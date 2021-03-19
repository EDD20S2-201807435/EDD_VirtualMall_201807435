import {React,useEffect,useState} from 'react'
const axios = require('axios').default
function ImportList() {
    const{productos,setproductos}=useState([])
    useEffect(()=>{
        async function obtener(){
            if(productos.length === 0){
                const data=await axios.get(url)

            }
        }
    })
    return (
        <div>
            
        </div>
    )
}

export default ImportList
