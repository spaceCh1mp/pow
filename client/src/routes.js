import React from 'react'
// setup react router
import {Route, Switch, BrowserRouter as Router} from 'react-router-dom'
//import pages
import Register from './pages/0-Register'
import Login from './pages/1-Login'
import Home from './pages/2-Home'
import Profile from './pages/3-Profile'
import Dash from './pages/Dashboard/Dash'
import Sub from './pages/Dashboard/Sub/Sub'

//404

function NotFound(){
    return (
        <h2>Not Found</h2>
    )
}

function Routes(){
    return (
        <Router>
            <Switch>
                <Route exact path="/" component={Home} />
                <Route path="/login" component={Login} />
                <Route path="/register" component={Register} />
                <Route path="/profile" component={Profile} />
                <Route exact path="/dashboard" component={Dash} />
                <Route path="/dashboard/sub" component={Sub} />
                <Route component={NotFound} />
            </Switch>
        </Router>
    )
}

export default Routes