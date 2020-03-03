# Sudoku Builder

[![Go Report Card](https://goreportcard.com/badge/github.com/armsnyder/sudoku-builder)](https://goreportcard.com/report/github.com/armsnyder/sudoku-builder)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/armsnyder/sudoku-builder)](https://hub.docker.com/r/armsnyder/sudoku-builder)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/armsnyder/sudoku-builder)](https://hub.docker.com/r/armsnyder/sudoku-builder)

Pet project to explore using AI to assist with puzzle building

## Quick Start (Docker)

Requirements
* Docker

```
$ docker build -t sudoku-builder .
$ docker run -p 8080:8080 sudoku-builder
```

Browse to http://localhost:8080

## Quick Start (With UI Hot Reloading)

Requirements
* Golang (v1.13+)
* npm
* yarn

```
$ cd go
$ go run .
```

```
$ cd vue
$ yarn install
$ yarn serve
```

Browse to http://localhost:8080
