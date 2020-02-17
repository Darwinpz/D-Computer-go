import React, { Component } from 'react'

import { Link } from 'react-router-dom'

export default class editarproducto extends Component {
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

                            <Link class="btn btn-danger ml-2 " to="/:id/verproducto" role="button">Cancelar</Link>

                        </div>

                        <div class="card-body">

                            <form action="/:id/verproducto/editarproducto" method="POST">

                                <div class="form-group">
                                    <p class="card-text mb-1">Nombre:</p>
                                    <input type="text" name="nombre_prod" class="form-control" value="" placeholder="Nombre del producto" required />

                                </div>

                                <div class="form-group">
                                    <p class="card-text mb-1">Descripcion:</p>
                                    <textarea name="descripcion_prod" class="form-control" rows="2" placeholder="Descripcion del producto"
                                        required>Descripcion</textarea>
                                </div>

                                <div class="form-group">
                                    <p class="card-text mb-1">Categoria:</p>
                                    <input type="text" value="" name="categoria_prod" class="form-control" placeholder="Categoria del producto" required/>
                                </div>


                                <div class="form-group">
                                    <p class="card-text mb-1">Stock: </p>
                                    <input type="number" name="stock" value="3434" class="form-control" placeholder="Cantidad del producto" required/>
                                </div>

                                <div class="form-group">
                                    <p class="card-text mb-1">Precio:</p>
                                    <input type="number" min="0" step=".01" name="precio" value="3434" class="form-control" placeholder="Precio del producto" required/>
                                </div>

                                <Link class="btn btn-success ml-2 btn-block" id="btn-actualizar"  data-id="">Actualizar</Link>
                             </form>

                        </div>

                    </div>
                </div>
            </div>
        )
    }
}
