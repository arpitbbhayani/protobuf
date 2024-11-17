package main

import (
	"fmt"
	"testing"

	"github.com/arpitbbhayani/protobuf/models"
)

var order = models.Order{
	OrderId: "o1",
	Amount:  100,
	Customer: &models.Customer{
		Id:    "1",
		Name:  "Arpit Bhayani",
		Email: "arpit@bhayani.com",
	},
}

func Test1(t *testing.T) {
	fmt.Println(&order)
}
