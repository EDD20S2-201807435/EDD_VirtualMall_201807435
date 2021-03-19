  
import React from 'react'
import Navbar from './Components/NavBar'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import UserList from './Components/UserList'
import CreateUser from './Components/CreateUser'
import ImportList from './Components/ImportList'
import ViewUser from './Components/ViewUser'


function App() {
  return (
        <Router>
          <Navbar />
          <Route path="/listado" component={UserList} />
          <Route path="/formulario" component={CreateUser} />
          <Route path="/listado2" component={ImportList} />
          <Route path="/view/:id" component={ViewUser} />
        </Router>
  )
}

export default App