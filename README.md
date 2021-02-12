# HelloGoService
A simple REST service for demos (written in Go). <br>
by William Mortl

## To Run

### Visual Studio Code

This solution contains a *devcontainer* that allows one to run and debug this REST service within Visual Studio Code. For a tutorial on 
devcontainers, visit [this link](https://code.visualstudio.com/docs/remote/containers-tutorial).

Simply go to the "Debug" icon (on the left) and click the "play" button (near the top).

### Command Line

#### Building 

1. Clone the repository into your Go path
2. In the project directory, first build the project using: **make**
3. Next, again in the project directory, create a Docker container: **docker build -t {image name}:{tag name} .**

#### Running the Container

From the CLI, run the container: **docker container run -p 8080:8080 -it {image name}:{tag}**

## Services Provided

Ping service: http://localhost:8080/ping <br>
Hello service: http://localhost:8080/hello?name=YourNameHere <br>
Prometheus metrics service: http://localhost:8080/metrics
