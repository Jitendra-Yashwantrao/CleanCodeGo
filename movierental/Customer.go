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

func (c Customer) HTMLStatement() string {

	result := "<h1>Rental Record for <em>" + c.name + "</em></h1>\n\t\t\t<table>\n"

	for _, rental := range c.rentals.rentals {
		result += "\t\t\t  <tr><td>" + rental.movie.title + "</td><td>" +
			fmt.Sprintf("%f", rental.amount()) + "</td></tr>\n"
	}
	result += "\t\t\t</table>\n\t\t\t<p>Amount owed is <em>" + fmt.Sprintf("%f", c.rentals.totalAmountForRentals()) + "</em></p>\n"
	result += "\t\t\t<p>You earned <em>" + fmt.Sprintf("%v", c.rentals.totalRenterPointsFor()) + "</em> frequent renter points</p>"
	return result
}
