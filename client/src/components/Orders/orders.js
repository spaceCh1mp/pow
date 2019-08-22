import React, {useEffect, useState} from 'react';

//Material UI core
import Link from '@material-ui/core/Link';
import Table from '@material-ui/core/Table';
import TableRow from '@material-ui/core/TableRow';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import { makeStyles } from '@material-ui/core/styles';

//Native component
import Title from './../Title/title';

// Generate Order Data
function createData(id, date, name, type, amount, due, status) {
  return { id, date, name, type, amount, due, status};
}

const testRows = [
  createData(0, '16 Mar, 2019', 'Elvis Presley', 'member', null, null, 'approved'),
  createData(2, '16 Mar, 2019', 'Paul McCartney', 'payment', '600', '30 Aug, 2019', 'approved'),
  createData(3, '16 Mar, 2019', 'Elvis Presley', 'member', null, null, 'approved'),
  createData(7, '16 Mar, 2019', 'Paul McCartney', 'payment', '300', '30 Sept, 2019', 'rejected'),
  createData(8, '16 Mar, 2019', 'Paul McCartney', 'payment', '900', '30 Aug, 2019', 'approved'),
  createData(9, '16 Mar, 2019', 'Elvis Presley', 'member', null, null, 'rejected'),
];

const useStyles = makeStyles(theme => ({
  seeMore: {
    marginTop: theme.spacing(3),
  },
}));

export default function Orders() {
  const classes = useStyles();
  const [logs, setLogs] = useState([])
  
  useEffect(() => {
    //fetch data when component mounts
    setLogs(testRows)
  }, [])

  return (
    <React.Fragment>
      <Title h="h2" color="primary">Logs</Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell>Date</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>Type</TableCell>
            <TableCell>Amount</TableCell>
            <TableCell>Due</TableCell>
            <TableCell align="right">Status</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {logs.map(row => (
            <TableRow key={row.id}>
              <TableCell>{row.date}</TableCell>
              <TableCell>{row.name}</TableCell>
              <TableCell>{row.type}</TableCell>
              <TableCell>{row.amount}</TableCell>
              <TableCell>{row.due}</TableCell>
              <TableCell align="right">{status(row.status)}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <div className={classes.seeMore}>
        <Link color="primary" href="/sub">
          View Family
        </Link>
      </div>
    </React.Fragment>
  );
}

function status(p){
  if (p === 'approved'){
    return p
  }
  return p
}