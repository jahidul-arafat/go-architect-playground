Goda: https://github.com/loov/goda
Gomod: https://github.com/Helcaraxan/gomod/tree/master

# Comprehensive Command Reference for Project Management
This project utilizes various tools and commands to streamline setup, testing, generation, and visualization tasks. This section provides a categorized list of commands for easy reference.

---

## Setup

| **Command**                         | **Description**                                                                                           |
|-------------------------------------|-----------------------------------------------------------------------------------------------------------|
| `go mod init <module_name>`         | Initialize a new Go module.                                                                               |
| `go mod tidy`                       | Clean up unused dependencies and update `go.mod` and `go.sum`.                                            |
| `go mod download`                   | Download all dependencies for the project.                                                                |
| `go env`                            | Display current Go environment variables.                                                                 |
| `go mod vendor`                     | Create a `vendor` directory with module dependencies.                                                     |
| `go get <package>@<version>`        | Add or update a dependency to a specific version.                                                         |
| `go mod verify`                     | Verify dependencies against their checksums.                                                              |
| `go list -m all`                    | List all modules used in the build.                                                                       |
| `gvm install go1.20`                | Install Go version 1.20 using Go Version Manager (GVM).                                                   |
| `gvm use go1.20`                    | Switch to Go version 1.20.                                                                                |
| `go env -w GO111MODULE=on`          | Enable Go modules explicitly.                                                                             |
| `go clean -modcache`                | Clear module cache, useful if dependency issues arise.                                                    |
| `go mod graph | dot -Tsvg -o graph.svg` | Generate an SVG dependency graph.                                                                        |
| `go mod edit -replace=<module>=<new_source>` | Replace a module source temporarily in `go.mod`.                                                   |
| `go get -u all`                     | Update all dependencies to their latest compatible versions.                                              |

---

## Dependency Management

| **Command**                                   | **Description**                                                                                           |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| `gomod list all`                              | List all dependencies (direct and indirect) in the project.                                               |
| `gomod analyse`                               | Analyze and list all project dependencies, including indirect ones.                                       |
| `gomod graph -p`                              | Generate a package-level dependency graph.                                                                |
| `gomod outdated`                              | List outdated dependencies with newer available versions.                                                 |
| `gomod paths <from> <to>`                     | Display dependency paths between specified modules.                                                       |
| `gomod exec go list ./...`                    | List project packages and dependencies recursively.                                                       |
| `gomod graph -a`                              | Annotate dependency graph with version information.                                                       |
| `gomod weight`                                | Analyze weight and size impact of dependencies.                                                           |
| `goda list all`                               | List all dependencies in the project.                                                                     |
| `goda graph -short`                           | Create a simplified dependency graph for the project.                                                     |
| `goda cycles ./...`                           | Detect dependency cycles within the project.                                                              |
| `goda flatten ./...`                          | Generate a flattened list of dependencies.                                                                |
| `goda stats ./...`                            | Show statistics about dependencies, such as package count and import frequency.                           |
| `goda import <package>`                       | Show all packages importing a specific package.                                                           |
| `go list -m -json all`                        | List all modules with detailed JSON information.                                                          |
| `go mod why -m <module>`                      | Explain why a specific module is used.                                                                    |
| `go mod tidy -e`                              | Remove unused dependencies, ignoring errors.                                                              |
| `goda tree ./...`                             | Display a hierarchical dependency tree for in-depth analysis.                                             |

---

## Testing

