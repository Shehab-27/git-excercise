package main

import "fmt"

type customer struct {
	tel     string
	balance float64
	name    string
}

var bank = map[int]customer{
	1: {tel: "0123456766", balance: 80909, name: "awad"},
	2: {tel: "012323232", balance: 2909322, name: "abdelslam"},
	3: {tel: "012456123", balance: 900000, name: "hassan"},
	4: {tel: "0151212156", balance: 1500000, name: "alaaeldin"},
}

func sales(m map[int]customer) {
	for _, cust := range m {
		if cust.balance > 1000000 {
			fmt.Println("customer's name is :", cust.name, "customer's phone number is :", cust.tel)
		}
	}
}

func main() {

	sales(bank)

}
