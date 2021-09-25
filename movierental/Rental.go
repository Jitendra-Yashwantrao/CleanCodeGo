package movierental

type Rental struct {
	movie      Movie
	daysRented int
}

func (r Rental) amount() float64 {
	var amount float64
	switch r.movie.priceCode {
	case Movie_REGULAR:
		amount += 2
		if r.daysRented > 2 {
			amount += (float64(r.daysRented) - 2) * 1.5
		}
		break
	case Movie_NEW_RELEASE:
		amount += float64(r.daysRented) * 3
		break
	case Movie_CHILDRENS:
		amount += 1.5
		if r.daysRented > 3 {
			amount += (float64(r.daysRented) - 3) * 1.5

		}
		break
	}
	return amount
}

func (r Rental) renterPointFor() int {
	frequentRenterPoints := 1
	if r.movie.priceCode == Movie_NEW_RELEASE && r.daysRented > 1 {
		frequentRenterPoints++
	}
	return frequentRenterPoints
}
