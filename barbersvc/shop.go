package barbersvc

type Shop struct {
	barber      *Barber
	waitingRoom chan *Customer
}

func NewShop(barber *Barber, seats int) *Shop {
	shop := &Shop{barber: barber}
	shop.waitingRoom = make(chan *Customer, seats)
	return shop
}
