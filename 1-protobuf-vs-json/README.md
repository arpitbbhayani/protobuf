```sh
protoc --go_out=. person.proto
go test -v . -bench=.
```

```
{
  "name": "Jane Doe",
  "id": 2,
  "email": "jane.doe@example.com",
  "phone": "123-456-7890",
  "address": "123 Main St, Anytown, USA",
  "age": 28,
  "gender": "Female",
  "occupation": "Engineer",
  "nationality": "American",
  "marital_status": "Single"
}
=== RUN   TestSizeComparison
    main_test.go:86: JSON size: 220 bytes
    main_test.go:87: Protobuf size: 113 bytes
--- PASS: TestSizeComparison (0.00s)
goos: linux
goarch: amd64
pkg: github.com/arpitbbhayani/protobuf
cpu: AMD Ryzen 7 4800U with Radeon Graphics         
BenchmarkJSONSerialization
BenchmarkJSONSerialization-16             847700              1821 ns/op
BenchmarkJSONDeserialization
BenchmarkJSONDeserialization-16           160700              6994 ns/op
BenchmarkProtobufSerialization
BenchmarkProtobufSerialization-16        2009774               707.7 ns/op
BenchmarkProtobufDeserialization
BenchmarkProtobufDeserialization-16      1000000              1531 ns/op
```
