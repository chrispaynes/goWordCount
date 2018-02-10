src = ./pkg
main = ./cmd/main.go
pkgDir = $(src)/$(pkg)

.PHONY: benchmarks build coverage cpuProfiledockerUp pprof start test package vet

benchmarks:
	@go test -run=NONE -benchmem -bench=. $(src)/...

build:
	docker-compose build

coverage:
	@set -e;
	@echo "mode: set" > acc.out;

	@for Dir in $$(find . -type d); do \
		if ls "$$Dir"/*.go &> /dev/null; then \
			go test -coverprofile=profile.out "$$Dir"; \
			go tool cover -html=profile.out; \
		fi \
	done

	@rm -rf ./profile.out;
	@rm -rf ./acc.out;

cpuProfile:
	@go test -run=NONE -cpuprofile=cprof -bench=. $(src)/splitter
	@make pprof

dockerUp:
	@docker-compose down
	@docker-compose up -d
	@python -m webbrowser "http://localhost:8080" &> /dev/null

package:
	@mkdir -p $(pkgDir)
	@echo package $(pkg) | tee $(pkgDir)/$(pkg).go $(pkgDir)/$(pkg)_test.go

pprof:
	@go tool pprof --web cprof

start: 
	@rm -f $(main)/main
	@go run $(main)

test:
	@go test -v $(src)/...

vet:
	@go vet ./...
