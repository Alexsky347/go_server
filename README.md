## Install all dependencies
go get ./...

## List upgradable dependencies
go list -m -u all

## Synchronize all dependencies
go mod tidy

An example of go server

## How to run

### Required Environment

- Postgres
- Redis

### Config
- You should modify `pkg/config/config.yaml`
````
$ docker compose -f docker/docker-compose.yml up -d
```

```yaml
environment: production
http_port: 8888
hostname: localhost
auth_secret: ######
database_uri: postgres://username:password@host:5432/database
redis_uri: localhost:6380
redis_password:
redis_db: 0
```

### Run
```shell script
$ make run
```
```
2023-09-12T15:18:36.684+0700    INFO    http/server.go:58       HTTP server is listening on PORT: 8888
2023-09-12T15:18:36.684+0700    INFO    grpc/server.go:53       GRPC server is listening on PORT: 8889
```

### Test
```shell script
$ make test
```

### Test with Coverage
```shell script
make test -timeout 9000s -a -v -coverprofile=coverage.out -coverpkg=./... ./...
```