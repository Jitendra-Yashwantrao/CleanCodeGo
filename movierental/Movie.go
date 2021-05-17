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

func (m Movie) getPriceCode() int {
	return m.priceCode
}

func (m Movie) getTitle() string {
	return m.title
}

func (m Movie) setPriceCode(code int) {

	m.priceCode = code
}
