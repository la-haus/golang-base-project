# Sample Golang API

## Cloning into the GOPATH

```sh
git clone git@github.com:la-haus/golang-base-project.git $GOPATH/src/github.com/la-haus/sample
```

## Go to the project directory

```sh
cd $GOPATH/src/github.com/la-haus/sample
```

## Requirements

* go v1.16+
* go module
* mockery
* swag

## Build

* Copy the sample configuration file and fill the right values according to your environment

```sh
cp .env.sample .env
```

* Install dependencies:

```sh
go mod download
```

* Generate Swagger doc:

```sh
go get github.com/swaggo/swag/cmd/swag
swag init -g router/router.go
```

* Run test:

```sh
go test ./...
```

* How to (re)generate mocks using **Mockery** tool:
```sh
mockery --all --dir interfaces --output ./tests/mocks
```


## Open the API Docs

Then, once app is up and running, you can go to http://`<host>`:`<port>`/api/sample/docs/index.html

## Environments

### Required environment variables

* `SERVER_HOST`: Host for the server
* `SERVER_PORT`: Port for the server
* `DD_ENV`: Datadog environment, values can be `staging` or `production`
* `DD_SERVICE`: Datadog service name
* `DD_VERSION`: Datadog service version
* `DATADOG_AGENT_HOST`: Datadog agent hostname
* `DATADOG_AGENT_PORT`: Datadog agent port


## Contributors

* Luis Pereira

## License

This project is property of La Haus
