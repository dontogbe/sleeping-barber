package main

import (
	"dontogbe/sleeping-barber/barbersvc"
	"fmt"
	"time"
)

func main() {
	barber := barbersvc.NewBarber("Don")
	shop := barbersvc.NewShop(barber, 5)
	go barber.Work(shop)
	time.Sleep(time.Second * 1)
	custs := []string{"AB", "Seth", "Grace", "Comfort", "Nancy"}
	for _, c := range custs {
		cust := barbersvc.NewCustomer(c)
		go cust.AddCustomer(shop)
	}
	// press a key to finish
	fmt.Scanln()
}
