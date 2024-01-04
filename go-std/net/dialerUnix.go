package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	var d net.Dialer

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	d.LocalAddr = nil // if you have a local addr, add it here

	raddr := net.UnixAddr{
		Name: "/path/to/unix.sock",
		Net:  "unix",
	}

	conn, err := d.DialContext(ctx, "unix", raddr.String())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, socket!")); err != nil {
		log.Fatal(err)
	}
}
