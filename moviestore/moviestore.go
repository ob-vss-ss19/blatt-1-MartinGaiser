package moviestore

import "errors"

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
	ms.available[nextSerial] = Movie{title, fsk, nextSerial}
	return nextSerial
}

func (ms *moviestoreImpl) AddUser(name string, age Age) UserID {
	nextUserID := ms.nextUserID
	ms.nextUserID++
	ms.users[nextUserID] = User{name, age, nextUserID}
	ms.rented[nextUserID] = []Movie{}
	return nextUserID
}

func (ms *moviestoreImpl) Rent(serial Serial, userID UserID) (User, Movie, error) {
	movie, isAvailable := ms.available[serial]
	user, userExists := ms.users[userID]
	if userExists {
		if isAvailable {
			delete(ms.available, serial)
			ms.rented[userID] = append(ms.rented[userID], movie)
			return user, movie, nil
		}
		return User{}, Movie{}, errors.New("the Movie with the Serial: " + string(serial) + " was not found")
	}
	return User{}, Movie{}, errors.New("no User with the ID: " + string(userID) + " was found")
}

func (ms *moviestoreImpl) RentedByUser(userID UserID) ([]Movie, error) {
	rentedMovies, userExists := ms.rented[userID]
	if userExists {
		return rentedMovies, nil
	}
	return []Movie{}, errors.New("the User with the UserID: " + string(userID) + " does not exist")

}

func (ms *moviestoreImpl) Return(serial Serial) (User, Movie, error) {
	_, movieAvailable := ms.available[serial]
	if movieAvailable {
		return User{}, Movie{}, errors.New("the Movie can't be returned since it's already available")
	}
	var returnMovie Movie
	var renterUserID UserID
	for user, rentedMovies := range ms.rented {
		newRentedMovies := rentedMovies[:0]
		for _, movie := range rentedMovies {
			if movie.Serial != serial {
				newRentedMovies = append(newRentedMovies, movie)
			} else {
				returnMovie = movie
				renterUserID = user
			}
		}
		ms.rented[user] = newRentedMovies
	}
	if returnMovie.Serial != serial {
		return User{}, Movie{}, errors.New("the Movie can't be returned since it it not rented")
	}
	return ms.users[renterUserID], returnMovie, nil
}
