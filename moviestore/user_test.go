package moviestore

import "testing"

func TestUser(t *testing.T) {
	cases := []struct {
		in   User
		want string
	}{
		{User{"Heinz Ketchup", 12, 12}, "  12. Heinz Ketchup, 12"},
		{User{"Hendlmaier Senf", 13, 8}, "   8. Hendlmaier Senf, 13"},
		{User{"Löwensenf", 14, 13}, "  13. Löwensenf, 14"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("%q.String() == %q, want %q", c.in, got, c.want)
		}
	}
}
