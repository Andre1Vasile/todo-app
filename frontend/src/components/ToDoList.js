import React, {useState, useEffect} from 'react'
import Todo from './Todo'
import AddTask from './AddTask'




export default function ToDoList({toDoList, removeTodo}){



    return(
        toDoList.map(todo => {
            return(
                <ul className="toDoList">
                    <Todo key={todo.id} todo={todo} removeTodo={removeTodo}  />
                </ul>
            )
        })
    )
}