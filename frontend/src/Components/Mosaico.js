import React from 'react'
import Carta from './Carta'
function Mosaico(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.map((c, index) => (
                    <Carta nombre={c.Puntos}
                        categoria={c.Departamento}
                        descripcion={c.Indice}
                        
                        
                    />
                ))}
            </div>
        </div>
    )
}

export default Mosaico