# Flights Test Assignment

## Prerequisites:

* Install Golang
    * Choose a binary release for your system [from downloads](https://golang.org/dl/).
    * After install, check with `go version` in terminal.
* Install `docker`
    * For example, [Docker Desktop for Mac](https://docs.docker.com/desktop/mac/install/)

## Run the service

```shell
make start
```

## Run tests

```shell
make test
```

## Benchmark

```
$ make bench
goos: darwin
goarch: arm64
pkg: flights-test/internal/services/calculator
BenchmarkCalculator_Calculate
BenchmarkCalculator_Calculate-10    	       9	 118881315 ns/op
PASS
ok      flights-test/internal/services/calculator       2.948s

```

## API Documentation

### POST `/calculate`

Request: array of flights. Each flight is an array with the departure and arrival airports

> The max possible number of flights is about 1051200. When someone who born in some airport travels every day for the whole 120-years life without any breaks. And it's avg flight takes 30 minutes and 30 minutes waiting in airport (1 hour per flight in total)

```json
[
  [ "SFO", "ATL" ],
  [ "ATL", "GSO" ],
  [ "GSO", "IND" ],
  [ "IND", "EWR" ]
]
```

Response: 

```json
200 OK

[ "SFO", "EWR" ]
```

```json
400 Bad Request

{ "error": "incorrect path" }
```

Example:

```shell
curl -X POST --location "http://localhost:8080/calculate" \
    -H "Content-Type: application/json" \
    -d "[
          [ \"SFO\", \"ATL\" ],
          [ \"ATL\", \"GSO\" ],
          [ \"GSO\", \"IND\" ],
          [ \"IND\", \"EWR\" ]
        ]"
```