  
import React from 'react'
import Navbar from './NavBar'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import UserList from './UserList'
import CreateUser from './CreateUser'
import ImportList from './ImportList'
import Departamentos from './Departamentos'
import Tiendas from './Tiendas'
import Calificaciones from './Calificaciones'
import Tienda from './Tienda'
import './Css/app.css'



function Principal() {
  return (

    <div  id="containerLogin">
    <img src="https://www.iconpacks.net/icons/1/free-user-login-icon-305-thumb.png"></img>
    <center><h4>Ingrese Su numero de DPI</h4>
    <div class="ui input"><input type="text" placeholder="ejemplo: 256542215"/></div>
    <h4>Ingrese Su numero de Contraseña</h4>
    <div class="ui input"><input type="password"/></div>
    <br></br><br></br>
    
    
    <a class="ui button" href="http://localhost:4000/listado">Ingresar</a>
    </center>
  </div>
  )
}

export default Principal    