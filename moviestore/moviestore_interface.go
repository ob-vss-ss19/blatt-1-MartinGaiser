package moviestore

type Moviestore interface {
	AddMovie(string, FSK) Serial
	AddUser(string, Age) UserID
	Rent(Serial, UserID) (User, Movie, error)
	Return(Serial) (User, Movie, error)
	RentedByUser(UserID) ([]Movie, error)
}

func NewMoviestore() Moviestore {
	return nil
}

