import React, { Component } from 'react'

export default class facturas extends Component {
    render() {
        return (
            <div className="container mt-5">
                <div class="card mb-3">

                    <div class="row">

                        <div class="col">

                            <div class="card-body">

                                <div class="d-flex bd-highlight d-flex align-items-center">

                                    <div class="flex-grow-1 bd-highlight "><h5>Factura n°: </h5></div>

                                </div>

                                <div class="table-responsive mt-2">

                                    <table class="table table-bordered table-md">
                                        <thead>
                                            <tr>
                                                <th scope="col">Producto</th>
                                                <th scope="col">Cantidad</th>
                                                <th scope="col">Precio</th>
                                                <th scope="col">Subtotal</th>
                                            </tr>

                                        </thead>

                                        <tbody>

                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                                <td></td>

                                            </tr>
                                            <tr>
                                                <td colspan="3" style={{ fontWeight: "bold" }}>TOTAL</td>
                                                <td style={{ fontWeight: "bold" }}></td>
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
