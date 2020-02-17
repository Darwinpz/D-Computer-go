import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import axios from 'axios'

export default class productos extends Component {

    state = {
        productos: []
    }

    onInputChange = (e) => {

        this.setState({

            [e.target.name]: e.target.value

        })

    }

    async componentDidMount() {
        const productos = await axios.get('http://192.168.50.5:4000/dcomputer/api/productos');
        this.setState({ productos: productos.data })

    }


    render() {
        return (


            <div className="card mt-5">

                <div className="card-header d-flex justify-content-between align-items-center">

                    <h3 class="card-title text-black">Productos</h3>
                    <Link class="btn btn-info" id="btn-toggle-producto" to="/subirproducto">Subir un producto</Link>

                </div>

                <div class="card-body">

                    <div class="row">

                        <div class="col-md-3 mb-2" >

                            <div class="card" >
                                <img src="/imagenes/d-computer.ico" class="img-fluid card-img-top " alt="" />

                                <div class="card-body">
                                    {
                                        this.state.productos.map(producto => (
                                            <div className="container">
                                                <p class="text-center mb-1" style={{ fontWeight: "bold" }}>{producto.Prod_nombre}</p>
                                                <h5 class="text-center mb-2" style={{ color: "green" }}>{producto.Prod_precio_venta}</h5>
                                                <Link class="btn btn-primary btn-block" to="/:id/verproducto" role="button">Ver producto</Link>
                                            </div>
                                        ))
                                    }

                                </div>
                            </div>

                        </div>

                    </div>
                </div>
            </div>

        )
    }
}
