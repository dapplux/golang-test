package main

import "fmt"

/*
	The code bellow should do the next:
	 - increment provided orders price
	 - add order id as suffix to the order name
	 - create map of order names with price values
	 - create slice of prices ans slice of names, fill them from created map
	 - basing on the prices slice print price of the first order, expected result is - Bar_1 - 1.56

	 Please, make a code review of this file. You may use your code editor for this.
*/

type Order struct {
	ID    uint64
	Name  OrderName
	Price float64
}

type OrderName string

func (t *OrderName) AddSuffix(n string) error {
	if n == "" {
		return fmt.Errorf("can't set order name, empty name")
	}
	val := (OrderName)(fmt.Sprintf("%s_%s", t, n))
	t = &val

	return nil
}

func (t OrderName) String() string {
	return string(t)
}

func renderOrdersPrice(orders []Order) map[OrderName]float64 {
	res := make(map[OrderName]float64)

	for i := range orders {
		_ = orders[i].Name.AddSuffix(string(orders[i].ID))
		res[orders[i].Name] = orders[i].Price
	}

	return res
}

func main() {
	orders := []Order{
		{
			ID:    1,
			Name:  "Bar",
			Price: 0.56,
		},
		{
			ID:    2,
			Name:  "Foo",
			Price: 1.18,
		},
	}

	for _, order := range orders {
		order.Price++
	}

	priceMap := renderOrdersPrice(orders)

	names := []OrderName{}
	prices := []uint64{}

	for name, price := range priceMap {
		names = append(names, name)
		prices = append(prices, uint64(price))
	}

	fmt.Printf("%s - %d", names[0], prices[0])
}
