import {React, useEffect,useState} from 'react'
import './Css/UserList.css'
function Calificaciones(props) {
    let id1 = props.match.params.id1
    let id2 = props.match.params.id2
    const url = "http://localhost:3000/crear";
    fetch(url)
   .then(response => response.json())
   .then(json =>  {
    var cards = '';
    var ind = ''; 
	for (var i =  1; i < 6; i++) {
        
            
              
                cards = cards+ '<div class="card" style="width: 200px;height: 100px;">'+
                '<h4 class="title">'+'<a href=http://localhost:4000/Tiendas/'+id1+'/'+id2+'/'+i+'>'+i+'</a></h4>'+
                '</div>';
                ind = json.Vector[i].Departamento;
               
            
	
		
	}
	document.getElementById("container").innerHTML = cards;//add cards to container
          
   });
    return (
        <div class="container" id="container">	
        </div>
    )
}

export default Calificaciones