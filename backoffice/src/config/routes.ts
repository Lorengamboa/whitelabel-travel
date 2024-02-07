import {
  SignInScreen,
  SignUpScreen,
  DashboardScreen,
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
];

export default routes;