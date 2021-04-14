  
import React from 'react'
import Navbar from './Components/NavBar'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import UserList from './Components/UserList'
import CreateUser from './Components/CreateUser'
import ImportList from './Components/ImportList'
import Departamentos from './Components/Departamentos'
import Tiendas from './Components/Tiendas'


function App() {
  return (

        <Router>
          <Navbar />
          <Route path="/listado" component={UserList} />
          <Route path="/formulario" component={CreateUser} />
          <Route path="/listado2" component={ImportList} />
          <Route path="/depa/:id" component={Departamentos} />
          <Route path="/Tiendas/:id1/:id2" component={Tiendas} />
        </Router>
  )
}

export default App