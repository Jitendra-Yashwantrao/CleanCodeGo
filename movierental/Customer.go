package movierental

import "fmt"

type Customer struct {
	name    string
	rentals Rentals
}

func (c Customer) Statement() string {

	result := "Rental Record for " + c.name + "\n"

	for _, rental := range c.rentals.rentals {
		result += "\t" + rental.movie.title + "\t" +
			fmt.Sprintf("%f", rental.amount()) + "\n"
	}
	result += "\tAmount owed is " + fmt.Sprintf("%f", c.rentals.totalAmountForRentals()) + "\n"
	result += "\tYou earned " + fmt.Sprintf("%v", c.rentals.totalRenterPointsFor()) + " frequent renter points"
	return result
}