| **Command**                               | **Description**                                                                                           |
|-------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| `go test ./...`                           | Run all tests in the project recursively.                                                                 |
| `go test -cover ./...`                    | Run tests and show code coverage.                                                                         |
| `go test -coverprofile=coverage.out ./...` | Generate a coverage profile for detailed coverage analysis.                                               |
| `go tool cover -html=coverage.out`        | View coverage report in a browser.                                                                        |
| `go test -race ./...`                     | Run tests with race condition detection enabled.                                                          |
| `go test -bench .`                        | Run benchmarks to evaluate performance.                                                                   |
| `go test -count=1 ./...`                  | Run tests without caching results.                                                                        |
| `go test -cpu 1,2,4 ./...`                | Run tests with multiple CPU cores.                                                                        |
| `go test -memprofile=mem.out`             | Generate a memory profile for analysis.                                                                   |
| `go test -cpuprofile=cpu.out`             | Generate a CPU profile for analysis.                                                                      |
| `go test -trace trace.out`                | Generate a trace file for test execution analysis.                                                        |
| `go tool trace trace.out`                 | Visualize and analyze the trace file for performance diagnostics.                                         |
| `go test -short ./...`                    | Run short tests only, skipping longer cases.                                                              |
| `go test -parallel 4 ./...`               | Run tests in parallel with a specified max goroutine limit.                                               |
| `go tool pprof cpu.out`                   | Analyze CPU profile to identify performance bottlenecks.                                                  |
| `go tool pprof mem.out`                   | Analyze memory profile for potential optimizations.                                                       |
| `go test -benchmem ./...`                 | Benchmark with memory allocation analysis.                                                                |
| `go test -run <TestName>`                 | Run a specific test by name.                                                                              |

---

## Code Generation

| **Command**                                    | **Description**                                                                                           |
|------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| `go generate ./...`                            | Run all `go:generate` directives in the project.                                                          |
| `go generate -run stringer`                    | Run only `go:generate` commands using `stringer`.                                                         |
| `go run generate_schema.go`                    | Run a Go script to generate a JSON schema for `Student`.                                                  |
| `go install github.com/alecthomas/jsonschema`   | Install JSON schema generation tool.                                                                      |
| `//go:generate stringer -type=Grade`           | Generate `String` method for `Grade` type.                                                                |
| `go generate -run mockgen`                     | Run only `mockgen` commands in `go:generate` directives.                                                  |
| `go install github.com/onsi/ginkgo/v2/ginkgo@latest` | Install `ginkgo` for enhanced testing capabilities.                                                  |
| `go install golang.org/x/tools/cmd/stringer@latest` | Install `stringer` for generating `String` methods.                                                   |
| `go install github.com/golang/mock/mockgen@latest` | Install `mockgen` to generate mock files for interfaces.                                           |
| `go install github.com/google/wire/cmd/wire@latest` | Install Wire for dependency injection code generation.                                                 |
| `wire`                                         | Generate dependency injection code with Wire.                                                             |
| `//go:generate protoc --go_out=. student.proto` | Generate code from a protobuf file.                                                                       |
| `go install github.com/fatih/gomodifytags`     | Install tool to modify struct tags easily.                                                                |
| `gomod exec go build`                          | Run build with all modules included via `gomod`.                                                          |
| `go build`                                     | Build the project to verify code generation completeness.                                                 |
| `go generate -x ./...`                         | Display commands run by `go generate` for debugging purposes.                                             |
| `go install github.com/cweill/gotests/...`     | Install `gotests` to automatically generate test files from code.                                         |

---

## Visualization

| **Command**                                 | **Description**                                                                                           |
|---------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| `gomod graph -o dependencies.dot`           | Output dependency graph in DOT format.                                                                    |
| `dot -Tpng dependencies.dot -o dependencies.png` | Convert the DOT file to PNG.                                                                         |
| `go-callvis -focus main ./`                 | Generate a call graph in the browser.                                                                     |
| `goda graph -short -cluster -o goda_graph.svg` | Create a dependency graph with `goda` and output to SVG.                                               |
| `gomod analyse`                             | Analyze project dependencies in detail.                                                                   |
| `goda tree ./...`                           | Display a hierarchical dependency tree for in-depth analysis.                                             |
| `goda cycles ./...`                         | Detect dependency cycles within the project.                                                              |
| `goda paths <from> <to>`                    | Show dependency path between two packages.                                                                |
| `gomod list unused`                         | List unused dependencies to identify potentially removable modules.                                       |
| `goda stats ./...`                          | Show statistics about dependencies, such as package count and import frequency.                           |
| `gomod graph -a`                            | Display dependency graph with version annotations.                                                        |
| `goda import <package>`                     | Show all packages importing a specific package.                                                           |
| `gomod exec go list -json ./...`            | Visualize packages and dependencies in JSON format.                                                       |
| `gomod weight`                              | Analyze dependency weights and size impact.                                                               |

---

This `README.md` provides a categorized list of commands for managing dependencies, running tests, generating code, visualizing dependencies, and more.

This Markdown format should work well in .md files with tables appearing correctly.
