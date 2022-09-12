package storage

import (
	"network_communication/pkg/user"
)

var counter int

// Storage - user repository
type Storage struct {
	store map[int]*user.User
}

//NewStorage - creates new user repository
func NewStorage() *Storage {
	return &Storage{
		make(map[int]*user.User),
	}
}

//AddUser - add new user to the storage
func (s *Storage) AddUser(name string, age int, friends []int) int {
	if name != "" { //username can't be empty
		defer counterIncrement() //increments the ID counter after creating a new user

		newUser := user.NewUser()
		newUser.SetName(name)
		newUser.SetAge(age)
		newUser.SetFriends(friends)

		s.store[counter] = newUser
	}
	return counter
}

func (s *Storage) MakeFriends(id1, id2 int) {
	s.store[id1].AddFriend(id2)
	s.store[id2].AddFriend(id1)
}

//GetFriendsID - returns a list of friends of the user with the specified id
func (s *Storage) GetFriendsID(id int) []int {
	friendsID := s.store[id].GetFriends()
	return friendsID
}

//DeleteFromFriends - removes the user with the specified id from the friend lists of other users in storage
func (s *Storage) DeleteFromFriends(id int) {
	for userID, user := range s.store {
		if userID != id {
			user.DeleteFriend(id)
		}
	}
}

//DeleteUser - delete user with specified id from storage
//used after the command DeleteFromFriends
func (s *Storage) DeleteUser(id int) {
	delete(s.store, id)
}

//GetAll - returns all users from store (without user's private fields)
func (s *Storage) GetAll() map[int]*user.User {
	return s.store
}

//GetUser - returns user with specified id from store (without user's private fields)
func (s *Storage) GetUser(id int) *user.User {
	return s.store[id]
}

//UpdateAge - update user's age
func (s *Storage) UpdateAge(id, age int) {
	s.store[id].SetAge(age)
}

func (s *Storage) GetCount() int {
	return counter
}

func counterIncrement() {
	counter++
}
