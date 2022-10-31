import React, {useEffect, useState} from 'react'
import ToDoList from './ToDoList'
import AddTask from './AddTask'
import {ContentSection} from './styles/ContentSection.styled'





export default function UserList({userList}){
    const [selected, setSelected] = useState(userList[0].value);
    const [todos, setTodos] = useState([])

    function removeTodo(taskId){
        const newTodos = todos.filter(todo => todo.id != taskId)
        setTodos(newTodos)
    }

    function updateTodos(todo){
        setTodos(todos.concat(todo))
    }

    function handleChange(event) {
        setSelected(event.target.value);
        var endpoint_url = `/api/todos?id=` + event.target.value.toString()
        fetchTodos(endpoint_url)
    } 
    

    const fetchTodos = async (url) => {
        try{
            var response = await fetch(url)

            var json = await response.json();
            setTodos(json)
            setSelected(json[0].userId)
        }
        catch (error){
            console.log("error: ", error)
        }
    }
    useEffect(() => {
        fetchTodos(`/api/todos?id=` + userList[0].id)
    }, [])


    return(
        <ContentSection>
            { todos && <AddTask userId = {parseInt(selected)} updateTodos = {updateTodos}/>}
            <br></br>
            <label for="users">Choose a user:</label>
            <select name="users" id="users" value={selected} onChange={handleChange}>
                {userList.map(user => (
                    <option key={user.id} value={user.id}>{user.username}</option>
                ))}
            </select>
            {todos && <ToDoList toDoList = {todos} removeTodo = {removeTodo}/>}
        </ContentSection>
    )

}