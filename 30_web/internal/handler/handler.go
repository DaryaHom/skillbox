package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"network_communication/pkg/storage"
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

//GetAll - return all users from storage
func GetAll(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		elems := s.GetAll()
		if err := json.NewEncoder(w).Encode(elems); err != nil {
			log.Fatalln(err)
		}
	}
}

//CreateUser - create a new user & put it into storage
func CreateUser(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]interface{}{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		defer closeRequestBody(r.Body)

		name := fmt.Sprintf("%v", request["name"])
		age, err := strconv.Atoi(fmt.Sprintf("%v", request["age"]))
		if err != nil {
			log.Fatalln(err)
		}
		var friends []int
		if request["friends"] != nil { //protect from panic
			friends = convertSlice(request["friends"].([]interface{}))
		}

		userID := s.AddUser(name, age, friends)
		response := fmt.Sprintf("User %s with ID %d was created\n", name, userID)
		writeCreateResponse(w, response)
	}
}

//MakeFriends - reads user's id from req & put in to the slice of friend's id for each of two users
func MakeFriends(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		defer closeRequestBody(r.Body)

		id1, err := strconv.Atoi(request["source_id"])
		if err != nil {
			log.Fatalln(err)
		}
		id2, err := strconv.Atoi(request["target_id"])
		if err != nil {
			log.Fatalln(err)
		}

		s.MakeFriends(id1, id2)
		response := fmt.Sprintf("%s and %s are friends now\n", s.GetUser(id1).GetName(), s.GetUser(id2).GetName())
		writeOkResponse(w, response)
		return
	}
}

//DeleteUser - reads user's id from req & delete user from friends lists and from the store
func DeleteUser(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		defer closeRequestBody(r.Body)

		id, err := strconv.Atoi(request["target_id"])
		if err != nil {
			log.Fatalln(err)
		}

		response := fmt.Sprintf("User %s is deleted\n", s.GetUser(id).GetName())

		s.DeleteFromFriends(id)
		s.DeleteUser(id)
		writeOkResponse(w, response)
	}
}

//GetAllFriends - returns a list of friends of the user with id specified in req
func GetAllFriends(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, "user_id")
		if val != "" {
			userID, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalln(err)
			}
			friendsID := s.GetFriendsID(userID)
			response := fmt.Sprintf("Friends of user %d, %s: %v \n", userID, s.GetUser(userID).GetName(), friendsID)
			writeOkResponse(w, response)
			return
		}
		fmt.Fprintf(w, "Need valid user_id in URL")
	}
}

//UpdateAge - update user's age
func UpdateAge(s *storage.Storage) http.HandlerFunc {
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
			defer closeRequestBody(r.Body)

			age, err := strconv.Atoi(request["new age"])
			if err != nil {
				log.Fatalln(err)
			}

			s.UpdateAge(userID, age)
			response := fmt.Sprintf("User %d's age has been updated to %d\n", userID, age)
			writeOkResponse(w, response)
			return
		}
		fmt.Fprintf(w, "Need valid user_id in URL")
	}
}

func closeRequestBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

//writeOkResponse - returns 200-status
func writeOkResponse(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}

//writeCreateResponse - returns 201-status
func writeCreateResponse(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
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
