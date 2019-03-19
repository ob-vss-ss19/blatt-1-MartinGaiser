package moviestore

type Moviestore interface {
	AddMovie(string, FSK) Serial
	AddUser(string, Age) UserID
	Rent(Serial, UserID) (User, Movie, error)
	Return(Serial) (User, Movie, error)
	RentedByUser(UserID) ([]Movie, error)
}

type moviestoreImpl struct {
	available map[Serial] Movie
	rented map[UserID][] Movie
	users map[UserID] User
	nextSerial Serial
	nextUserID UserID
}


func NewMoviestore() Moviestore{
	return nil
}

func AllowedAtAge(m *Movie, age Age) bool{
	return false
}

func (ms *moviestoreImpl) AddMovie(title string, fsk FSK) Serial{
	return 0
}

func (ms *moviestoreImpl) AddUser(name string, age Age) UserID {
	return 0
}

func (ms *moviestoreImpl) Rent(serial Serial, userID UserID) (User, Movie, error){
	return User{"Null",0,0},Movie{"Null",FSK16,0},nil
}

func (ms *moviestoreImpl) RentedByUser(userID UserID) ([]Movie, error){
	return []Movie{{"Null", FSK16, 0},{"Null",FSK16,0}},nil
}

func (ms *moviestoreImpl) Return(serial Serial) (User, Movie, error){
	return User{"Null",0,0},Movie{"Null",FSK16,0},nil
}


