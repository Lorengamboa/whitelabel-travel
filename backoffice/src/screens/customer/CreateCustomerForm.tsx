import React from 'react';
import { useForm, Resolver, Controller } from 'react-hook-form';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

import { CustomerBody } from '@/types/customer';
import { useCreateCustomerMutation } from '@/services/queries/customer.query';
import { CustomerFormErrorProps } from './customer.types';


const useCustomResolver: Resolver<CustomerBody> =
  (data: CustomerBody) => {
    const REQUIRED_FIELD_MESSAGE = `This field is required`;
    const errors: CustomerFormErrorProps = {};

    if (!data.name) {
      errors.name = { message: REQUIRED_FIELD_MESSAGE };
    }

    if (!data.phone_number) {
      errors.phone_number = { message: REQUIRED_FIELD_MESSAGE };
    }

    if (!data.email) {
      errors.email = { message: REQUIRED_FIELD_MESSAGE };
    }

    if (!data.address) {
      errors.address = { message: REQUIRED_FIELD_MESSAGE };
    }

    if (!data.url) {
      errors.url = { message: REQUIRED_FIELD_MESSAGE };
    }

    return { values: data, errors };
  };


const CreateCustomerForm = () => {
  const {
    handleSubmit,
    formState: { errors },
    control,
  } = useForm<CustomerBody>({
    resolver: useCustomResolver, defaultValues: {
      name: "",
      phone_number: "",
      email: "",
      address: "",
    },
  });

  const { mutate: createCustomer } = useCreateCustomerMutation();

  const onSubmit = async (data) => {
    console.log(data);
    await createCustomer(data)
    // Perform form submission logic here
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Typography variant="h4" gutterBottom>
        Customer creation form
      </Typography>
      <Controller control={control} name='name' render={({ field }) => (
        <TextField
          {...field}
          type="text"
          variant='outlined'
          color='secondary'
          label="Name"
          fullWidth
          required
          error={!!errors.name}
          sx={{ mb: 4 }}
        />
      )} />
      <Controller control={control} name='email' render={({ field }) => (
        <TextField
          {...field}
          type="email"
          variant='outlined'
          color='secondary'
          label="Email"
          fullWidth
          required
          error={!!errors.email}
          sx={{ mb: 4 }}
        />
      )} />
      <Controller control={control} name='phone_number' render={({ field }) => (
        <TextField
          {...field}
          type="text"
          variant='outlined'
          color='secondary'
          label="Phone number"
          required
          fullWidth
          sx={{ mb: 4 }}
        />
      )} />
      <Controller control={control} name='address' render={({ field }) => (
        <TextField
          {...field}
          type="text"
          variant='outlined'
          color='secondary'
          label="Address"
          fullWidth
          required
          sx={{ mb: 4 }}
        />
      )} />
      <Controller control={control} name='url' render={({ field }) => (
        <TextField
          {...field}
          type="text"
          variant='outlined'
          color='secondary'
          label="URL"
          fullWidth
          required
          sx={{ mb: 4 }}
        />
      )} />
      <Button variant="contained" color="secondary" type="submit">Register</Button>
    </form>
  );
};

export default CreateCustomerForm;