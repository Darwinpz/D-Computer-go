import React, { Component } from 'react'

export default class mantenimiento extends Component {
    render() {
        return (
            <div className="container mt-5">
                <div class="card mb-3">

                    <div class="row">

                        <div class="col">

                            <div class="card-body">

                                <div class="d-flex bd-highlight d-flex align-items-center">

                                    <div class="flex-grow-1 bd-highlight "><h5>Mantenimiento: </h5></div>

                                </div>

                                <div class="table-responsive mt-2">

                                    <table class="table table-bordered table-md">
                                        <thead>
                                            <tr>
                                                <th scope="col">Fecha</th>
                                                <th scope="col">Cliente</th>
                                                <th scope="col">Atendido por</th>
                                                <th scope="col">Nombre del equipo</th>
                                                <th scope="col">Problema</th>
                                            </tr>

                                        </thead>

                                        <tbody>
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                                <td></td>

                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
