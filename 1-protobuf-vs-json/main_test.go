package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

var person = Person{
	Name:          "Jane Doe",
	Id:            2,
	Email:         "jane.doe@example.com",
	Phone:         "123-456-7890",
	Address:       "123 Main St, Anytown, USA",
	Age:           28,
	Gender:        "Female",
	Occupation:    "Engineer",
	Nationality:   "American",
	MaritalStatus: "Single",
}

func init() {
	v, _ := json.MarshalIndent(&person, "", "  ")
	fmt.Println(string(v))
}

func BenchmarkJSONSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&person)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONDeserialization(b *testing.B) {
	data, err := json.Marshal(&person)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		var p Person
		err := json.Unmarshal(data, &p)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkProtobufSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&person)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkProtobufDeserialization(b *testing.B) {
	data, err := proto.Marshal(&person)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		var p Person
		err := proto.Unmarshal(data, &p)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestSizeComparison(t *testing.T) {
	jsonData, err := json.Marshal(&person)
	if err != nil {
		t.Error(err)
	}

	protoData, err := proto.Marshal(&person)
	if err != nil {
		t.Error(err)
	}

	t.Logf("JSON size: %d bytes", len(jsonData))
	t.Logf("Protobuf size: %d bytes", len(protoData))
}
