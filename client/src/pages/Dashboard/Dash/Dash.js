import React, { useState, useEffect } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions'

import Orders from './../../../components/Orders/orders';
import Prompts from './../../../components/Prompts/prompts';
import Copyright from './../../../components/Copyright/copyright'


import AddIcon from '@material-ui/icons/Add'
import GroupAddIcon from '@material-ui/icons/GroupAdd'
import Family from './../../../components/Family/family'

const drawerWidth = 240;

const useStyles = makeStyles(theme => ({
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  cardAction:{
    position: 'relative',
    height: '100px'
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  promptPaper:{
    margin: `${theme.spacing(1)}px auto`,
    marginBottom: theme.spacing(2),
    padding: theme.spacing(2),
    maxWidth: 750,
    overflow: 'auto',
  }
}));

const testRows = {
  'id': 0,
  'name': 'kene',
  'family': 'spc_ops',
  'level': 2
}  
//createData(0, 'kene', null, 0)
  //createData(0, 'kene', 'spc_ops', 1)
  //createData(0, 'kene', 'spc_ops', 2)


export default function Dash() {
  const classes = useStyles();
  const [user, setUser] = useState({})
  const [update, setUpdate] = useState(false)
  const updateHandler = () => {
    //setUser(user.family = 'spc_ops')
    setUpdate(!update)
  }

  useEffect(() =>{
    //set user data
    setUser(testRows)
  },[update])

  return (
      <main className={classes.content}>
        <div className={classes.appBarSpacer} />
        <Container maxWidth="lg" className={classes.container}>
          <Grid container item xs={12} spacing={2}>
            {(user.family === null)
            ? <Init 
                cls={classes}
                handler={updateHandler}
              />
            : dash(user, classes)
            }
          </Grid>
        </Container>
        <Copyright />
      </main>
  );
}

function dash(u, cls) {
    if (u.level === 1){
      return admin(cls)
    }
    return member(cls)
}

function admin(classes) {
  //build admin page
  return(
    <React.Fragment>
      {/*Awaiting confirmation */}
      <Grid item xs={12}>
        <Prompts class={classes.promptPaper} />
      </Grid>
      <Grid item xs={12}>
        {member(classes)}
      </Grid>
      {/* Recent Orders */}
      <Grid item xs={12}>
        <Paper className={classes.paper}>
          <Orders />
        </Paper>
      </Grid>
    </React.Fragment>
  )
}

function member(cls){
  //build member page
  return (
    <React.Fragment>
      <Grid item xs={12}>
        <Card className={cls.card}>
          <CardContent>

            <Paper className={cls.paper}>
                Group Rate
            </Paper>


            <Paper className={cls.paper}>
                Due
            </Paper>
          
          </CardContent>
          <CardActions className={cls.cardAction}>
            <GroupAddIcon />
          </CardActions>
        </Card>
      </Grid>

      <Grid item xs={12}>
        <Paper className={cls.paper}>
          <Orders />
        </Paper>
      </Grid>
    </React.Fragment>
  ) 
}

function Init(props){
  return(
    <React.Fragment>
      <Grid item xs={12} sm={6}>
        <Family
          action='Create'
          title="Create"
          desc="Get in on the fun"
          icon={<AddIcon />} 
          handler={props.handler} 
        />
      </Grid>
      <Grid item xs={12} sm={6}>
        <Family 
          action="Join"
          title="Join"
          desc="Get in" 
          icon={<GroupAddIcon />}
          handler={props.handler}
        />
      </Grid>
    </React.Fragment>
  )
  //build initial board
}