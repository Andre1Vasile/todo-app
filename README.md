# go-react-todo

to run server: /todoapp/ $ go run main.go

to build frontend: /todoapp/frontend/ $ npm run build

DB used was MySql. DB tables are users(id_user, username) and todos(id_todo, id_user, completed, title)

ConnectionString must be changed to fit your DB settings and the fileServerPath must be changed to /frontend/build (where index.html will be located)
