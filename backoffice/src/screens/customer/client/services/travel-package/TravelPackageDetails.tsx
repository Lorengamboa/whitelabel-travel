import * as React from 'react';
import { useParams } from 'react-router-dom';

import { Typography, Grid } from '@mui/material';
import { useGetTravelPackageQuery } from '@/services/queries/client.query';

const TravelPackageDetails = () => {
  const { customerId, clientId, travelPackageId } = useParams();

  const { data: travelPackage = {} } = useGetTravelPackageQuery(customerId, clientId, travelPackageId);

  return (
    <>
      <Typography variant="h3" gutterBottom>
        {travelPackage.packageName}
      </Typography>
      <Grid container spacing={2}>
        <Grid item xs={12} sm={6}>
          <Typography variant="body1" gutterBottom sx={{ mb: 2 }}>
            <b>Duration</b> {travelPackage.duration}
          </Typography>
          <Typography variant="body1" gutterBottom sx={{ mb: 2 }}>
            <b>Difficulty</b> {travelPackage.difficultyLevel}
          </Typography>
          <Typography variant="body1" gutterBottom sx={{ mb: 2 }}>
            <b>Iterinary:</b> {travelPackage.itinerary}
          </Typography>
          <Typography variant="body1" gutterBottom sx={{ mb: 2 }}>
            <b>Packages excluded:</b> {travelPackage.packageExcludes}
          </Typography>
          <Typography variant="body1" gutterBottom sx={{ mb: 2 }}>
            <b>Packages included:</b> {travelPackage.packageIncludes}
          </Typography>
          <Typography variant="body1" gutterBottom sx={{ mb: 2 }}>
            <b>Price:</b> {travelPackage.price}
          </Typography>
        </Grid>
      </Grid>
    </>
  );
};

export default TravelPackageDetails;