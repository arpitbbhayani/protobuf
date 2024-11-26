```sh
protoc --go_out=./models ./protos/*
go test -v . -bench=.    
```
