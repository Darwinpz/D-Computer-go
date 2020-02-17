import React, { Component } from 'react'

import { Link } from 'react-router-dom'

export default class verproducto extends Component {

    state={

    }

    render() {
        return (
            <div class="card mb-3 mt-5">

                <div class="row">

                    <div class="col-md-6 ">


                        <div class="container text-center mt-4">

                            <img src="/imagenes/d-computer.ico" class="img-fluid" alt="" />

                        </div>
                    </div>

                    <div class="col-md-6">

                        <div class="d-flex justify-content-end m-2 ">

                            <Link class="btn btn-danger ml-2 " to="/productos" role="button">Volver</Link>

                        </div>

                        <div class="card-body">

                            <h3 class="card-title text-black">Nombre del producto</h3>
                            <p class="card-text text-break text-justify">Descripción</p>
                            <p class="card-text"><strong>Categoria: </strong>Categoria</p>
                            <p class="card-text"><strong>Disponible: </strong>2636237</p>
                            <h1 class="card-text mb-3" style={{ color: "green" }}>$ Precio</h1>
                        </div>

                        <div class="d-flex justify-content-end m-2 mb-4">

                            <Link class="btn btn-success ml-2" to="/:id/editarproducto" role="button">Editar producto</Link>

                            <button class="btn btn-danger ml-2 " id="btn-eliminar">Eliminar Producto</button>
                        </div>

                    </div>

                </div>
            </div>

        )
    }
}
