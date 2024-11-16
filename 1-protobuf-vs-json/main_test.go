package main

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/proto"
)

func BenchmarkJSONSerialization(b *testing.B) {
	person := &Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(person)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONDeserialization(b *testing.B) {
	person := &Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	data, err := json.Marshal(person)
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
	person := &Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(person)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkProtobufDeserialization(b *testing.B) {
	person := &Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	data, err := proto.Marshal(person)
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
	person := Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}

	jsonData, err := json.Marshal(person)
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
