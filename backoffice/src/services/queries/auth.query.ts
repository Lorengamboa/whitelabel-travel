import { UseMutationResult, useMutation } from '@tanstack/react-query';

import { type LoginBody } from '@/types/auth';
import { endpoints } from '@/config/endpoints';
import { post, ResponseData } from '@/services/network/network-service';

export const useLoginQuery = (): UseMutationResult<unknown, Error, LoginBody, unknown> => {
  return useMutation({
    mutationFn: async (loginBody: LoginBody) => {
      const { data }: ResponseData = await post(endpoints.backoffice.login, loginBody);
      return data;
    }
  })
}