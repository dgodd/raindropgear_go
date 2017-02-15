Build with

```
GOOS=linux go build -ldflags="-s -w" -o app main.go
docker run --rm -v $PWD:/data  lalyos/upx app
```

NOTE: docker run --rm -v $PWD:/data  lalyos/upx -k --best --ultra-brute app

