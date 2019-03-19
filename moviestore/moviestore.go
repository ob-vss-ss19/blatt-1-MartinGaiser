package moviestore

type moviestoreImpl struct {
	available  map[Serial]Movie
	rented     map[UserID][]Movie
	users      map[UserID]User
	nextSerial Serial
	nextUserID UserID
}

func (ms *moviestoreImpl) AddMovie(title string, fsk FSK) Serial {
	nextSerial := ms.nextSerial
	ms.nextSerial++
	ms.available[nextSerial] = Movie{title,fsk,nextSerial}
	return nextSerial
}

func (ms *moviestoreImpl) AddUser(name string, age Age) UserID {
	nextUserID := ms.nextUserID
	ms.nextUserID++
	ms.users[nextUserID] = User{name,age,nextUserID}
	return nextUserID
}

func (ms *moviestoreImpl) Rent(serial Serial, userID UserID) (User, Movie, error) {
	return User{}, Movie{}, nil
}

func (ms *moviestoreImpl) RentedByUser(userID UserID) ([]Movie, error) {
	return []Movie{{}, {}}, nil
}

func (ms *moviestoreImpl) Return(serial Serial) (User, Movie, error) {
	return User{}, Movie{}, nil
}
