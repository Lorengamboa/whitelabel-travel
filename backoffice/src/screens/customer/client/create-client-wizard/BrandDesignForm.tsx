import * as React from 'react';
import { MuiColorInput } from "mui-color-input";
import { MuiFileInput } from "mui-file-input";
import { useController } from "react-hook-form";
import Grid from '@mui/material/Grid';

export default function AddressForm() {
  const { field: fileInput } = useController({ name: "logo" });
  const { field: primaryColorInput } = useController({ name: "primaryColor" });
  const { field: secondaryColorInput } = useController({ name: "secondaryColor" });

  return (
    <>
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <MuiFileInput
            itemType="file"
            inputProps={{ accept: 'image/*' }}
            label="Logo"
            {...fileInput}
          />
        </Grid>
        <Grid item container spacing={2}>
          <Grid item>
            <MuiColorInput
              label="Primary color"
              format="hex"
              {...primaryColorInput}
            />
          </Grid>
          <Grid item>
            <MuiColorInput
              label="Secondary color"
              format="hex"
              {...secondaryColorInput}
            />
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}