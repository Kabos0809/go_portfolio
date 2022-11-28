import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
} from 'react-router-dom'
import { Login } from "./components/user/userLogIn"

export default class App extends Component {
  render() {
    return (
      <Router>
        <Switch>
        <Route exact path="/login" component={Login}/>
        <Admin>
          <Route exact path="/logout"/>
          <Route exact path="/admin"/>
        </Admin>
        </Switch>
      </Router>
    )
  }
}
