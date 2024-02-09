import * as React from 'react';

import { useGetUsersQuery } from '@/services/queries/user.query';

export default function Dashboard() {
  const { data: users = [] } = useGetUsersQuery();

  return (
    <>
    </>
  );
}