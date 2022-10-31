import React, {useEffect, useState} from 'react'
import {Button} from './styles/Button.styled'



export default function AddTask({userId, updateTodos}){

    const [titleInput, setTitleInput] = useState('');
    const [alert, setAlert] = useState(false);
    const [message, setMessage] = useState('')


    useEffect(() => {
        if(alert) {
          setTimeout(() => {
            setAlert(false);
          }, 1000)
        }
      }, [alert, handleSubmit])


    const handleSubmit = async (e) => {
        e.preventDefault()

        userId = parseInt(userId)
        var title = titleInput

        if (title.length == 0){

            setMessage("Title can not be empty!")
        }

        else{

            try{
                const response = await fetch(`/api/todos/add`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        title: title,
                        userId: userId
                        })
                    })
            if(!response.ok){
                setMessage("Error!")
                throw new Error(`Error! status: ${response.status}`)

            }

            const result = await response.json();

            var todo = result
            console.log(result)

            } catch (err){
                console.log(err.message);
            } finally {
                setTitleInput('')
                setAlert(true)
                setMessage("Task added!")
                updateTodos(todo)
            }
        }
    };

    return(
        <form>
            {alert && <h2>{message}</h2>}
            <br></br>
            <label>
                <p>Title</p>
                <input refs="title" type="text" onChange={(e) => setTitleInput(e.target.value)} value={titleInput} required/>
            </label>
            <br></br>
            <Button type="submit" value="Submit" onClick={handleSubmit}>Add Task</Button>
        </form>
    )

}