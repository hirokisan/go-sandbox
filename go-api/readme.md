# go-api

## Prepare
```
$ go get -u github.com/goadesign/goa/...
```

## Add API
```
$ touch design/xxx.go
$ goagen bootstrap -d github.com/hirokisan/go-sandbox/go-api/design/xxx -o xxx
```

## Install
```
$ make install
```

## Create&Update design
```
$ API_NAME=xxx make design
```

## Create&Update API
```
$ API_NAME=xxx make api
```

## Todo
- [ ] goaをバージョン管理する？
- [ ] goaをforkして利用する？
- [ ] Makefileで動作管理する
- [ ] scaffoldを用意する？
