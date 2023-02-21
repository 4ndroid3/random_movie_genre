package users

import (
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {
	var total float32
	for _, v := range item {
		total += v
	}
	return total * 1.25
}
