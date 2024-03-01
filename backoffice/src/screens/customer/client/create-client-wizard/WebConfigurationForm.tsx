import * as React from 'react';
import {
  useFormContext,
} from "react-hook-form";
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';

export default function AddressForm() {
  const { register, formState } = useFormContext();
  const errors = formState.errors;

  return (
    <>
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <TextField
            required
            id="website-url"
            label="Website url"
            fullWidth
            autoComplete="given-name"
            variant="standard"
            error={!!errors.url}
            helperText={errors.url?.message as string}
            {...register("url")}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            required
            id="title"
            label="Title"
            fullWidth
            autoComplete="family-name"
            variant="standard"
            error={!!errors.title}
            helperText={errors.title?.message as string}
            {...register("title")}
          />
        </Grid>
      </Grid>
    </>
  );
}