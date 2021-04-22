import React from 'react'
import Navbar from './Components/NavBar'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import UserList from './Components/UserList'
import CreateUser from './Components/CreateUser'
import ImportList from './Components/ImportList'
import Departamentos from './Components/Departamentos'
import Tiendas from './Components/Tiendas'
import Calificaciones from './Components/Calificaciones'
import Tienda from './Components/Tienda'
import Principal from './Components/Principal'
import Login from './Login'


function App() {
  return (
    <>
        <Router>
          
          <Route path="/listado" component={UserList} />
          <Route path="/Login" component={Login} />
          <Route path="/Principal" component={Principal} />
          <Route path="/formulario" component={CreateUser} />
          <Route path="/listado2" component={ImportList} />
          <Route path="/depa/:id" component={Departamentos} />
          <Route path="/Tiendas/:id1/:id2/:id3" component={Tiendas} />
          <Route path="/Tienda/:id1/:id2/:id3" component={Tienda} />
          <Route path="/Calificaciones/:id1/:id2" component={Calificaciones} />
        </Router>
        </>
  )
}

export default App