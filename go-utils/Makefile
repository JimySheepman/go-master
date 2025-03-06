SERVICE=example-service
COMMIT_ID=$(shell git rev-list --abbrev-commit -1 HEAD -- .)
REVISION_ID=$(shell git rev-list --count HEAD)
BRANCH=$(shell git symbolic-ref --short HEAD)
GIT_USERNAME="test_user"
GIT_ACCESS_TOKEN="tSyGeCcWp7PYdsjjYViB"
COMMITTER_NAME=$(shell git config user.name)
COMMITTER_EMAIL=$(shell git config user.email)

all: prevent_dirty run_tests clean build package

build: clean
	mkdir -p build
	docker build -t "registry.birlesikodeme.com/${SERVICE}:${COMMIT_ID}" --build-arg commitID=${COMMIT_ID} --build-arg revisionID=${REVISION_ID} --build-arg GIT_USERNAME=${GIT_USERNAME} --build-arg GIT_ACCESS_TOKEN=${GIT_ACCESS_TOKEN} .

package: build
	docker push registry.birlesikodeme.com/${SERVICE}:${COMMIT_ID}

deploy: prevent_dirty alert_slack
  echo deploy

clean:
	rm -rf ./build

test:
	go test ./...

test_db:
	go test --tags=db ./...

alert_slack:
	curl -X POST -H 'Content-type: application/json' --data '{"text":"*Application:* ${SERVICE} *Branch:* ${BRANCH} *Commit:* ${COMMIT_ID} *Author:* ${COMMITTER_NAME}"}' ${SLACK_HOOKED}

# Prevents command from executing with dirty git revision
prevent_dirty:
	@if [ -z "$(shell git status --porcelain .)" ]; then \
		echo "Working directory clean, proceeding..."; \
	else \
		echo "\033[31mError: Can not continue with uncommitted changes."; exit 1; \
	fi

run_tests:
	@if [ "$(shell go test ./... > /dev/null && echo $$?)" = "0" ]; then \
		echo "All tests passed successfully, proceeding..."; \
	else \
		if [ "$(BRANCH)" = "master" ]; then \
			echo "\033[31mError: Can not continue with failed test(s)"; exit 1; \
		else \
			echo "\033[33mWarning: Proceeding with failed test(s)\e[0m"; \
		fi \
	fi
