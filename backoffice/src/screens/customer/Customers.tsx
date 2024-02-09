import { useNavigate } from 'react-router-dom';
import Fab from '@mui/material/Fab';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';

import { useGetCustomersQuery } from '@/services/queries/customer.query';
import CustomerList from './CustomerList';


const fabStyle = {
  position: 'absolute',
  bottom: 120,
  right: 32,
};

const Customers = () => {
  const { data: customers = [] } = useGetCustomersQuery();
  const navigate = useNavigate();

  return (
    <>
      <CustomerList data={customers} />

      <Fab sx={fabStyle} onClick={() => navigate('/customers/new')}>
        <AddCircleOutlineIcon />
      </Fab>
    </>
  );
};

export default Customers;