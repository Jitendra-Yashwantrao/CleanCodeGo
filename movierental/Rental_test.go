package movierental

import (
	"testing"
)

func TestRental_renterPointFor(t *testing.T) {
	type fieldsRental struct {
		movie      Movie
		daysRented int
	}
	tests := []struct {
		name   string
		fields fieldsRental
		want   int
	}{
		{
			name: "Should return 2 renter point for movie of type new release when days rented greater than 1 ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_NEW_RELEASE,
				},
				daysRented: int(2)},
			want: 2,
		},
		{
			name: "Should return 1 renter point for movie of type new release when days rented is 1 ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_NEW_RELEASE,
				},
				daysRented: int(1)},
			want: 1,
		},
		{
			name: "Should return 1 renter point for movie of type other than new release when days rented greater than 1 ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_CHILDRENS,
				},
				daysRented: int(2)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rental{
				movie:      tt.fields.movie,
				daysRented: tt.fields.daysRented,
			}
			if got := r.renterPointFor(); got != tt.want {
				t.Errorf("Rental.renterPointFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRental_amount(t *testing.T) {
	type fieldsRental struct {
		movie      Movie
		daysRented int
	}
	tests := []struct {
		name   string
		fields fieldsRental
		want   float64
	}{
		{
			name: "Should return amount 2 for regular movie rented for 2 or less days ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_REGULAR,
				},
				daysRented: int(2)},
			want: 2,
		},
		{
			name: "Should return amount 5 for regular movie rented for 4  days ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_REGULAR,
				},
				daysRented: int(4)},
			want: 5,
		},
		{
			name: "Should return amount 6 for New release movie rented for 2  days ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_NEW_RELEASE,
				},
				daysRented: int(2)},
			want: 6,
		},
		{
			name: "Should return amount 1.5 for children movie rented for 3 or less days ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_CHILDRENS,
				},
				daysRented: int(3)},
			want: 1.5,
		},
		{
			name: "Should return amount 3 for children movie rented for 4 days ",
			fields: fieldsRental{
				movie: Movie{
					title:     "Happy Minds",
					priceCode: Movie_CHILDRENS,
				},
				daysRented: int(4)},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rental{
				movie:      tt.fields.movie,
				daysRented: tt.fields.daysRented,
			}
			if got := r.amount(); got != tt.want {
				t.Errorf("Rental.amount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRentals_totalRenterPointsFor(t *testing.T) {
	type fieldsRental struct {
		rentals []Rental
	}
	tests := []struct {
		name   string
		fields fieldsRental
		want   int
	}{
		{
			name: "Should return total 3 renter points for two rentals listed",
			fields: fieldsRental{rentals: []Rental{
				{
					movie: Movie{
						title:     "Happy Minds",
						priceCode: Movie_NEW_RELEASE,
					},
					daysRented: int(2),
				},
				{
					movie: Movie{
						title:     "Happy Minds",
						priceCode: Movie_NEW_RELEASE,
					},
					daysRented: int(1),
				},
			},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := Rentals{
				rentals: tt.fields.rentals,
			}
			if got := rs.totalRenterPointsFor(); got != tt.want {
				t.Errorf("Rentals.totalRenterPointsFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRentals_totalAmountForRentals(t *testing.T) {
	type fieldsRental struct {
		rentals []Rental
	}
	tests := []struct {
		name   string
		fields fieldsRental
		want   float64
	}{
		{
			name: "Should return total amount 9  for two rentals listed",
			fields: fieldsRental{rentals: []Rental{
				{
					movie: Movie{
						title:     "Happy Minds",
						priceCode: Movie_NEW_RELEASE,
					},
					daysRented: int(2),
				},
				{
					movie: Movie{
						title:     "Happy Minds",
						priceCode: Movie_NEW_RELEASE,
					},
					daysRented: int(1),
				},
			},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := Rentals{
				rentals: tt.fields.rentals,
			}
			if got := rs.totalAmountForRentals(); got != tt.want {
				t.Errorf("Rentals.totalAmountForRentals() = %v, want %v", got, tt.want)
			}
		})
	}
}
