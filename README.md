# Go GRPC testing

Sample of GRPC tests written on go

[Demo allure report](https://nikita-filonov.github.io/sample_go_grpc_testing/)

## Setup project

You have to install [go](https://go.dev/doc/install)

Clone the project and install go packages

```shell
git clone https://github.com/Nikita-Filonov/sample_go_grpc_testing.git`
cd ./sample_go_grpc_testing
go mod download
```

## Run tests

```shell
make test
```

Generate allure report

```shell
cd tests
allure generate
```

Serve allure report

```shell
cd tests
allure serve
```