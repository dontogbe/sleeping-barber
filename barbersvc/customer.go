package barbersvc

import (
	"fmt"
	"time"
)

type Customer struct {
	name string
}

func (c *Customer) AddCustomer(shop *Shop) {
	select {
	case shop.waitingRoom <- c:
		// yes ! client found a seat
		fmt.Printf(":) %s found a seat\n", c.name)
		select {
		// we check that the barber is not sleeping
		case shop.barber.wakeMe <- c:
			// lazy barber ! wake him up !
			fmt.Printf("WAKE UP ! screams %s\n", c.name)
		default:
			// default... do nothing (default is necessary there)
		}
		return // stop the problem, client is now managed by barber

	default:
		// no seat (write on channel failed)
		fmt.Printf("... %s will be back later ...\n", c.name)
		time.Sleep(time.Millisecond * 100)
	}
}

// NewCustomer initializes a Customer struct
func NewCustomer(name string) *Customer {
	c := Customer{name: name}
	return &c
}
