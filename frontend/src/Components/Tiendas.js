import {React, useEffect,useState} from 'react'
import './Css/Tienda.css'
import NavBar from './NavBar'
function Tiendas(props) {
    let id1 = props.match.params.id1
    let id2 = props.match.params.id2
    let id3 = props.match.params.id3
    const url = "http://localhost:3000/listadotiendas/"+id1+"/"+id2+"/"+id3;
    console.log(url);
    fetch(url)
   .then(response => response.json())
   .then(json =>  {
    var cards = '';
   
	for (var i =  0; i < Object.keys(json.Datos).length; i++) {
        if (json.Datos[i].Nombre != "") {

                cards = cards+ '<div class="card" style="width: 300px;height: 400px;">'+
			'<a href=""><img style="width: 150px;height: 150px; margin-top:10px;" src="'+json.Datos[i].Logo+'" ></a>'+
			'<h4 class="title1" >'+json.Datos[i].Nombre+'</h4>'+
			'<p class="DescTienda">Descripcion: '+json.Datos[i].Descripcion+'</p>'+
			'<p>Contacto:  '+json.Datos[i].Contacto+'</p>'+
			'</div>';

                
            
	
		}
	}
	document.getElementById("container").innerHTML = cards;//add cards to container
          
   });
    return (
        <>
        <NavBar/>
        <div class="container" id="container">	
        </div>
        </>
    )
}

export default Tiendas