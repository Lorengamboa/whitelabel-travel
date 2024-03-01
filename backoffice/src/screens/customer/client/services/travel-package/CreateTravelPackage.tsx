import React from 'react';
import { useForm, Resolver, Controller } from 'react-hook-form';
import { useParams } from 'react-router-dom';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import Rating from '@mui/material/Rating';

import { TravelPackageBody } from '@/types/client';
import { useCreateTravelPackageMutation } from '@/services/queries/client.query';
import { TravelPackageFormErrorProps } from './travelPackage.types';


const useCustomResolver: Resolver<TravelPackageBody> =
  (data: TravelPackageBody) => {
    const REQUIRED_FIELD_MESSAGE = `This field is required`;
    const errors: TravelPackageFormErrorProps = {};

    if (!data.packageName) {
      errors.packageName = { message: REQUIRED_FIELD_MESSAGE };
    }
    if (!data.duration) {
      errors.duration = { message: REQUIRED_FIELD_MESSAGE };
    }

    if (!data.itinerary) {
      errors.itinerary = { message: REQUIRED_FIELD_MESSAGE };
    }
    if (!data.packageIncludes) {
      errors.packageIncludes = { message: REQUIRED_FIELD_MESSAGE };
    }
    if (!data.packageExcludes) {
      errors.packageExcludes = { message: REQUIRED_FIELD_MESSAGE };
    }
    if (!data.recommendedGear) {
      errors.recommendedGear = { message: REQUIRED_FIELD_MESSAGE };
    }
    if (!data.difficultyLevel) {
      errors.difficultyLevel = { message: REQUIRED_FIELD_MESSAGE };
    }
    if (!data.price) {
      errors.price = { message: REQUIRED_FIELD_MESSAGE };
    }

    return { values: data, errors };
  };


const CreateTravelPackageForm = () => {
  const { customerId, clientId } = useParams();

  console.log(customerId, clientId)
  const {
    handleSubmit,
    formState: { errors },
    control,
  } = useForm<TravelPackageBody>({
    resolver: useCustomResolver, defaultValues: {
      packageName: '',
      duration: 0,
      itinerary: '',
      packageIncludes: '',
      packageExcludes: '',
      recommendedGear: '',
      difficultyLevel: '',
      price: 0,
    },
  });

  const { mutate: createTravelPackage } = useCreateTravelPackageMutation(customerId, clientId);

  const onSubmit = async (formData: TravelPackageBody) => {
    const data = {
      ...formData,
      duration: Number(formData.duration),
      price: Number(formData.price)
    }
    await createTravelPackage(data)
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Typography variant="h4" gutterBottom sx={{ mb: 4 }}>
        Travel package creation form
      </Typography>
      <Controller
        name="packageName"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            label="Package Name"
            error={!!errors.packageName}
            fullWidth
            sx={{ mb: 4 }}
          />
        )}
      />
      <Controller
        name="duration"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            type='number'
            label="Duration (days)"
            error={!!errors.duration}
            fullWidth
            sx={{ mb: 4 }}
          />
        )}
      />
      <Controller
        name="itinerary"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            multiline
            label="Itinerary"
            error={!!errors.itinerary}
            fullWidth
            minRows={4}
            sx={{ mb: 4 }}
          />
        )}
      />
      <Controller
        name="packageIncludes"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            multiline
            label="Package Includes"
            error={!!errors.packageIncludes}
            fullWidth
            minRows={4}
            sx={{ mb: 4 }}
          />
        )}
      />
      <Controller
        name="packageExcludes"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            multiline
            label="Package Excludes"
            error={!!errors.packageExcludes}
            minRows={4}
            fullWidth
            sx={{ mb: 4 }}
          />
        )}
      />
      <Controller
        name="recommendedGear"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            multiline
            label="Recommended Gear"
            minRows={4}
            error={!!errors.recommendedGear}
            fullWidth
            sx={{ mb: 4 }}
          />
        )}
      />
      <Controller
        name="difficultyLevel"
        control={control}
        render={({ field }) => (
          <>
            <Typography component="legend">Difficulty</Typography>
            <Rating {...field} sx={{ mb: 4 }} />
          </>
        )}
      />
      <Controller
        name="price"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            type='number'
            label="Price"
            error={!!errors.price}
            fullWidth
            sx={{ mb: 4 }}
          />
        )}
      />

      <Button variant="contained" color="secondary" type="submit">Register</Button>
    </form>
  );
};

export default CreateTravelPackageForm;