package servidor

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("fail body format"))
		return
	}

	var user user
	user.ID = uuid.NewString()

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error to convert json in struct"))
		return
	}
	db, err := db.OpenConn()

	if err != nil {
		w.Write([]byte("fail to connect database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (id, name, email) values ($1, $2, $3)")
	if err != nil {
		w.Write([]byte("fail to create statement user"))
		return
	}

	defer statement.Close()
	insert, err := statement.Exec(user.ID, user.Name, user.Email)
	if err != nil {
		w.Write([]byte("fail to exec statement"))
		return
	}

	rows, err := insert.RowsAffected()
	if err != nil {
		w.Write([]byte("fail to return id"))
		return
	}

	if rows >= 1 {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("user created ID: %s", user.ID)))
	}
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConn()

	if err != nil {
		w.Write([]byte("fail to connect database"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Error to search users on database"))
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error to scan user"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error to convert users to json"))
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	db, err := db.OpenConn()

	if err != nil {
		w.Write([]byte("fail to connect database"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from users where id = $1", ID)
	if err != nil {
		w.Write([]byte("error to search user"))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("erro to scan user"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error to convert users to json"))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("fail body format"))
		return
	}

	var user user
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("erro to json to struct"))
		return
	}

	db, err := db.OpenConn()

	if err != nil {
		w.Write([]byte("fail to connect database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = $1, email = $2 where id = $3")
	if err != nil {
		w.Write([]byte("erro to prepare statement"))
		return
	}

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("error to update user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	db, err := db.OpenConn()

	if err != nil {
		w.Write([]byte("fail to connect database"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("delete from users where id = $1")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("erro to prepare statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("erro to delete user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
