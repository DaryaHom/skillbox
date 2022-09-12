package user

import (
	"database/sql"
	"log"
)

type User struct {
	id      int    `json: "id"`
	name    string `json: "name"`
	age     int    `json: "age, omitempty"`
	friends []int  `json: "friends, omitempty"`
}

func NewUser() *User {
	return &User{}
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

//AddUserParams - add new user parameters to database
func (u *User) AddUserParams(db *sql.DB) (int, error) {
	result, err := db.Exec("insert into users(name, age) values ($1, $2)", u.name, u.age)
	if err != nil {
		return -1, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	u.id = int(lastInsertId)

	for _, friendID := range u.friends {
		err := MakeFriends(db, u.id, friendID)
		if err != nil {
			log.Println(err)
		}
		err = MakeFriends(db, friendID, u.id)
		if err != nil {
			log.Println(err)
		}
	}
	return int(lastInsertId), nil
}

func GetName(db *sql.DB, id int) string {
	var name string
	if err := db.QueryRow("select name from users where id = $1", id).Scan(&name); err != nil {
		log.Fatalln(err)
	}
	return name
}

func UpdateAge(db *sql.DB, id, age int) error {
	_, err := db.Exec("update users set age = $1 where id = $2", age, id)
	if err != nil {
		return err
	}
	return nil
}

//GetFriendsID - returns a list of id friends of the specified user
func GetFriendsID(db *sql.DB, id int) ([]int, error) {
	rows, err := db.Query("select friendID from friends where id = $1", id)
	if err != nil {
		return nil, err
	}

	var friends []int
	for rows.Next() {
		var f int
		if err := rows.Scan(&f); err != nil {
			return nil, err
		}
		friends = append(friends, f)
	}
	return friends, nil
}

//MakeFriends - adds the user with friendID to the friends list if it's not already there
func MakeFriends(db *sql.DB, id1, id2 int) error {
	_, err := db.Exec("insert into friends(id, friendID) values ($1, $2)", id1, id2)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromDB(db *sql.DB, id int) error {
	_, err := db.Exec("delete from friends where id = $1 or friendID = $1", id)
	if err != nil {
		return err
	}

	_, err = db.Exec("delete from users where id = $1", id)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
