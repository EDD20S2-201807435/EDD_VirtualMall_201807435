import {React, useEffect,useState} from 'react'
import './Css/UserList.css'
import { Container } from 'semantic-ui-react';
import Mosaico from './Mosaico'
const axios =require('axios').default
var cors = require('cors')
function UserList() {
    const url = "http://localhost:3000/crear";
    fetch(url)
   .then(response => response.json())
   .then(json =>  {
    var cards = '';
    var ind = ''; 
	for (var i =  0; i < Object.keys(json.Vector).length; i++) {
        if (json.Vector[i].Puntos != 0) {
            if(json.Vector[i].Indice != ind){
                cards = cards+ '<div class="card" style="width: 100px;height: 100px;">'+
                '<h4 class="title">'+'<a href=http://localhost:4000/depa/'+json.Vector[i].Indice+'>'+json.Vector[i].Indice+'</a></h4>'+
                '</div>';
                ind = json.Vector[i].Indice;
            }
	
		}
	}
	document.getElementById("container").innerHTML = cards;//add cards to container
          
   });
    return (
        <div class="container" id="container">	
        </div>
    )
}

export default UserList
