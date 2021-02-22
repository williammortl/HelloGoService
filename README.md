# HelloGoService
A simple REST service for demos (written in Go). <br>
by William Mortl <br><br>

This example collection of Go REST services demostrate several important Go principles: <br><br>
1. **ping.go** - a simple Go REST service
2. **hello.go** - a Go REST service that reads from the query string
3. **dbadd.go** & **dbget.go** - adds and retrieves records from an in-memory database, shows how to read JSON from the post body, and demonstrates how to read a variable from the REST path
4. **math.go** - demonstrates Go routines... which is a Go "thread"

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

### REST Services Documentation
**Ping service**: http://localhost:8080/Ping <br><br>
**Hello service**: http://localhost:8080/Hello?name=YourNameHere <br><br>
**Database Get service** (retrieves by ID which is an int value, replace the trailing 0 with any int): http://localhost:8080/Db/0 <br><br>
**Database Add / Update service** (adds / updates by ID, replace the trailing 4 with any int): http://localhost:8080/Db/4 <br><br>
Post JSON in this format to the above URL to add this person to the database with id 4: <br>
```
{
    "name": "Taylor Swift",
    "address": "10th Street",
    "phone": "6155551212",
}
``` 
<br> **Math service** (allowed operators: Add, Subtract, Multiple): http://localhost:8080/Math/Add
<br><br>
Post JSON in this format to the above URL to sum the numbers below: <br>
```
{
    "numbers": [
        1,
        2,
        3,
        4
    ]
}
```

### Telemetry Metrics Documentation

Metrics are provided by [Prometheus](https://prometheus.io/). <br>

**Prometheus metrics service**: http://localhost:8080/metrics

### Swagger Documentation

Service documentation provided by [swaggo](https://github.com/swaggo/http-swagger). <br>

**Swagger service**: http://localhost:8080/swagger/index.html