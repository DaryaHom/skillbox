package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proxy/internal/req"
	"proxy/internal/user"
	"strconv"

	"github.com/go-chi/chi"
)

//Get - a hello function
func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Welcome")); err != nil {
			log.Fatalln(err)
		}
	}
}

//CreateUser - creates a new user & writes it to the database
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]interface{}{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		defer req.CloseRequestBody(r.Body)

		name := fmt.Sprintf("%v", request["name"])
		age, err := strconv.Atoi(fmt.Sprintf("%v", request["age"]))
		if err != nil {
			log.Fatalln(err)
		}
		var friends []int
		if request["friends"] != nil { //protect from panic
			friends = convertSlice(request["friends"].([]interface{}))
		}

		u := user.NewUser()
		u.SetName(name)
		u.SetAge(age)
		u.SetFriends(friends)

		userID, err := u.AddUserParams(db)
		if err != nil {
			response := fmt.Sprintf("%v", err)
			req.WriteBadRequest(w, response)
			return
		}
		response := fmt.Sprintf("User %s with ID %d was created\n", name, userID)
		req.WriteCreateResponse(w, response)
	}
}

//MakeFriends - reads user's id from req & put it to the database table for each of two users
func MakeFriends(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		defer req.CloseRequestBody(r.Body)

		id1, err := strconv.Atoi(request["source_id"])
		if err != nil {
			log.Fatalln(err)
		}
		id2, err := strconv.Atoi(request["target_id"])
		if err != nil {
			log.Fatalln(err)
		}

		err = user.MakeFriends(db, id1, id2)
		if err != nil {
			response := fmt.Sprintf("%v", err)
			req.WriteBadRequest(w, response)
			return
		}

		err = user.MakeFriends(db, id2, id1)
		if err != nil {
			response := fmt.Sprintf("%v", err)
			req.WriteBadRequest(w, response)
			return
		}

		response := fmt.Sprintf("%s and %s are friends now\n", user.GetName(db, id1), user.GetName(db, id2))
		req.WriteOkResponse(w, response)
	}
}

//DeleteUser - reads user's id from req & delete user from the database
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		defer req.CloseRequestBody(r.Body)

		id, err := strconv.Atoi(request["target_id"])
		if err != nil {
			log.Fatalln(err)
		}

		err = user.DeleteFromDB(db, id)
		if err != nil {
			response := fmt.Sprintf("%v", err)
			req.WriteBadRequest(w, response)
			return
		}
		response := fmt.Sprintf("User %d is deleted\n", id)
		req.WriteOkResponse(w, response)
	}
}

//GetAllFriends - returns a list of friends of the user with id specified in req
func GetAllFriends(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, "user_id")
		if val != "" {
			userID, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalln(err)
			}
			friendsID, err := user.GetFriendsID(db, userID)
			if err != nil {
				response := fmt.Sprintf("%v", err)
				req.WriteBadRequest(w, response)
				return
			}
			response := fmt.Sprintf("Friends of user %d: %v \n", userID, friendsID)
			req.WriteOkResponse(w, response)
		}
	}
}

//UpdateAge - update user's age in the database
func UpdateAge(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, "user_id")
		if val != "" {
			userID, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalln(err)
			}

			request := map[string]string{}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				log.Fatalln(err)
			}
			defer req.CloseRequestBody(r.Body)

			age, err := strconv.Atoi(request["new age"])
			if err != nil {
				log.Fatalln(err)
			}

			err = user.UpdateAge(db, userID, age)
			if err != nil {
				response := fmt.Sprintf("%v", err)
				req.WriteBadRequest(w, response)
				return
			}
			response := fmt.Sprintf("User %d's age has been updated to %d\n", userID, age)
			req.WriteOkResponse(w, response)
		}
	}
}

//convertSlice - makes slice of integers from interface with slice of integers
func convertSlice(in []interface{}) (out []int) {
	b := make([]int, len(in))
	for i := range in {
		b[i] = in[i].(int)
	}
	return
}
