package main

import (
	"go_study/c32_customer_management/view"
)

func main() {
	customerView := view.NewCustomerView()
	customerView.Menu()
}
