# go-react-todo

to run server: /todoapp/ $ go run main.go

to build frontend: /todoapp/frontend/ $ npm run build

DB used was MySql. DB tables are users(id_user, username) and todos(id_todo, id_user, completed, title)

ConnectionString must be changed to fit your DB settings and the fileServerPath must be changed to /frontend/build (where index.html will be located)


Simple ToDo app which loads a list of tasks from a database and shows it on the webpage.

The app shows the state of the task ( completed or not ). Completed tasks can be removed from the list using the 'REMOVE' button and incompleted tasks can be marked as completed by using the 'DONE' button.

You can add another task by selecting a user and typing a title in the input field. The app creates a new task which is automatically marked as incompleted.

Users can filter the tasks by choosing a user.

known bugs: alert doesn't pop when the title input is left empty.
