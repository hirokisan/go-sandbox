# go-api

## Install
```
$ go get -u github.com/goadesign/goa/...
```

## Add API
```
$ touch design/xxx.go
$ goagen bootstrap -d github.com/hirokisan/go-sandbox/go-api/design/xxx -o xxx
```

## Run
```
$ go install ./api/...
```

## Todo
- [ ] goaをバージョン管理する？
- [ ] goaをforkして利用する？
- [ ] Makefileで動作管理する
