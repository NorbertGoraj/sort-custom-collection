package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"sort"
)

func main() {
	eList := []product{
		product{name: "TV", price: decimal.NewFromFloat(12350.20)},
		product{name: "iPad", price: decimal.NewFromFloat(16679.22)},
		product{name: "Mobile", price: decimal.NewFromFloat(12000)},
	}
	sort.Sort(productList(eList))
	for _, employee := range eList {
		fmt.Printf("Name: %s Salary %s\n", employee.name, employee.price)
	}
}

type product struct {
	name  string
	price decimal.Decimal
}

type productList []product

func (product productList) Len() int {
	return len(product)
}

func (product productList) Less(i, j int) bool {
	return product[i].price.LessThan(product[j].price)
}

func (product productList) Swap(i, j int) {
	product[i], product[j] = product[j], product[i]
}
