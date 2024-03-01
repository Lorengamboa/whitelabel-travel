import {
  SignInScreen,
  SignUpScreen,
  DashboardScreen,
  CustomerDetailsScreen,
  CustomersScreen,
  UsersScreen,
  CreateCustomerFormScreen,
  ClientsScreen,
  CreateClientFormScreen,
  ClientDetailsScreen,
  TravelPackagesScreen,
  CreateTravelPackageFormScreen,
  TravelPackageDetailsScreen,
} from '../screens';

const routes = [
  {
    path: '',
    isPrivate: true,
    component: DashboardScreen,
  },
  {
    path: 'login',
    isPrivate: false,
    component: SignInScreen,
  },
  {
    path: 'register',
    isPrivate: false,
    component: SignUpScreen,
  },
  {
    path: 'users',
    isPrivate: true,
    component: UsersScreen,
  },
  {
    path: 'customers',
    isPrivate: true,
    component: CustomersScreen,
  },
  {
    path: 'customers/:id',
    isPrivate: true,
    component: CustomerDetailsScreen,
  },
  {
    path: 'customers/new',
    isPrivate: true,
    component: CreateCustomerFormScreen,
  },
  {
    path: 'customers/:id/clients',
    isPrivate: true,
    component: ClientsScreen,
  },
  {
    path: 'customers/:customerId/clients/:clientId',
    isPrivate: true,
    component: ClientDetailsScreen,
  },
  {
    path: 'customers/:customerId/clients/:clientId/travel-packages',
    isPrivate: true,
    component: TravelPackagesScreen,
  },
  {
    path: 'customers/:customerId/clients/:clientId/travel-packages/new',
    isPrivate: true,
    component: CreateTravelPackageFormScreen,
  },
  {
    path: 'customers/:customerId/clients/:clientId/travel-packages/:travelPackageId',
    isPrivate: true,
    component: TravelPackageDetailsScreen,
  },
  {
    path: 'customers/:id/clients/new',
    isPrivate: true,
    component: CreateClientFormScreen,
  },
];

export default routes;