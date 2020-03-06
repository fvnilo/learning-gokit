# stringsvc1
> The most basic example of a gokit service.

## Start
### Run service
```
go run main.go
```

### Make requests
```
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
```
