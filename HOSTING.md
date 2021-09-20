# Self-Hosting

[esm.sh](https://esm.sh) provides a fast, global content delivery network publicly, but you also can deploy your own CDN for ES Modules.<br>
You will need [Go](https://golang.org/dl) 1.16+ to compile the server. The server runtime will install the nodejs (14 LTS) automatically.

## Clone code

```bash
TARGET_DIR=...
REPO=https://github.com/alephjs/esm.sh ## or alternative fork (eg, `https://github.com/rivy-go/esm.sh`)
git clone "${REPO}" "${TARGET_DIR}"
cd "${TARGET_DIR}"
```

## Run the sever locally

```bash
go run main.go --port=8080 --dev
## or install and run
# go install
# esm.sh --port=8080 --dev
```

then you can import `React` from http://localhost:8080/react

## Deploy to single host

Please ensure the [supervisor](http://supervisord.org/) installed on your host machine.

```bash
sh ./scripts/deploy.sh
```

## Deploy to multiple hosts

- deploy manually
- deploy automatically

_We are working on it._

## Deploy with Docker

An example [Dockerfile](./Dockerfile) is found in the root of this project.
