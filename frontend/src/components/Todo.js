import React, { useEffect, useState } from 'react'
import {ListItem} from './styles/ListItem.styled'
import {Button} from './styles/Button.styled'



export default function Todo({todo, removeTodo}){
    const [alert, setAlert] = useState(false);

    useEffect(() => {
        if(alert) {
          setTimeout(() => {
            setAlert(false);
            removeTodo(todo.id)
          }, 2000)
        }
      }, [alert, handleRemove])

      const [visible, setVisible] = useState(true);

      const removeElement = () => {
        setVisible((prev) => !prev);
      };


    const handleRemove = async (e) => {
        e.preventDefault()

        try{
            const response = await fetch(`/api/todos/remove`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    id: todo.id,
                    userId: todo.userId,
                    })
                })

        if(!response.ok){
            throw new Error(`Error! status: ${response.status}`)
        }

        const result = await response.json();
        console.log(result)
        } catch (err){
            console.log(err.message);
        } finally {
            setAlert(true)
        }
    };


    const handleDone = async (e) => {
        e.preventDefault()

        try{
            console.log("id: " + todo.id)
            console.log("completed: " + todo.completed)
            const response = await fetch(`/api/todos/done`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    id: todo.id,
                    completed: true,
                    })
                })

        if(!response.ok){
            throw new Error(`Error! status: ${response.status}`)
        }

        const result = await response.json();
        console.log(result)
        } catch (err){
            console.log(err.message);
        } finally {
            todo.completed = true
            removeElement(this)
        }
    };

    var button
    if (todo.completed){
        button = <Button  className="remove" onClick={handleRemove}>REMOVE</Button>
    }
    else{
        button = <Button primary className="done" onClick={handleDone}>DONE</Button>
    }


    return(
        <ListItem>
            <li className="todo">
                <h1 className="title">{todo.title}</h1>
                <h3 className="userId"> User ID: {todo.userId}</h3>
                <h3 className="id">Task ID: {todo.id}</h3>
                <input type="checkbox" className="completed" checked={todo.completed}></input>
                {button}
                {alert && <h2>Task removed</h2>}
            </li>
        </ListItem>
    )
}