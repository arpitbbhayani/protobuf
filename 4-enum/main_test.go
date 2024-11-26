package main

import (
	"fmt"
	"testing"

	"github.com/arpitbbhayani/protobuf/models"
)

var order2 = models.Order2{}
var order3 = models.Order3{}

func TestDefaults(t *testing.T) {
	fmt.Println("default value when member accessed", order2.Status)
	fmt.Println("default value when called GetStatus", order2.GetStatus()) // Defauult 1st value
	order2.Status = models.OrderStatus2_ORDER_STATUS2_DELIVERED.Enum()
	fmt.Println("value after setting Status", order2.GetStatus())
	fmt.Println("for proto3, default value", order3.Status)
}
