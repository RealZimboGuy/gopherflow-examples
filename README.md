# GopherFlow Example

Two workflows implemented using the [GopherFlow](https://github.com/RealZimboGuy/gopherflow) library:

* **DemoWorkflow** — walks through a few fictitious steps and sets state variables
* **GetIpWorkflow** — fetches the public IP address from ifconfig.io into a state variable

Prerequisites:
- Go 1.26+ (or Docker if you prefer containers)

## Run it

    go run ./cmd/gopherflow_example

The example sets `GFLOW_DATABASE_TYPE=SQLLITE` in `main.go`, so it starts against a
local `gflow.db` file with no other setup. Access the web console at
http://localhost:8080/

    Username : admin
    Password : admin

## Run it in Docker

    docker build -t gopher-flow-example:1.0 -f Dockerfile.multistage .

    docker run -p 8080:8080 \
    -e GFLOW_DATABASE_TYPE=SQLLITE \
    -e GFLOW_DATABASE_SQLLITE_FILE_NAME=/data/gflow.db \
    -v "$(pwd)/data:/data" \
    gopher-flow-example:1.0

GopherFlow uses a pure-Go SQLite driver, so the image is a static binary and needs
no special container privileges.

## What to look at

| Path | What it shows |
|---|---|
| `cmd/gopherflow_example/main.go` | registering workflows and starting the engine |
| `internal/workflows/` | the two workflow definitions |
| `controllers/` | adding your own HTTP routes alongside the console |
