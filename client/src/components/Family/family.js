import React, {useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Fab from '@material-ui/core/Fab'
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions'
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';


const useStyles = makeStyles(theme => ({
  card: {
    minWidth: 275,
  },
  cardAction:{
    position: 'relative',
    height: '100px'
  },
  fab: {
    margin: theme.spacing(1),
    position: 'absolute',
    bottom: theme.spacing(2),
    right: theme.spacing(2),
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
}));

const Family = (props) => {
  const classes = useStyles();

  const comp = ()=>{
    if(props.action === 'Join'){
      return <JoinGroup icon={props.icon} />
    }
    if(props.action === 'Create'){
      return <CreateGroup icon={props.icon} />
    }
    return 'No type selected'
  }
  return (
    <Card className={classes.card}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {props.title}
        </Typography>
        <Typography variant="body2" component="p">
          {props.desc}
        </Typography>
      </CardContent>
      <CardActions className={classes.cardAction}>
        {comp()}
      </CardActions>
    </Card>
  );
}

function JoinGroup(props){
  const classes = useStyles();
  const [open, setOpen] = useState(false)

  const handleDialog = () =>{
    setOpen(!open)
  }

  const joinHandler = () => {
    //pass params to joinGroup function
    //run update function if passed
    let d = document.getElementById("code").value
    joinGroup(d)
  }

  return(
    <React.Fragment>
      <Fab color="primary" aria-label="add" className={classes.fab} onClick={handleDialog}>
        {props.icon}
      </Fab>
      
      <Dialog open={open} onClose={handleDialog} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Subscribe</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Group code.
            </DialogContentText>
            <TextField
              autoFocus
              margin="dense"
              id="code"
              label="Group Code"
              type="text"
              fullWidth/>
          </DialogContent>
          <DialogActions>
          <Button onClick={handleDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={joinHandler} color="primary">
            Join
          </Button>
          </DialogActions>
      </Dialog>

    </React.Fragment>
  )
}

function joinGroup(a){
  console.log(a)
}

function CreateGroup(props){
  const classes = useStyles();
  const [open, setOpen] = useState(false)

  const handleDialog = () => {
    setOpen(!open)
  }

  const createHandler = () => {
    //call createGroup function and pass params
  }

  return(
    <React.Fragment>
      <Fab color="primary" aria-label="add" className={classes.fab} onClick={handleDialog}>
        {props.icon}
      </Fab>
      
      <Dialog open={open} onClose={handleDialog} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Subscribe</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Fill the form below to create a new family.
            </DialogContentText>
            <TextField
              autoFocus
              margin="dense"
              id="name"
              label="Group Name"
              type="text"
              fullWidth/>
          </DialogContent>
          <DialogActions>
          <Button onClick={handleDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={createHandler} color="primary">
            Create
          </Button>
          </DialogActions>
      </Dialog>

    </React.Fragment>
  )
}

export default Family





/*
  What to do
    when floating button is clicked there are two possible methods
      1: Create Family - the handler has been sent so make api call to create family then call the handler with response
      2: Join family - get family 


      */