# Build
```
go build -ldflags="-s -w" -o timer main.go
```
# run
```
./timer 5s or ./timer 0h0m5s
```

```bash
$ go run main.go 5s
Counting on <<<0:0:5
time‘s up
```