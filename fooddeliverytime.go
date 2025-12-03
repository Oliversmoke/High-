package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var item food

	if order == "burger" {
		item.preptime = 15
	} else if order == "chips" {
		item.preptime = 10
	} else if order == "nuggets" {
		item.preptime = 12
	} else {
		return 404
	}

	return item.preptime
}
