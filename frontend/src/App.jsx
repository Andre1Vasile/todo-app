import React, {useEffect, useState} from 'react'
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import UserList from './components/UserList'
import {Container} from './components/styles/Container.styled'


function App() {

  const [users, setUsers] = useState(null)
  const fetchUsers = async () => {
      try{
          var response1 = await fetch(`/api/users`)
          var json1 = await response1.json();
          setUsers(json1)
      }
      catch (error){
          console.log("error: ", error)
      }
  }
  useEffect(() => {
      fetchUsers()
  }, [])


  return (
        <Container>
            {users && <UserList userList = {users}/>}
        </Container>
  )
}

export default App;
