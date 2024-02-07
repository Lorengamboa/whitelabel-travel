import { useQuery } from '@tanstack/react-query';

import { endpoints } from '@/config/endpoints';
import { get, ResponseData } from '@/services/network/network-service';

// Get all users
export const useGetUsersQuery = () => {
  return useQuery({
    queryKey: ['users'],
    queryFn: async () => {
      const data: ResponseData = await get(endpoints.backoffice.users);
      return data;
    }
  })
}