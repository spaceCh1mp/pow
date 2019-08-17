import React from 'react'

import {HashRouter, Route} from 'react-router-dom'

import Dash from './Dash/Dash'
import Sub from './Sub/Sub'

function Dashboard(){
    return(
        <HashRouter>
            <Route exact path="/" component={Dash} />
            <Route path="/sub" component={Sub} />
        </HashRouter>
    )
}


export default Dashboard