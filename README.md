# ALKETAS

## ABOUT THE PROJECT NAME

I use the "Random article"-Function for chose projectnames.

More About the Wikipedia article [ALKETAS](https://en.wikipedia.org/wiki/Alcetas_of_Macedon)

## PROJECT MISSION

An easy way to renew tokens for OpenBao in a Kubernetes environment

## ARCHITECTURAL DECISIONS

### Why wasn't the Golang lib used for OpenBao?

- https://github.com/openbao/openbao
- https://pkg.go.dev/github.com/openbao/openbao/api

The documentation is pretty poor, and I only needed one function. It would have taken me too much time to just deal with the lib.

## BUILD

### Application native

```bash
$ go build ./src/main.go
```

### Docker image

```bash
 $ podman build -t local-bao -f ./Dockerfile .
 $ podman run -it -v $(pwd)/examples/tokens.yaml:/alketas/confs/tokens.yaml:z \
   local-bao
```

You find the ready builded images hier: https://hub.docker.com/repository/docker/olafradicke/alketas/general

## RUN

### CONFIGURATION FILE

In this format:

```yaml
# This is an example

tokens:
  - name: "group-01"
    token: qwertyui
  - name: "grpup-02"
    token: asdfghjk
  - name: "grpup-03"
    token: zxcvbnm
```

### RUN AS SCRIPT (DEV/TEST)

```bash
$ cd src
$ go run ./main.go
```
