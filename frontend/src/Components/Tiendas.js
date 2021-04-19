import {React, useEffect,useState} from 'react'
import './Css/UserList.css'
function Tiendas(props) {
    let id1 = props.match.params.id1
    let id2 = props.match.params.id2
    let id3 = props.match.params.id3
    const url = "localhost:3000/listadotiendas/"+id1+"/"+id2+"/"+id3;
    fetch(url)
   .then(response => response.json())
   .then(json =>  {
    var cards = '';
    var ind = ''; 
	for (var i =  0; i < Object.keys(json.Datos).length; i++) {
        if (json.Datos[i].Nombre != "") {

                cards = cards+ '<div class="card" style="width: 300px;height: 370px;">'+
			'<a href=http://localhost:4000/Tiendas/'+json.Datos[i].Letra+'/'+json.Vector[i].Departamento+'/'+i+'><img src="'+json.Vector[i].Logo+'" ></a>'+
			'<h4 class="title">'+json.Vector[i].Nombre+'</h4>'+
			'<p>Descripcion: '+json.Vector[i].Descripcion+'</p>'+
			'<p>Contacto:  '+json.Vector[i].Contacto+'</p>'+
			'</div>';

                
            
	
		}
	}
	document.getElementById("container").innerHTML = cards;//add cards to container
          
   });
    return (
        <div class="container" id="container">	
        </div>
    )
}

export default Tiendas