import React from 'react'

import {HashRouter, Route} from 'react-router-dom'


import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

import Header from './../../components/Header/header'

import Sub from './Sub/Sub'
import Dash from './Dash/Dash'
import Setup from './Setup/Setup'

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
  },
  fixedHeight: {
    height: 240,
  },
}));

function Dashboard(){
    const classes = useStyles()
    return(
        <div className={classes.root}>
            <CssBaseline />
            <Header title="Dashboard" />
            <HashRouter>
                <Route exact path="/" component={Dash} />
                <Route path="/sub" component={Sub} />
                <Route path="/setup" component={Setup} />
            </HashRouter>
        </div>
    )
}


export default Dashboard