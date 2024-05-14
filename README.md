# go-mono-repo-templater

go-mono-repo-templater is a command-line tool written in Go that generates a mono repo service structure for managing microservices-based projects. It creates a directory structure with files organized for each service, including command-line executables, internal packages, APIs, and Dockerfiles.

## Features

- Quickly create a mono repo structure for managing microservices.
- Automatically generates directory structure and files for each service.
- Easy customization with command-line arguments.

## Installation

To use go-mono-repo-templater, you need to have Go installed on your system. You can install it using the following command:

```bash
go install github.com/mcmuralishclint/go-mono-repo-templater
```

## Usage
To generate a mono repo structure, run the following command:

```bash
go-mono-repo-templater -dir /path/to/ -services service1,service2,service3
```

Replace /path/to/ with the path to your repo and service1,service2,service3 with a comma-separated list of service names.

## Directory Structure
```bash
scheduliz/
  |- services/
  |   |- auth/
  |   |   |- cmd/
  |   |   |   |- main.go
  |   |   |- internal/
  |   |   |   |- auth.go
  |   |   |- api/
  |   |   |   |- handler.go
  |   |   |   |- router.go
  |   |   |- Dockerfile
  |   |
  |   |- scheduler/
  |   |   |- ...
  |   |
  |   |- ...
  |
  |- pkg/
  |   |- utils/
  |   |   |- ...
  |
  |- Makefile
  |- go.mod
  |- go.sum
```

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

You can simply copy and paste this content into a file named `README.md` in the root directory of your GitHub repository.
