import {
  SignInScreen,
  SignUpScreen,
  DashboardScreen,
  CustomerDetailsScreen,
  CustomersScreen,
  UsersScreen,
  CreateCustomerFormScreen,
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
];

export default routes;