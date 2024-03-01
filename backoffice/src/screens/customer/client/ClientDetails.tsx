import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import { Typography, Avatar, Grid, Button } from '@mui/material';
import LuggageIcon from '@mui/icons-material/Luggage';
import ColorLensIcon from '@mui/icons-material/ColorLens';
import CardMembership from '@mui/icons-material/CardMembership';
import Box from '@mui/material/Box';
import { useGetClientQuery, useDeployClientMutation } from '@/services/queries/client.query';

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  justifyContent: 'center',
  alignItems: 'center',
  color: theme.palette.text.secondary,
  display: 'flex',
  flexDirection: 'column',
}));

// Define the CustomerWidgets component
const CustomerWidget = ({ label, value, icon }) => (
  <Box sx={{ display: 'flex', flexDirection: 'row', alignItems: 'center', marginBottom: 2 }}>
    {icon}
    <Box sx={{ marginLeft: 2 }}>
      <Typography variant="subtitle1" gutterBottom>
        {label}
      </Typography>
      <Typography variant="body1" gutterBottom>
        {value}
      </Typography>
    </Box>
  </Box>
);

const ClientDetails: React.FC = () => {
  const { customerId, clientId } = useParams();
  const { data: client = {} } = useGetClientQuery(customerId, clientId);
  const { mutate: deployClient } = useDeployClientMutation();
  const navigate = useNavigate();

  const handleDeployClick = async () => {
    await deployClient({ clientId, customerId });
  };

  return (
    <>
      <Grid container justifyContent="flex-end">
        {/* Other components and markup */}
        <Button variant="contained" color="primary" onClick={handleDeployClick} size='large'>
          Deploy
        </Button>
      </Grid>
      <Typography variant="h4" gutterBottom>
        {client.title}
      </Typography>
      <Grid container spacing={2}>
        <Grid item xs={12} sm={6}>
          <Avatar
            alt={client.name}
            src={client.logo}
            sx={{ width: 100, height: 100, backgroundColor: client.primaryColor }}
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <Typography variant="h6" gutterBottom>
            {client.name}
          </Typography>
          <Typography variant="body1" gutterBottom>
            {client.url}
          </Typography>
          <Typography variant="body1" gutterBottom>
            {client.email}
          </Typography>
          <Typography variant="body1" gutterBottom>
            {client.address1}
          </Typography>
          <Typography variant="body1" gutterBottom>
            {client.address2}
          </Typography>
          <Typography variant="body1" gutterBottom>
            {client.country}
          </Typography>
          <Typography variant="body1" gutterBottom>
            {client.city}
          </Typography>
        </Grid>
      </Grid>

      <Grid marginBottom={2}>
        <Typography component="h1" variant="h5" marginBottom={2}>
          Brand identity configuration
        </Typography>

        <Grid container spacing={4}>
          <Grid item xs={4}>
            <Item>
              <CustomerWidget
                label="Colors"
                icon={<ColorLensIcon
                  sx={{ fontSize: 40 }}
                />}
              />
            </Item>
          </Grid>
        </Grid>
      </Grid>

      <Grid>
        <Typography component="h1" variant="h5" marginBottom={2}>
          Services
        </Typography>
        <Grid container spacing={4}>
          <Grid item xs={4}>
            <Item onClick={() => navigate('travel-packages')}>
              <CustomerWidget
                label="Trip offers"
                icon={<LuggageIcon
                  sx={{ fontSize: 40 }}
                />}
              />
            </Item>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
};

export default ClientDetails;