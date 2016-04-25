all: gingx

gingx:  gingx.go
	go build gingx.go

run: gingx.go
	go run gingx.go .

