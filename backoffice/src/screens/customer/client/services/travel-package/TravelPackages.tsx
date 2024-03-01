import { useNavigate, useParams } from 'react-router-dom';
import Fab from '@mui/material/Fab';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';

import { useGetTravelPackagesQuery } from '@/services/queries/client.query';
import TravelPackageList from './TravelPackageList';

const fabStyle = {
  position: 'absolute',
  bottom: 120,
  right: 32,
};

const TravelPackages = () => {
  const { customerId, clientId } = useParams();
  const { data: packages = [] } = useGetTravelPackagesQuery(customerId, clientId);
  const navigate = useNavigate();

  return (
    <>
      <TravelPackageList data={packages} />

      <Fab sx={fabStyle} onClick={() => navigate('new')}>
        <AddCircleOutlineIcon />
      </Fab>
    </>
  );
};

export default TravelPackages;