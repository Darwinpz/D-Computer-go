import React, { Component } from 'react'
import axios from 'axios'

export default class registrarse extends Component {

    state = {
        User_ced: "",
        User_nombre: "",
        User_correo: "",
        User_telefono: "",
        User_clave: "",
        Tipo_user_cod: 0
    }

    onSubmit = async (e) => {
        
        e.preventDefault();

        await axios.post('http://192.168.50.5:4000/dcomputer/api/usuarios', {
            User_ced: this.state.User_ced,
            User_nombre: this.state.User_nombre,
            User_correo: this.state.User_correo,
            User_telefono: this.state.User_telefono,
            User_clave: this.state.User_clave,
            Tipo_user_cod: this.state.Tipo_user_cod
        });

        window.location.href="/"


    }


    onInputChange = (e) => {

        this.setState({

            [e.target.name]: e.target.value

        })

    }

    render() {
        return (
            <div class="row">

                <div class="col-md-4 mx-auto">

                    <div class="card mt-4 text-center">

                        <div class="card-header">

                            <h4> Registro de Usuario</h4>

                        </div>

                        <div class="card-body">

                            <form onSubmit={this.onSubmit}>

                                <div class="form-group">

                                    <input type="text" className="form-control" onChange={this.onInputChange} name="User_ced" placeholder="Cédula" value={this.state.User_ced} />

                                </div>


                                <div class="form-group">

                                    <input type="text" className="form-control" onChange={this.onInputChange} name="User_nombre" placeholder="Nombre" value={this.state.User_nombre} />

                                </div>

                                <div class="form-group">

                                    <input type="email" class="form-control" name="User_correo" onChange={this.onInputChange} placeholder="Correo" value={this.state.User_correo}/>

                                </div>

                                
                                <div class="form-group">

                                    <input type="text" class="form-control" name="User_telefono" onChange={this.onInputChange} placeholder="Teléfono" value={this.state.User_telefono}/>

                                </div>

                                <div class="form-group">

                                    <input type="password" class="form-control" name="User_clave" onChange={this.onInputChange} placeholder="Contraseña" value={this.state.User_clave}/>
                                </div>

                                <div class="form-group">

                                    <input type="text" class="form-control" name="Tipo_user_cod" onChange={this.onInputChange} placeholder="Tipo de usuario" value={this.state.Tipo_user_cod}/>
                                </div>

                                <div class="form-group">

                                    <button type="submit" class="btn btn-primary btn-block">Registarse</button>

                                </div>
                            </form>

                        </div>

                    </div>

                </div>

            </div>
        )
    }

}
