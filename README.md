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

### Build and test Docker image

```bash
 $ podman build -t local-bao -f ./Dockerfile .
 $ podman run -it -v $(pwd)/examples/tokens.yaml:/alketas/confs/tokens.yaml:z \
   local-bao
```

You find the ready builded images hier: https://hub.docker.com/repository/docker/olafradicke/alketas/general

## RUN / DEPLOY IN KUBERNETES

You find a configuration example in th directory `/examples` for a `CronJob`-Deployment and the needed `Secret`s. Changed and enter:

```bash
$ kubectl apply -n <my-namespace> -f /examples
```

### RUN AS SCRIPT FOR DEV/TEST

```bash
$ cd src
$ export VAULT_RENEW_TOKENS='../examples/tokens.yaml'
$ go run ./main.go
```

## DEBUGGING

For security reasons, the CronJob logs very little. To efficiently search for sources of error, there is in `examples` a [manifest](examples/Pod-alketas-debug.yaml) for a debug pod that you can log into and use interactively.
