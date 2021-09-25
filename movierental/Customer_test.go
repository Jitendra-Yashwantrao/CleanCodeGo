package movierental

import (
	"fmt"
	"testing"
)

func TestCustomer_Statement(t *testing.T) {
	type CustomerTest struct {
		name    string
		rentals []Rental
	}

	tests := []struct {
		name   string
		fields CustomerTest
		want   string
	}{

		{name: "Should return statment for customer of ",
			fields: CustomerTest{name: "Jitendra", rentals: []Rental{
				{daysRented: 7, movie: Movie{title: "IceAge", priceCode: Movie_CHILDRENS}},
				{daysRented: 6, movie: Movie{title: "Temrminator", priceCode: Movie_REGULAR}},
				{daysRented: 5, movie: Movie{title: "Happy Minds", priceCode: Movie_NEW_RELEASE}}}},
			want: `Rental Record for Jitendra
	IceAge	7.500000
	Temrminator	8.000000
	Happy Minds	15.000000
	Amount owed is 30.500000
	You earned 4 frequent renter points`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Customer{
				name:    tt.fields.name,
				rentals: tt.fields.rentals,
			}
			got := c.Statement()
			fmt.Println("got=", got)
			if got != tt.want {
				t.Errorf("Customer.Statement() got =%v= want =%v=", got, tt.want)
			}
		})
	}
}
