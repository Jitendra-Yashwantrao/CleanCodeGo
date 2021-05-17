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
	for _, each := range c.rentals {
		var thisAmount float64
		//determine amounts for each line
		switch each.getMovie().getPriceCode() {
		case Movie_REGULAR:
			thisAmount += 2
			if each.getDaysRented() > 2 {
				thisAmount += (float64(each.getDaysRented()) - 2) * 1.5
			}
			break
		case Movie_NEW_RELEASE:
			thisAmount += float64(each.getDaysRented()) * 3
			break
		case Movie_CHILDRENS:
			thisAmount += 1.5
			if each.getDaysRented() > 3 {
				thisAmount += (float64(each.getDaysRented()) - 3) * 1.5

			}
			break
		}
		// add frequent renter points
		frequentRenterPoints++
		// add bonus for a two day new release rental
		if each.getMovie().getPriceCode() == Movie_NEW_RELEASE && each.getDaysRented() > 1 {
			frequentRenterPoints++
		}

		//show figures for this rental
		result += "\t" + each.getMovie().getTitle() + "\t" +
			fmt.Sprintf("%f",thisAmount) + "\n"
		totalAmount += thisAmount
	}
	//add footer lines result
	result += "Amount owed is " + fmt.Sprintf("%f",totalAmount) + "\n"
	result += "You earned " + string(frequentRenterPoints) + " frequent renter points"
	return result
}
