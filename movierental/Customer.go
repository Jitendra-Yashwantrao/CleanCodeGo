package movierental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func (c Customer) Statement() string {

	var totalAmount float64
	result := "Rental Record for " + c.name + "\n"

	for _, rental := range c.rentals {
		result += "\t" + rental.movie.title + "\t" +
			fmt.Sprintf("%f", rental.amount()) + "\n"
	}

	totalAmount = totalAmountForRentals(c.rentals)

	result += "\tAmount owed is " + fmt.Sprintf("%f", totalAmount) + "\n"
	result += "\tYou earned " + fmt.Sprintf("%v", totalRenterPointsFor(c.rentals)) + " frequent renter points"
	return result
}

func totalRenterPointsFor(rentals []Rental) int {
	frequentRenterPoints := 0
	for _, rental := range rentals {
		frequentRenterPoints += rental.renterPointFor()
	}
	return frequentRenterPoints
}

func totalAmountForRentals(rentals []Rental) float64 {
	var totalAmount float64
	for _, rental := range rentals {
		totalAmount += rental.amount()
	}
	return totalAmount
}
