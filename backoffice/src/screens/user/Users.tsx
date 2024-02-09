import * as React from 'react';
import Fab from '@mui/material/Fab';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';

import { useGetUsersQuery } from '@/services/queries/user.query';
import UserList from './UserList';

const fabStyle = {
  position: 'absolute',
  bottom: 120,
  right: 32,
};

export default function Users() {
  const { data: users = [] } = useGetUsersQuery();

  return (
    <>
      <UserList data={users} />


      <Fab sx={fabStyle}>
        <AddCircleOutlineIcon />
      </Fab>
    </>
  );
}