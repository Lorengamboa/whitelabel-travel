import * as React from 'react';
import { type SubmitHandler, useForm, Resolver, Controller } from 'react-hook-form';
import { Navigate } from 'react-router-dom';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';

import { LoginBody } from '@/types/auth';
import useAuthStore from '@/store/useAuthStore';
import { useLoginQuery } from '@/services/queries/auth.query';

import { LoginFormErrorProps } from './authentication.types';

const defaultTheme = createTheme();


const useCustomResolver: Resolver<LoginBody> =
  (data: LoginBody) => {
    const REQUIRED_FIELD_MESSAGE = `This field is required`;
    const errors: LoginFormErrorProps = {};

    if (!data.email) {
      errors.email = { message: REQUIRED_FIELD_MESSAGE };
    }

    if (!data.password) {
      errors.password = { message: REQUIRED_FIELD_MESSAGE };
    }

    return { values: data, errors };
  };


function SignIn() {
  const { setIsAuthenticated, isAuthenticated } = useAuthStore((state) => state);
  const { mutateAsync: login } = useLoginQuery();
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginBody>({
    resolver: useCustomResolver, defaultValues: {
      email: "",
      password: "",
    },
  });


  if (isAuthenticated) {
    return (
      <Navigate to="/" replace />
    )
  }

  const onSubmit: SubmitHandler<LoginBody> = async (data) => {
    await login(data);
    setIsAuthenticated(true);
  };


  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" sx={{ mt: 1 }} onSubmit={handleSubmit(onSubmit)}>
            <Controller
              control={control}
              name='email'
              render={({ field }) => (
                <TextField
                  error={!!errors.email}
                  margin="normal"
                  fullWidth
                  id="email"
                  label="Email Address"
                  autoComplete="email"
                  autoFocus
                  {...field}
                />
              )}
            />
            <Controller
              control={control}
              name='password'
              render={({ field }) => (
                <TextField
                  error={!!errors.password}
                  margin="normal"
                  fullWidth
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
                  {...field}
                />
              )}
            />
            <FormControlLabel
              control={<Checkbox value="remember" color="primary" />}
              label="Remember me"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign In
            </Button>
            <Grid container>
              {/* <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid> */}
              <Grid item>
                <Link href="register" variant="body2">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
      </Container>
    </ThemeProvider >
  );
}

export default SignIn;