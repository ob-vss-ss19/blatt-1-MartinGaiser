package moviestore

import "testing"

func TestMovie(t *testing.T) {
	cases := []struct {
		in          Movie
		want        string
		age         Age
		rentAllowed bool
	}{
		{Movie{"Am Limit", FSK6, 12}, "  12. Am Limit (Ab 6)", 5, false},
		{Movie{"Texas Chainsaw Massacre", FSK18, 8}, "   8. Texas Chainsaw Massacre (Ab 18)", 18, true},
		{Movie{"Inglourious Basterds", FSK16, 13}, "  13. Inglourious Basterds (Ab 16)", 17, true},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("%q.String() == %q, want %q", c.in, got, c.want)
		}
	}

	for _, c := range cases {
		got := AllowedAtAge(&c.in, c.age)
		if got != c.rentAllowed {
			t.Errorf("%s: AgeAllowed == %t, want %t", c.in.Title, got, c.rentAllowed)
		}
	}
}
