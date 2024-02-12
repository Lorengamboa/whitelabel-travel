import * as React from 'react';
import { useParams } from 'react-router-dom';
import { styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import MailOutlineIcon from '@mui/icons-material/mailOutline';
import ReceiptLongIcon from '@mui/icons-material/ReceiptLong';
import StorefrontIcon from '@mui/icons-material/Storefront';
import Box from '@mui/material/Box';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';

import { useGetCustomerQuery } from '@/services/queries/customer.query';


const CustomItem = ({ label, description }) => (
  <div style={{ display: 'flex', flexDirection: 'column' }}>
    <Typography variant="body1" component="h2" gutterBottom>
      {label}
    </Typography>
    <Typography variant="body2" gutterBottom>
      {description}
    </Typography>
  </div>
);

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

const CustomerDetails = () => {
  const { id } = useParams();
  const { data: customer = {} } = useGetCustomerQuery(id);

  return (
    <Grid container spacing={4}>
      <Grid container item xs={12} columnSpacing={{ xs: 2 }}>
        <Grid item xs={3}>
          <Item sx={{ height: '100%' }}>
            <Avatar alt="Remy Sharp" src="https://placehold.co/60x60" sx={{ width: 60, height: 60 }} />
            <Typography variant="overline" gutterBottom>
              {customer.name}
            </Typography>
          </Item>
        </Grid>

        <Grid item xs={9}>
          <Item sx={{ height: '100%', flexDirection: 'row', display: 'flex', justifyContent: 'space-evenly' }}>
            <CustomItem label="URL" description={customer.url} />
            <CustomItem label="Phone Number" description={customer.phone_number} />
            <CustomItem label="Address" description={customer.address} />
          </Item>
        </Grid>
      </Grid>

      <Grid item xs={4}>
        <Item>
          <CustomerWidget
            label="Email"
            value={customer.email}
            icon={<MailOutlineIcon
              sx={{ fontSize: 40 }}
            />}
          />
        </Item>
      </Grid>
      <Grid item xs={4}>
        <Item>
          <CustomerWidget
            label="Paid invoice (coming soon)"
            value="Must be set up"
            icon={<ReceiptLongIcon
              sx={{ fontSize: 40 }}
            />}
          />
        </Item>
      </Grid>
      <Grid item xs={4}>
        <Item>
          <CustomerWidget
            label="Clients"
            value="0"
            icon={<StorefrontIcon
              sx={{ fontSize: 40 }}
            />}
          />
        </Item>
      </Grid>
    </Grid>
  );
};

export default CustomerDetails;