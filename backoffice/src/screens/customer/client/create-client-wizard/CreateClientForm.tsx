import * as React from 'react';
import { type SubmitHandler, useForm, FormProvider } from 'react-hook-form';
import { useParams, useNavigate } from 'react-router-dom';
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { matchIsValidColor } from "mui-color-input";
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

import WebConfigurationForm from './WebConfigurationForm';
import BrandDesignForm from './BrandDesignForm';
import CompanyInformation from './CompanyInformation';

import { ClientBody } from '@/types/client';
import { useCreateClientMutation } from '@/services/queries/client.query';

const steps = ['Website configuration', 'Brand design', 'Company information'];

function getStepContent(step: number) {
  switch (step) {
    case 0:
      return <WebConfigurationForm />;
    case 1:
      return <BrandDesignForm />;
    case 2:
      return <CompanyInformation />;
    default:
      throw new Error('Unknown step');
  }
}

const defaultValues = {
  "url": "",
  "title": "",
  "primaryColor": "#ffffff",
  "secondaryColor": "#ffffff",
  logo: undefined,
  name: "",
  email: "",
  address1: "",
  address2: "",
  city: "",
  country: "",
};


const validationSchema = [
  //validation for step1
  yup.object({
    url: yup.string().required(),
    title: yup.string().required()
  }),
  //validation for step2
  yup.object({
    logo: yup.string().required(),
    primaryColor: yup.string().test("is-valid-color", "Invalid color", matchIsValidColor).required(),
    secondaryColor: yup.string().test("is-valid-color", "Invalid color", matchIsValidColor).required(),
  }),
  //validation for step3
  yup.object({
    name: yup.string().required(),
    email: yup.string().required(),
    address1: yup.string().required(),
    address2: yup.string().required(),
    city: yup.string().required(),
    country: yup.string().required(),
  })
];


export default function CreateClientForm() {
  const [activeStep, setActiveStep] = React.useState(0);
  const { id } = useParams();
  const createClientMutation = useCreateClientMutation(id);
  const currentValidationSchema = validationSchema[activeStep];

  const methods = useForm({
    shouldUnregister: false,
    defaultValues,
    resolver: yupResolver(currentValidationSchema),
    mode: "onChange"
  });

  const { handleSubmit, trigger } = methods;

  const handleNext = async () => {
    const isStepValid = await trigger();
    if (isStepValid) setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };

  const onSubmit: SubmitHandler<ClientBody> = async (formData) => {
    // File logo to base64
    const reader = new FileReader();
    reader.readAsDataURL(formData.logo);
    reader.onload = function () {
      const data = {
        ...formData,
        logo: reader.result,
      }
      createClientMutation.mutateAsync(data);
    };
    reader.onerror = function (error) {
      console.log('Error: ', error);
    };
  };


  return (
    <React.Fragment>
      <CssBaseline />

      <Container component="main" maxWidth="md" sx={{ mb: 4 }}>
        <Paper variant="outlined" sx={{ my: { xs: 3, md: 6 }, p: { xs: 2, md: 3 } }}>
          <Typography component="h1" variant="h4" align="center">
            Client creation
          </Typography>
          <Stepper activeStep={activeStep} sx={{ pt: 3, pb: 5 }}>
            {steps.map((label) => (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
              </Step>
            ))}
          </Stepper>
          {activeStep === steps.length ? (
            <React.Fragment>
              <Typography variant="h5" gutterBottom>
                Thank you for your order.
              </Typography>
              <Typography variant="subtitle1">
                Your order number is #2001539. We have emailed your order
                confirmation, and will send you an update when your order has
                shipped.
              </Typography>
            </React.Fragment>
          ) : (
            <Box sx={{ mt: 1 }}>
              <FormProvider {...methods}>
                <form onSubmit={handleSubmit(onSubmit)}>
                  {getStepContent(activeStep)}
                  <Box sx={{ display: 'flex', justifyContent: 'flex-end' }}>
                    {activeStep !== 0 && (
                      <Button onClick={handleBack} sx={{ mt: 3, ml: 1 }}>
                        Back
                      </Button>
                    )}

                    {activeStep === steps.length - 1 ?
                      <Button
                        variant="contained"
                        sx={{ mt: 3, ml: 1 }}
                        type="submit">
                        Submit
                      </Button>
                      :
                      <Button
                        variant="contained"
                        onClick={handleNext}
                        sx={{ mt: 3, ml: 1 }}
                        type="button">
                        Next
                      </Button>
                    }

                  </Box>
                </form>
              </FormProvider>
            </Box>
          )}
        </Paper>
      </Container>
    </React.Fragment>
  );
}