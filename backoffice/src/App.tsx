import { Route, createBrowserRouter, createRoutesFromElements } from 'react-router-dom';

import routes from './config/routes';
import AuthenticationLayout from './hocs/AuthenticationLayout';

// @TODO: Needs proper refactoring
const renderRoutes = () => routes.map(({ isPrivate, path, component: Component }, id) => {
  return (
    isPrivate
      ? <Route element={<AuthenticationLayout />} key={id}>
        <Route
          key={id}
          path={path}
          element={
            <Component />
          }
        />
      </Route>
      : <Route
        key={id}
        path={path}
        element={
          <Component />
        }
      />
  );
});


export const router = createBrowserRouter(
  createRoutesFromElements(
    <Route>
      {renderRoutes()}
    </Route>
  )
);