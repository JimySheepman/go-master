package integration

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	username    = "sa"
	image       = "mcr.microsoft.com/mssql/server"
	exposedPort = "1433/tcp"
)

// Container represents a mock-server container
type Container struct {
	Container testcontainers.Container
	req       ContainerRequest
}

// ContainerRequest a container request specification
type ContainerRequest struct {
	testcontainers.GenericContainerRequest
	Username string
	Password string
	DbName   string
	Image    string
	Logger   *testcontainers.LogConsumer
}

// CreateMsSqlContainer creates a SQL Server for Linux container
func CreateMsSqlContainer(ctx context.Context, req ContainerRequest) (*Container, error) {
	log.Println("Starting mssql container...")
	if req.Env == nil {
		req.Env = make(map[string]string)
	}
	req.Env["ACCEPT_EULA"] = "Y"
	if req.Image == "" {
		req.Image = image
		req.Env["MSSQL_PID"] = "Express"
	}
	req.GenericContainerRequest.Image = req.Image
	if req.ExposedPorts == nil {
		req.ExposedPorts = []string{exposedPort}
	}

	if req.Username == "" {
		req.Username = username
	}
	if req.Password == "" {
		return nil, errors.New("a password must be provided")
	}
	req.Env["SA_PASSWORD"] = req.Password
	if req.WaitingFor == nil {
		req.WaitingFor = wait.ForSQL(exposedPort, "sqlserver", func(port nat.Port) string {
			fmt.Printf("sqlserver://%s:%s@localhost:%s\n", req.Username, req.Password, port.Port())
			return fmt.Sprintf("sqlserver://%s:%s@localhost:%s", req.Username, req.Password, port.Port())
		}).Timeout(time.Minute)
	}

	provider, err := req.ProviderType.GetProvider()
	if err != nil {
		return nil, err
	}

	result := &Container{
		req: req,
	}

	req.Started = false
	if result.Container, err = provider.CreateContainer(ctx, req.ContainerRequest); err != nil {
		return result, errors.Wrap(err, "could not create container")
	}

	if req.Logger != nil {
		if err = result.Container.StartLogProducer(ctx); err != nil {
			return result, errors.Wrap(err, "could not start log producer")
		}
		result.Container.FollowOutput(*req.Logger)
	}

	if err = result.Container.Start(ctx); err != nil {
		return result, errors.Wrap(err, "could not start container")
	}
	return result, nil
}

// GoConnectionString returns a connection string suitable for usage in Go
func (c *Container) GoConnectionString(ctx context.Context) (string, error) {
	host, port, err := c.HostAndPort(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s", c.req.Username, c.req.Password, host, port.Port()), nil
}

// HostAndPort retrieves the external host and port of the container
func (c Container) HostAndPort(ctx context.Context) (string, nat.Port, error) {
	if c.Container != nil {
		return canned.GetHostAndPort(ctx, c.Container, exposedPort)
	}
	return "", "", errors.New("could not read host and port from a nil pointer")
}

// ReadFileAndExec read .sql file end exec
func ReadFileAndExec(name string, conn *sql.DB, panicOnError bool) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	content, err := os.Open(filepath.Join(path, name))
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(content)
	a := strings.Builder{}
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, "--") {
			a.Write([]byte(text))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	_, err = conn.Exec(a.String())
	if err != nil {
		if panicOnError {
			log.Fatal(err)
		} else {
			log.Println(err)
		}
	}
}
