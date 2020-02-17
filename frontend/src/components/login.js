import React, { Component } from 'react'
import axios from 'axios'
import Alertas from './alertas'

export default class login extends Component {

    state = {
        User_correo: "",
        User_clave: ""
    }

    onInputChange = (e) => {

        this.setState({

            [e.target.name]: e.target.value

        })

    }

    onSubmit = async (e) => {

        e.preventDefault();

        const response = await axios.post('http://192.168.50.5:4000/dcomputer/api/usuarios/login', {
            User_correo: this.state.User_correo,
            User_clave: this.state.User_clave
        });

        console.log(response.data)

        if (response.data.tipo === "error") {
            //const alert = useAlert();
            //alert.error("You just broke something!");
            document.getElementById('alerta').style.visibility = 'visible';

        } else {
            window.location.href = "/principal"
        }

        //


    }



    render() {
        return (
            <div className="container">
                <Alertas />

                <div class="row">

                   <div class="col-md-4 mx-auto">

                        <div class="card text-center">

                            <div class="card-header">

                                <h4>Iniciar sesión</h4>

                            </div>

                            <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" style={{width:"200px", height:"200px"}}/>

                            <h3>D-computer</h3>

                            <div class="card-body">

                                <form onSubmit={this.onSubmit}>

                                    <div class="form-group">

                                        <input type="email" class="form-control" name="User_correo" onChange={this.onInputChange} placeholder="Correo" value={this.state.User_correo} />

                                    </div>

                                    <div class="form-group">

                                        <input type="password" class="form-control" name="User_clave" onChange={this.onInputChange} placeholder="Contraseña" value={this.state.User_clave} />

                                    </div>

                                    <button type="submit" class="btn btn-primary btn-block">Ingresar</button>

                                </form>

                            </div>
                        </div>

                    </div>
                </div>
            </div>

        )
    }
}
