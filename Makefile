TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)

build:
	go build -ldflags="-s -w -X main.version=$(TAG)" -o ./bin/scry-darwin-arm64

run:
	go run main.go

dev:
	sh ./dev-setup.sh

clean:
	go mod tidy

serve-docs:
	cd docs
	bundle exec jekyll serve

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out | grep total:

# tag:
# 	git tag -a v0.0.0 -m "test tag"
# 	git push --tags
