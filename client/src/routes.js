import React from 'react'
// setup react router
import {Route, BrowserRouter as Router} from 'react-router-dom'
//import pages
import Register from './pages/0-Register'
import Login from './pages/1-Login'
import Profile from './pages/3-Profile'
import Dashboard from './pages/Dashboard/'

function Routes(){
    console.log(document.location)
    return (
        <Router>
                <Route path="/login" component={Login} />
                <Route path="/register" component={Register} />
                <Route path="/account" component={Profile} />
                <Route exact path="/dashboard" component={Dashboard} />
        </Router>
    )
}

export default Routes