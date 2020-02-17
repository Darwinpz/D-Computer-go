import React, { Component } from 'react'
import axios from 'axios'
import { Link } from 'react-router-dom'

export default class productos extends Component {

    state = {
        Prod_cod: "",
        Prod_nombre: "",
        Prod_cantidad: '',
        Prod_precio_costo: '',
        Prod_precio_venta: '',
        Cat_cod: '',
        Prod_descripcion:"",
        Categorias: []
    }

    onSubmit = async (e) => {

        e.preventDefault();

        await axios.post('http://192.168.50.5:4000/dcomputer/api/productos', {
            Prod_cod: this.state.Prod_cod,
            Prod_nombre: this.state.Prod_nombre,
            Prod_cantidad: parseInt(this.state.Prod_cantidad),
            Prod_precio_costo: parseFloat(this.state.Prod_precio_costo),
            Prod_precio_venta:parseFloat(this.state.Prod_precio_venta),
            Cat_cod:parseInt(this.state.Cat_cod),
            Prod_descripcion:this.state.Prod_descripcion

        });
        window.location.href = "/productos"
    }

    onInputChange = (e) => {

        this.setState({

            [e.target.name]: e.target.value

        })

    }

    async componentDidMount() {
        const categorias = await axios.get('http://192.168.50.5:4000/dcomputer/api/categorias');
        this.setState({ Categorias: categorias.data })

    }

    render() {
        return (
            <div class="card mt-5">

                <div class="card-header d-flex justify-content-between align-items-center">

                    <h3 class="card-title text-black">Subir producto:</h3>
                </div>

                <div class="card-body">

                    <form onSubmit={this.onSubmit}>

                        <div class="form-group">
                            <input type="text" name="Prod_nombre" onChange={this.onInputChange} class="form-control" placeholder="Nombre del producto" value={this.state.Prod_nombre}required />

                        </div>

                        <div class="form-group">
                            <input type="number" min="0" name="Prod_cantidad" onChange={this.onInputChange} class="form-control" placeholder="Cantidad del producto" value={this.state.Prod_cantidad}required />
                        </div>

                        <div class="form-group">
                            <input type="number" min="0" step=".01" name="Prod_precio_costo" onChange={this.onInputChange} class="form-control" placeholder="Precio de compra" value={this.state.Prod_precio_costo} required />
                        </div>

                        <div class="form-group">
                        <input type="number" min="0" step=".01" name="Prod_precio_venta" onChange={this.onInputChange} class="form-control" placeholder="Precio de venta" value={this.state.Prod_precio_venta} required />
                        </div>

                        <div class="form-group">
                            <select id="form-control" onChange={this.onInputChange} name="Cat_cod"  value={this.state.Cat_cod} class="form-control" required>
                                {
                                    this.state.Categorias.map(categoria => (
                                        <option key={categoria.cat_cod} value={categoria.cat_cod}>
                                            {categoria.cat_nombre}
                                        </option>
                                    ))
                                }
                            </select>

                        </div>

                        <div class="form-group">
                            <textarea name="Prod_descripcion" onChange={this.onInputChange} class="form-control" rows="2" placeholder="Descripcion del producto" value={this.state.Prod_descripcion} required></textarea>
                        </div>

                        <div class="d-flex justify-content-end m-2 mb-4">

                            <button type="submit" class="btn btn-success ml-2">Guardar producto</button>

                            <Link class="btn btn-danger ml-2 " to="/productos" id="btn-eliminar">Cancelar</Link>
                        </div>

                    </form>


                </div>

            </div>
        )
    }
}
