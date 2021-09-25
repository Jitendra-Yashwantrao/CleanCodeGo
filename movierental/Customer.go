package movierental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func (c Customer) addRental(arg Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) getName() string {
	return c.name
}

func (c Customer) Statement() string {

	var totalAmount float64
	frequentRenterPoints := 0
	result := "Rental Record for " + c.getName() + "\n"
	for _, rental := range c.rentals {
		frequentRenterPoints += renterPointFor(rental)
	}
	for _, rental := range c.rentals {
		thisAmount := amountFor(rental)
		result += "\t" + rental.getMovie().getTitle() + "\t" +
			fmt.Sprintf("%f", thisAmount) + "\n"
		totalAmount += thisAmount
	}

	result += "\tAmount owed is " + fmt.Sprintf("%f", totalAmount) + "\n"
	result += "\tYou earned " + fmt.Sprintf("%v", frequentRenterPoints) + " frequent renter points"
	return result
}

func renterPointFor(rental Rental) int {
	frequentRenterPoints := 1
	if rental.getMovie().getPriceCode() == Movie_NEW_RELEASE && rental.getDaysRented() > 1 {
		frequentRenterPoints++
	}
	return frequentRenterPoints
}

func amountFor(rental Rental) float64 {
	var amount float64
	switch rental.getMovie().getPriceCode() {
	case Movie_REGULAR:
		amount += 2
		if rental.getDaysRented() > 2 {
			amount += (float64(rental.getDaysRented()) - 2) * 1.5
		}
		break
	case Movie_NEW_RELEASE:
		amount += float64(rental.getDaysRented()) * 3
		break
	case Movie_CHILDRENS:
		amount += 1.5
		if rental.getDaysRented() > 3 {
			amount += (float64(rental.getDaysRented()) - 3) * 1.5

		}
		break
	}
	return amount
}
