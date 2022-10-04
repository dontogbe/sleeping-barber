package barbersvc

import (
	"fmt"
	"time"
)

type Barber struct {
	name   string
	wakeMe chan *Customer
}

func (b *Barber) Work(shop *Shop) {
	for {
		select {
		case c := <-shop.waitingRoom:
			fmt.Printf("%s cuts %s's hair\n", b.name, c.name)
			time.Sleep(time.Millisecond * 6)
			fmt.Printf("%s finished %s\n", b.name, c.name)
		default:
			fmt.Printf("%s wants to sleep...\n", b.name)
			c := <-b.wakeMe // not a "select case" so it blocks there
			fmt.Printf("%s awake %s\n", c.name, b.name)
		}
	}
}
func NewBarber(name string) *Barber {
	b := &Barber{name: name}
	b.wakeMe = make(chan *Customer)
	return b
}
