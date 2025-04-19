# Makefile

# Variables
VERSION_FILE=VERSION
VERSION=$(shell cat $(VERSION_FILE))
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
OUTPUT=task-cli

# Targets

# Build the binary with version info
build:
	go build -o $(OUTPUT) \
	-ldflags "-X 'task-tracker-cli/internal/cmd.version=$(VERSION)' \
	          -X 'task-tracker-cli/internal/cmd.commit=$(COMMIT)' \
	          -X 'task-tracker-cli/internal/cmd.date=$(DATE)'" \
	main.go

# Run the binary (builds first)
run: build
	./$(OUTPUT)

# Clean up the binary
clean:
	rm -f $(OUTPUT)

# Show current version
version:
	@echo $(VERSION)

# Bump patch version (e.g., 0.1.0 -> 0.1.1)
bump-patch:
	@$(eval NEW_VERSION=$(shell echo $(VERSION) | awk -F. '{$$3+=1; print $$1"."$$2"."$$3}'))
	@echo $(NEW_VERSION) > $(VERSION_FILE)
	@echo "Bumped patch version to $(NEW_VERSION)"

# Bump minor version (e.g., 0.1.0 -> 0.2.0)
bump-minor:
	@$(eval NEW_VERSION=$(shell echo $(VERSION) | awk -F. '{$$2+=1; $$3=0; print $$1"."$$2"."$$3}'))
	@echo $(NEW_VERSION) > $(VERSION_FILE)
	@echo "Bumped minor version to $(NEW_VERSION)"

# Bump major version (e.g., 0.1.0 -> 1.0.0)
bump-major:
	@$(eval NEW_VERSION=$(shell echo $(VERSION) | awk -F. '{$$1+=1; $$2=0; $$3=0; print $$1"."$$2"."$$3}'))
	@echo $(NEW_VERSION) > $(VERSION_FILE)
	@echo "Bumped major version to $(NEW_VERSION)"

# Full release (bump, build)
release: bump-patch build
	@echo "Released version $(VERSION)"

# Cross-compile for all platforms
cross-compile:
	GOOS=linux GOARCH=amd64 go build -o dist/task-cli-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/task-cli-mac main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/task-cli-mac-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/task-cli.exe main.go

# Zip all builds
package: cross-compile
	cd dist && zip task-cli-linux.zip task-cli-linux
	cd dist && zip task-cli-mac.zip task-cli-mac
	cd dist && zip task-cli-mac-arm64.zip task-cli-mac-arm64
	cd dist && zip task-cli.exe.zip task-cli.exe