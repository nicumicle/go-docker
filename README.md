# Go-Docker
[![PR Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)](https://github.com/nicumicle/go-docker/pulls)
[![Go Report Card](https://goreportcard.com/badge/github.com/hellofresh/health-go)](https://goreportcard.com/report/github.com/nicumicle/go-docker)
[![CICD](https://github.com/nicumicle/go-docker/actions/workflows/ci-cd.yaml/badge.svg)](https://github.com/nicumicle/go-docker/actions/workflows/ci-cd.yaml)
[![GitHub License](https://img.shields.io/github/license/nicumicle/go-docker)](https://github.com/nicumicle/go-docker/blob/main/LICENSE)

**Go-Docker** is a Golang application designed for an efficient Docker container management.

This tool simplifies the orchestration process, providing essential features to streamline your containerized workflows.

## Features

- **List Containers:** Quickly view all Docker containers running on your system.
- **Start/Stop/Pause/Unpause/Delete Containers:** Take control of your containers with simple commands to manage their lifecycle.
- **Execute Commands:** Seamlessly execute commands within containers, enhancing your control and interaction.
- **Display Logs:**  Quickly view container logs
- **View Container Network and IP:** Quickly see network details for your container
- **Filter containers:** Filter containers by name,IP or network

## Requirements
- docker

## Demo

![go-docker](https://github.com/nicumicle/go-docker/blob/main/go-docker.gif?raw=true)

## Running the APP

To run **Go-Docker**, you have three options: either run the app, download the pre-built binaries or build it yourself from the source code.

### Run the GO APP
1. Clone the repository
    ```shell
    git clone https://github.com/nicumicle/go-docker.git
    cd go-docker
    ```
2. Run the app
    ```shell
    go run main.go
    ```

### Downloading Binaries
To get started with using **Go-Docker**, follow these steps to download the appropriate binaries for your operating system:

#### Windows
1. Go to the [releases](https://github.com/nicumicle/go-docker/releases) page of this repository.
2. Download the Windows version of the binary.
3. Place the downloaded **go-docker.exe** file in your desired local directory.

#### Linux / macOS
1. Visit the [releases](https://github.com/nicumicle/go-docker/releases) page of this repository.
2. Download the macOS version of the binary.
3. Move the downloaded **go-docker** file to your preferred local directory.

### Build from source
1. Clone the repository
    ```shell
    git clone https://github.com/nicumicle/go-docker.git
    cd go-docker
    ```
2. Build the app
    ```shell
    go build -o out/go-docker main.go
    ```
3. Run the application
- Windows:
   ```shell
   .\out\go-docker.exe
   ```
- Linux / macOS:
   ```shell
   ./out/go-docker
   ```

## Contributing

Contributions are welcome! Fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

Happy Containerizing! 🐳✨
