import React, { Component } from 'react'
import { Link } from 'react-router-dom'

export default class navigation extends Component {


    render() {
        return (
            <nav className="navbar navbar-expand-lg navbar-light" style={{backgroundColor: "#e3f2fd"}}>
 
                <div className="container" style={{color:"black", fontWeight:"bold"}}>
                    <Link className="navbar-brand" to="/">D-computer</Link>        
                    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                         <span className="navbar-toggler-icon"></span>
                    </button>
                    
                    <div className="collapse navbar-collapse" id="navbarNav">
            
                        <ul className="navbar-nav ml-auto">


                                <li className="nav-item">
                                    <Link className="dropdown-item" to="/historial">Historial</Link>
                                </li>
                                                
                                <li className="nav-item dropdown">
                                    
                                    <div className="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
                                        <Link className="dropdown-item" to="/">Salir</Link>
                                    </div>
                                </li>
                                <li className="nav-item">
                                    <Link className="dropdown-item" to="/registrarse">Registrarse</Link>
                                </li>
            </ul>
        </div>

    </div>

</nav>
        )
    }
}
