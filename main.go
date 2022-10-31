package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var err error

const connectionString string = "root:testing123@tcp(localhost:3306)/go_todo"
const fileServerPath string = "C:/Programming/todoapp/frontend/build"

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func removeTodo(w http.ResponseWriter, r *http.Request) { // "api/todos/remove removes the todo with the specific ID from DB"
	w.Header().Set("Content-Type", "application/json")

	var re Todo
	err := json.NewDecoder(r.Body).Decode(&re)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err.Error())

	}

	delete_query := fmt.Sprintf("DELETE FROM todos WHERE id_todo = %d", re.Id)

	rows, err := db.Query(delete_query)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w)

}

func addTodo(w http.ResponseWriter, r *http.Request) { // "/api/todos/add adds a new todo in the database and returns it as a JSON response"

	w.Header().Set("Content-Type", "application/json")

	var re Todo
	err := json.NewDecoder(r.Body).Decode(&re)

	fmt.Println("Userid:", re.UserId)
	fmt.Println("title:", re.Title)

	if err != nil || len(re.Title) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	query := fmt.Sprintf("INSERT INTO todos (id_user, title, completed) VALUES (%d, '%s', %t) ", re.UserId, re.Title, false)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()

	var query_select = "SELECT MAX(id_todo) FROM todos;"

	res := db.QueryRow(query_select)

	if err != nil {
		panic(err.Error())
	}

	var id int

	err = res.Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	re.Id = id
	re.Completed = false

	fmt.Println(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(re)

}

func finishTodo(w http.ResponseWriter, r *http.Request) { // "/api/todos/done" ,  finds the Todo with the specific ID in DB and updates the row setting completed to true
	w.Header().Set("Content-Type", "application/json")

	var re Todo
	err = json.NewDecoder(r.Body).Decode(&re)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	query := fmt.Sprintf("UPDATE todos SET completed = %t WHERE id_todo = %d", re.Completed, re.Id)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w)

}

func getTodo(w http.ResponseWriter, r *http.Request) { // "/api/todos", returns a JSON response with all todos with a specific ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	query := r.URL.Query()

	_, present := query["id"]

	if !present {
		NotFoundHandler(w, r)
		return

	} else {
		id, _ := strconv.Atoi(r.FormValue("id"))
		json.NewEncoder(w).Encode(getTodoList(id))
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) { //custom 404
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Error 404")
}

func getTodoList(id int) []Todo { // gets todos from DB where id_todo = id and returns a list of Todos

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())

	}

	query := fmt.Sprintf("SELECT id_todo, id_user, title, completed FROM todos WHERE id_user = %d", id)

	res, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var toDoList []Todo
	for res.Next() {
		var todo Todo

		err := res.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Completed)

		if err != nil {
			log.Fatal(err)
		}

		toDoList = append(toDoList, todo)
	}

	return toDoList
}

func getUsers(w http.ResponseWriter, r *http.Request) { //gets users from DB and returns a JSON response with all users found
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())

	}

	var query = "SELECT id_user, username FROM users"

	res, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var userList []User
	for res.Next() {
		var user User

		err := res.Scan(&user.Id, &user.Username)

		if err != nil {
			log.Fatal(err)
		}

		userList = append(userList, user)
	}

	json.NewEncoder(w).Encode(userList)

}

func main() {

	// toDoList = append(toDoList, Todo{UserId: 1, Id: 1, Title: "delectus aut autem", Completed: false})
	// toDoList = append(toDoList, Todo{UserId: 2, Id: 2, Title: "quis ut nam facilis et officia qui", Completed: false})

	fileServer := http.FileServer(http.Dir(fileServerPath))
	http.Handle("/", fileServer)

	http.HandleFunc("/api/todos", getTodo)
	http.HandleFunc("/api/users", getUsers)
	http.HandleFunc("/api/todos/remove", removeTodo)
	http.HandleFunc("/api/todos/done", finishTodo)
	http.HandleFunc("/api/todos/add", addTodo)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
