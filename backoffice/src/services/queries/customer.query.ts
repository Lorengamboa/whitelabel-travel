import { useQuery, useMutation } from '@tanstack/react-query';

import { endpoints } from '@/config/endpoints';
import { get, post, ResponseData } from '@/services/network/network-service';
import { CustomerBody } from '@/types/customer';

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

// Create a customer
export const useCreateCustomerMutation = () => {
  return useMutation({
    mutationFn: async (customer: CustomerBody) => {
      const data: ResponseData = await post(endpoints.backoffice.customers, customer);
      return data;
    }
  })
}