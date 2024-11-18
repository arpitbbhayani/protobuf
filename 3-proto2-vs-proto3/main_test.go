package main

import (
	"fmt"
	"testing"

	"github.com/arpitbbhayani/protobuf/models"
)

var order = models.Order{}
var payment = models.Payment{}

func TestDefaults(t *testing.T) {
	fmt.Println(order.OrderId)
	fmt.Println(order.Amount)
	fmt.Println(order.Customer)
	fmt.Println(payment.Amount, payment.GetAmount())
}
