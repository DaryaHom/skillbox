package user

type User struct {
	name    string `json: "name"`
	age     int    `json: "age, omitempty"`
	friends []int  `json: "friends, omitempty"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetFriends(friends []int) {
	u.friends = friends
}

func (u *User) GetFriends() []int {
	return u.friends
}

//AddFriend - adds the user with friendID to the friends list if it's not already there
func (u *User) AddFriend(friendID int) {
	isFriends := u.checkFriendship(friendID)
	if !isFriends {
		u.friends = append(u.friends, friendID)
	}
}

func (u *User) DeleteFriend(id int) {
	for i := 0; i < len(u.friends); i++ {
		if u.friends[i] == id {
			if len(u.friends) != 0 && i < len(u.friends)-1 { // protect from panic
				u.friends = append(u.friends[:i], u.friends[i+1:]...)
				return
			}
			u.friends = u.friends[:len(u.friends)-1]
			return
		}
	}
}

//checkFriendship - checks if users are friends, returns true if yes
func (u *User) checkFriendship(id int) (isFriend bool) {
	for _, v := range u.friends {
		if v == id {
			isFriend = true
			return
		}
	}
	return
}
