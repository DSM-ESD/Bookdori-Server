# Bookdori-Server
북도리의 서버입니다!

## Build and Run

```bash
$ make
go build -o bin/bookdori-server main.go
$ bin/bookdori-server
[negroni] listening on :3000
...생략
```

만약 `cannot find package ...` 오류가 발생했다면 아래의 명령을 수행한 후에 다시 시도해주세요.

```bash
$ make install-pkg
go get "github.com/codegangsta/negroni"
go get "github.com/julienschmidt/httprouter"
...생략
```