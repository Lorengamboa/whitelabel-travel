import { useQuery } from '@tanstack/react-query';

import { endpoints } from '@/config/endpoints';
import { get, ResponseData } from '@/services/network/network-service';

// Get all customers
export const useGetCustomersQuery = () => {
  return useQuery({
    queryKey: ['customers'],
    queryFn: async () => {
      const data: ResponseData = await get(endpoints.backoffice.customers);
      return data;
    }
  })
}

// Get a single customer
export const useGetCustomerQuery = (customerId: string) => {
  return useQuery({
    queryKey: ['customer', 'customerId'],
    queryFn: async () => {
      const data: ResponseData = await get(`${endpoints.backoffice.customers}/${customerId}`);
      return data;
    }
  })
}