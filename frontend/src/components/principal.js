import React, { Component } from 'react';
import { Switch, Case } from 'react-if';
import { Link } from 'react-router-dom';
import { BrowserRouter as Router } from 'react-router-dom';
import Facturas from './facturas';
import Controlasistenca from './controlasistencia';

export default class principal extends Component {

    state = {
        tipo_cliente: "Administrador"
    }

    render() {
        return (

            <Switch>
                <Case condition={this.state.tipo_cliente === "Proveedor"}>
                    <Router>
                        <Facturas/>
                    </Router>
                
                </Case>
                <Case condition={this.state.tipo_cliente === "Cliente"}>
                    <div className="container mt-5">
                        <div className="row">
                            <div className="col">
                                <div className="card">
                                    <div class="card-header text-center">

                                        <h4>Mantenimiento</h4>
                                    </div>
                                    <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" />
                                    <div className="card-body">
                                        <Link type="submit" class="btn btn-danger btn-block" to="/mantenimiento">Ver</Link>
                                    </div>
                                </div>
                            </div>
                            <div className="col">
                                <div className="card">
                                    <div class="card-header text-center">
                                        <h4>Facturas</h4>
                                    </div>
                                    <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" />
                                    <div className="card-body">
                                        <Link type="submit" class="btn btn-danger btn-block" to="/facturas">Ver</Link>
                                    </div>
                                </div>
                            </div>
                        </div>

                    </div>
                </Case>
                <Case condition={this.state.tipo_cliente === "Trabajador"}>
                    <Router>
                        <Controlasistenca/>
                    </Router>
                </Case>
                <Case condition={this.state.tipo_cliente === "Administrador"}>

                    <div className="container mt-5">
                        <div className="row">
                            <div className="col-3">
                                <div className="card">
                                    <div class="card-header text-center">

                                        <h4>Productos</h4>
                                    </div>
                                    <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" />
                                    <div className="card-body">
                                        <Link type="submit" class="btn btn-danger btn-block" to="/productos">Ver</Link>
                                    </div>
                                </div>
                            </div>
                            <div className="col-3">
                                <div className="card">
                                    <div class="card-header text-center">
                                        <h4>Facturas</h4>
                                    </div>
                                    <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" />
                                    <div className="card-body">
                                        <Link type="submit" class="btn btn-danger btn-block" to="/facturas">Ver</Link>
                                    </div>
                                </div>
                            </div>
                            <div className="col-3">
                                <div className="card">
                                    <div class="card-header text-center">
                                        <h4>Mantenimiento</h4>
                                    </div>
                                    <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" />
                                    <div className="card-body">
                                        <Link type="submit" class="btn btn-danger btn-block" to="/mantenimiento">Ver</Link>
                                    </div>
                                </div>
                            </div>
                            <div className="col-3">
                                <div className="card">
                                    <div class="card-header text-center">
                                        <h4>Control Asistencia</h4>
                                    </div>
                                    <img src="/imagenes/d-computer.ico" class="rounded-circle mx-auto d-block m-2 logo" alt="" />
                                    <div className="card-body">
                                        <Link type="submit" class="btn btn-danger btn-block" to="/controlasistencia">Ver</Link>
                                    </div>
                                </div>
                            </div>
                        </div>

                    </div>
                </Case>
            </Switch>
        )
    }
}
