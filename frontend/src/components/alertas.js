import React, { Component } from 'react'

export default class alertas extends Component {
    render() {
        return (
            <div id="alerta" class="alert alert-danger mt-3" style={{visibility:"hidden"}}>
                <span>&times;</span>Contraseña o correo incorrectos.
            </div>
        )
    }
}
