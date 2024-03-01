import { useQuery, useMutation } from '@tanstack/react-query';

import { endpoints } from '@/config/endpoints';
import { get, post, ResponseData } from '@/services/network/network-service';
import { ClientBody, TravelPackageBody } from '@/types/client';
import { replaceParams } from '@/utils';


// Get all clients
export const useGetClientsQuery = (id: string) => {
  return useQuery({
    queryKey: ['clients'],
    queryFn: async () => {
      const data: ResponseData = await get(replaceParams(endpoints.backoffice.clients, { id }));
      return data;
    }
  })
}

// Create a client
export const useCreateClientMutation = (id: string) => {
  return useMutation({
    mutationFn: async (client: ClientBody) => {
      const sanitezedUrl = replaceParams(endpoints.backoffice.clients, { id });
      const data: ResponseData = await post(sanitezedUrl, client);
      return data;
    }
  })
}

// Get a client
export const useGetClientQuery = (customerId: string, clientId: string) => {
  return useQuery({
    queryKey: ['client', clientId],
    queryFn: async () => {
      const sanitezedUrl = replaceParams(endpoints.backoffice.clients, { id: customerId, clientId });
      const data: ResponseData = await get(sanitezedUrl);
      return data;
    }
  })
}

// deploy a client
export const useDeployClientMutation = () => {
  return useMutation({
    mutationFn: async ({ clientId, customerId }: { clientId: string, customerId: string }) => {
      const sanitezedUrl = replaceParams(endpoints.backoffice.clients, { id: customerId, clientId });
      const data: ResponseData = await post(`${sanitezedUrl}/deploy`);
      return data;
    }
  })
}

// Get all travel packages
export const useGetTravelPackagesQuery = (customerId: string, clientId: string) => {
  return useQuery({
    queryKey: ['travel-packages', clientId],
    queryFn: async () => {
      const sanitezedUrl = replaceParams(endpoints.backoffice.clients, { id: customerId, clientId });
      const data: ResponseData = await get(`${sanitezedUrl}/travel-packages`);
      return data;
    }
  })
}

// Create a travel package
export const useCreateTravelPackageMutation = (customerId: string, clientId: string) => {
  return useMutation({
    mutationFn: async (travelPackage: TravelPackageBody) => {
      const sanitezedUrl = replaceParams(endpoints.backoffice.clients, { id: customerId, clientId });
      const data: ResponseData = await post(`${sanitezedUrl}/travel-packages`, travelPackage);
      return data;
    }
  })
}

export const useGetTravelPackageQuery = (customerId: string, clientId: string, travelPackageId: string) => {
  return useQuery({
    queryKey: ['travel-package', travelPackageId],
    queryFn: async () => {
      const sanitezedUrl = replaceParams(endpoints.backoffice.clients, { id: customerId, clientId, travelPackageId });
      const data: ResponseData = await get(`${sanitezedUrl}/travel-packages/${travelPackageId}`);
      return data;
    }
  })
}