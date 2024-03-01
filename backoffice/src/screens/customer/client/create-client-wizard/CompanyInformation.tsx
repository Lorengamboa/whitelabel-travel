import * as React from 'react';
import {
  useFormContext,
} from "react-hook-form";
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';

export default function AddressForm() {
  const { register, formState } = useFormContext();
  const errors = formState.errors;

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Company information
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <TextField
            id="name"
            label="Company name"
            fullWidth
            autoComplete="given-name"
            variant="standard"
            error={!!errors.companyName}
            helperText={errors.companyName?.message as string}
            {...register("name")}
          />
        </Grid>
        <Grid item xs={12} sm={6} sx={{ mb: 4 }}>
          <TextField
            id="email"
            label="Company email"
            fullWidth
            autoComplete="family-name"
            variant="standard"
            error={!!errors.companyEmail}
            helperText={errors.companyEmail?.message as string}
            {...register("email")}
          />
        </Grid>

        <Grid item xs={12}>
          <Typography variant="h6" gutterBottom>
            Company address
          </Typography>
          <TextField
            id="address1"
            label="Address line 1"
            fullWidth
            autoComplete="shipping address-line1"
            variant="standard"
            error={!!errors.address1}
            helperText={errors.address1?.message as string}
            {...register("address1")}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            id="address2"
            label="Address line 2"
            fullWidth
            autoComplete="shipping address-line2"
            variant="standard"
            error={!!errors.address2}
            helperText={errors.address2?.message as string}
            {...register("address2")}
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            id="city"
            label="City"
            fullWidth
            autoComplete="shipping address-level2"
            variant="standard"
            error={!!errors.city}
            helperText={errors.city?.message as string}
            {...register("city")}
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            id="country"
            label="Country"
            fullWidth
            autoComplete="shipping country"
            variant="standard"
            error={!!errors.country}
            helperText={errors.country?.message as string}
            {...register("country")}
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}