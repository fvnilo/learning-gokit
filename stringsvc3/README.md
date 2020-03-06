# stringsvc2
> The same thing as stringsvc1 but with a better structure and middlewares.

## Start
### Run service
```
go run *.go
```

### Make requests
```
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
curl localhost:8080/metrics
```
