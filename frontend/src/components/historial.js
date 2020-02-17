import React, { Component } from 'react'

export default class historial extends Component {
    render() {
        return (
            <div class="row mt-5">

                <div class="col-md-4" style={{outerHeight: "550px", overflow: "auto"}}>

                    <div class="card border-secondary mb-1">
                        <div class="card-body user" >
                            <h6 class="card-title">General</h6>
                            <p class="card-text" style={{marginTop: "-10px"}}>Detalles</p>
                        </div>
                    </div>

                    
                    <div class="card border-secondary mb-1">
                        <div class="card-body user" style={{cursor: "pointer"}}>
                            <h6 class="card-title">Email</h6>
                            <p class="card-text" style={{marginTop: "-10px"}}>Nombre - Tipo</p>
                        </div>
                    </div>
                    
            
    </div>

                <div class="col-md-8">

                    <div class="card border-secondary ">
                        <div class="card-header" id="actividad">
                            ACTIVIDAD: General
                        </div>
                        <div class="card-body " id="historial" data-id="general" style={{outerHeight: "550px", overflow: "auto"}}>

                        </div>
                    </div>

                </div>

            </div>
        )
    }
}
