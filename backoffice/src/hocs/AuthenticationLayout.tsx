import { Navigate, Outlet, useLocation } from "react-router";
import useAuthStore from "@/store/useAuthStore";
import { HeaderElementLayout } from '@/containers';

const AuthenticationLayout: React.FC = () => {

  const location = useLocation();
  const { isAuthenticated } = useAuthStore();

  if (!isAuthenticated) {
    return (
      <Navigate to="/login" state={{ from: location }} replace />
    )
  }
  return <HeaderElementLayout><Outlet /></HeaderElementLayout>;
}


export default AuthenticationLayout;