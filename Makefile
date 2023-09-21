DEFAULT: build

# For building the project locally
install:


# Cleanup unused dependencies
cleanup:
	go mod tidy

build:
	@./scripts/build.sh
