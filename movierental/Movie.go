package movierental

const (
	Movie_CHILDRENS   int = 2
	Movie_REGULAR     int = 0
	Movie_NEW_RELEASE int = 1
)

type Movie struct {
	title     string
	priceCode int
}
