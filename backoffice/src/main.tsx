import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router-dom';
import { QueryClientProvider, QueryClient } from '@tanstack/react-query';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { blue } from '@mui/material/colors';

import { router as App } from './App.tsx'

const queryClient = new QueryClient()

const theme = createTheme({
  palette: {
    primary: {
      main: blue[500],
    },
  },
  typography: {
    allVariants: {
      color: "black",
    },
  },
});


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <ThemeProvider theme={theme}>
        <RouterProvider router={App} />
      </ThemeProvider>
    </QueryClientProvider>
  </React.StrictMode>,
)
