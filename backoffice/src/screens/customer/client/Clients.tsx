import { useNavigate, useParams } from 'react-router-dom';
import { isEmpty } from 'lodash';
import Alert from '@mui/material/Alert';
import Fab from '@mui/material/Fab';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';

import { useGetClientsQuery } from '@/services/queries/client.query';
import ClientsList from './ClientsList';


const fabStyle = {
  position: 'absolute',
  bottom: 120,
  right: 32,
};

const Clients = () => {
  const { id } = useParams();
  const { data: clients = [] } = useGetClientsQuery(id);
  const navigate = useNavigate();

  return (
    <>

      {isEmpty(clients)
        ? <Alert severity="info">There are not existing clients</Alert>
        : <ClientsList data={clients} />
      }
      <Fab sx={fabStyle} onClick={() => navigate(`/customers/${id}/clients/new`)}>
        <AddCircleOutlineIcon />
      </Fab>
    </>
  );
};

export default Clients;