import { Navigate, Outlet, useLocation } from "react-router";
import useAuthStore from "@/store/useAuthStore";

const AuthenticationLayout: React.FC = () => {

  const location = useLocation();
  const { isAuthenticated } = useAuthStore();

  if (!isAuthenticated) {
    return (
      <Navigate to="/login" state={{ from: location }} replace />
    )
  }
  return <Outlet />;
}


export default AuthenticationLayout;