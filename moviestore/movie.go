package moviestore

import "fmt"

type FSK uint8
type Serial uint

const (
	FSK0  FSK = 0
	FSK6  FSK = 6
	FSK12 FSK = 12
	FSK16 FSK = 16
	FSK18 FSK = 18
)

type Movie struct {
	Title  string
	Fsk    FSK
	Serial Serial
}

func AllowedAtAge(m *Movie, age Age) bool {
	return uint8(m.Fsk) <= uint8(age)
}

func (movie *Movie) String() string {
	return fmt.Sprintf("%4d. %s (Ab %d)", movie.Serial, movie.Title, movie.Fsk)
}
