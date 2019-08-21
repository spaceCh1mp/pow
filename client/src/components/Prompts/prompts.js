import React from 'react'

import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography'
import Grid from '@material-ui/core/Grid'
import Avatar from '@material-ui/core/Avatar'
import Button from '@material-ui/core/Button'
import Chip from '@material-ui/core/Chip'
import CancelIcon from '@material-ui/icons/Cancel';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';


import {withStyles, makeStyles } from '@material-ui/core/styles';
import {green, grey} from '@material-ui/core/colors';

const ColorButton = withStyles(theme => ({
  root: {
    color: theme.palette.getContrastText(grey[900]),
    backgroundColor: green[500],
    '&:hover': {
      backgroundColor: green[700],
    },
  },
}))(Button);

//
function createData(id, type, name, amount, family){
    return {id, type, name, amount, family}
}

const rows = [
    createData(0, 1, 'Agugua Kenechukwu', null, 'spc_ops'),
    createData(1, 0, 'Ejieji Nnaemeka', '500', 'spc_ops'),
]

const useStyle = makeStyles(theme => ({
    button: {
        margin: theme.spacing(1),
    },
    rightIcon: {
        marginLeft: theme.spacing(1),
    },
    chip: {
      margin: theme.spacing(1),
    },
}))


function Prompts(props){
    const classes = useStyle()
    return (
        <React.Fragment>
        { rows.map( row => (
            <Paper key={row.id} className={props.class}>
                <Grid container wrap="nowrap" spacing={2} zeroMinWidth>
                <Grid item>
                    <Avatar>W</Avatar>
                </Grid>
                <Grid item xs noWrap>
                    {Prompt(row)}
                </Grid>
                <Grid align="right" item>
                    <ColorButton variant="contained" color="primary" className={classes.button}>
                        Approve
                        <CheckCircleIcon className={classes.rightIcon} />
                    </ColorButton>
                    
                    <Button variant="contained" color="secondary" className={classes.button}>
                        Reject
                        <CancelIcon className={classes.rightIcon} />
                    </Button>
                </Grid>
                </Grid>
            </Paper>
        ))}
        </React.Fragment>
    )
}

function Prompt(r){
    const classes = useStyle()
    var b = {}
    if (r.type === 0){
        //Do stuff for payment
        b = {
            title:'Payment',
            content: payment(r.amount)
        }
    }else{
        b = {
            title:'Member',
            content: member()
        }
    }
    //Do member stuff
    return (
        <React.Fragment>
            <Grid container wrap="nowrap" spacing={2} zeroMinWidth>
                <Grid item >
                    <Typography component="h3" variant="h6">
                        {r.name}
                    </Typography>
                </Grid>
                <Grid item>
                    <Chip
                        size="small"
                        label={r.family}
                        className={classes.chip}
                        color="primary"
                        component="a"
                        href="#chip"
                        clickable
                    />
                </Grid>
            </Grid>
            <Typography component="p" noWrap>
                {b.content}
            </Typography>
        </React.Fragment>
    )
}

function member(){
    return(
        <React.Fragment>
            wants to join your family
        </React.Fragment>
    )
}

function payment(amount){
    return(
        <React.Fragment>
            paid {amount} for a duration 2 months
        </React.Fragment>
    )
}

export default Prompts 