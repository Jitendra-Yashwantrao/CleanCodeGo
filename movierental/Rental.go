package movierental

type Rental struct {
	movie      Movie
	daysRented int
}

func (r Rental) getDaysRented() int {
	return r.daysRented
}

func (r Rental) getMovie() Movie {
	return r.movie
}

func (r Rental) amount() float64 {
	var amount float64
	switch r.getMovie().getPriceCode() {
	case Movie_REGULAR:
		amount += 2
		if r.getDaysRented() > 2 {
			amount += (float64(r.getDaysRented()) - 2) * 1.5
		}
		break
	case Movie_NEW_RELEASE:
		amount += float64(r.getDaysRented()) * 3
		break
	case Movie_CHILDRENS:
		amount += 1.5
		if r.getDaysRented() > 3 {
			amount += (float64(r.getDaysRented()) - 3) * 1.5

		}
		break
	}
	return amount
}

func (r Rental) renterPointFor() int {
	frequentRenterPoints := 1
	if r.getMovie().getPriceCode() == Movie_NEW_RELEASE && r.getDaysRented() > 1 {
		frequentRenterPoints++
	}
	return frequentRenterPoints
}
