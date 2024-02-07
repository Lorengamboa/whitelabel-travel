import { create } from 'zustand';
import { persist } from 'zustand/middleware'

import { logger } from './logger';

interface AuthState {
  isAuthenticated: boolean;
}

export interface AuthStore extends AuthState {
  setIsAuthenticated: (args: AuthState['isAuthenticated']) => void;
}

const initialState: Pick<AuthStore, keyof AuthState> = {
  isAuthenticated: false,
};

const useAuthStore = create(
  logger(
    persist(
      (set) => ({
        ...initialState,
        setIsAuthenticated: (isAuthenticated: boolean) => {
          set(() => ({ isAuthenticated }));
        },
      }),
      {
        name: 'auth-store',
        partialize: (state: AuthState) => ({ isAuthenticated: state.isAuthenticated }),
      },
    )
  ));

export default useAuthStore;