import React from 'react'

import Link from '@material-ui/core/Link'
import ListItem from '@material-ui/core/ListItem'
import ListItemIcon from '@material-ui/core/ListItemIcon'
import ListItemText from '@material-ui/core/ListItemText'

import DashboardIcon from '@material-ui/icons/Dashboard'
import AccountCircleIcon from '@material-ui/icons/AccountCircle'

const Nav = (
  <div>
    <Link href="/">
      <ListItem button>
        <ListItemIcon>
          <DashboardIcon />
        </ListItemIcon>
        <ListItemText primary="Dashboard" />
      </ListItem>
    </Link>
    <Link href="/account">
      <ListItem button>
        <ListItemIcon> 
          <AccountCircleIcon />
        </ListItemIcon>
        <ListItemText primary="Account" />
      </ListItem>
    </Link>
  </div>
);

export default Nav