package movierental


type Rental struct {
	movie Movie
	daysRented int
}

func (r Rental) getDaysRented() int {
	return r.daysRented
}

func (r Rental) getMovie() Movie {
	return  r.movie
}