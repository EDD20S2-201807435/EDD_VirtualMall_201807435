import React from 'react'

function Carta(props) {
    return (
        <div className="column carta">
            <div className="ui card">
                
                <div className="content">
                    <div className="header">{props.nombre}</div>
                    <div className="meta">
                        <p>{props.categoria}</p>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={()=>{console.log(props.nombre)}}>Comprar</div>
                </div>
                
            </div>
        </div>
    )
}

export default Carta