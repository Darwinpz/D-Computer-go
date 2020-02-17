import React from 'react';

import { BrowserRouter as Router, Route } from 'react-router-dom';

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js';
import 'popper.js';

import Navegacion from './components/navegation';
import Login from './components/login';
import Registrarse from './components/registrarse';
import Principal from './components/principal';
import Facturas from './components/facturas';
import Controlasistencia from './components/controlasistencia';
import Mantenimiento from './components/mantenimiento';
import Subirproducto from './components/subirproducto';
import Productos from './components/productos';
import Verproducto from './components/verproducto';
import Editarproducto from './components/editarproducto';
import Historial from './components/historial';

function App() {
  return (
    <Router>
      <Navegacion/>
      <div className="container">
        <Route path="/" exact component={Login} />
        <Route path="/registrarse" exact component={Registrarse} />
        <Route path="/principal" exact component={Principal} />
        <Route path="/facturas" exact component={Facturas} />
        <Route path="/controlasistencia" exact component={Controlasistencia} />
        <Route path="/mantenimiento" exact component={Mantenimiento} />
        <Route path="/subirproducto" exact component={Subirproducto} />
        <Route path="/productos" exact component={Productos} />
        <Route path="/:id/verproducto" component={Verproducto} />
        <Route path="/:id/editarproducto" component={Editarproducto} />
        <Route path="/historial" component={Historial}/>
      </div>
    </Router>
    
  );
}

export default App;
