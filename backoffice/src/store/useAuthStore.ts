import { create } from 'zustand';
import { persist } from 'zustand/middleware'

import { logger } from './logger';

interface AuthState {
  isAuthenticated: boolean;
}

export interface AuthStore extends AuthState {
  login: (args: AuthState['isAuthenticated']) => void;
  logout: () => void
}

const initialState: Pick<AuthStore, keyof AuthState> = {
  isAuthenticated: false,
};

const useAuthStore = create(
  logger(
    persist(
      (set) => ({
        ...initialState,
        login: (isAuthenticated: boolean) => {
          set(() => ({ isAuthenticated }));
        },
        logout: () => {
          set(() => ({ isAuthenticated: false }));
        },
      }),
      {
        name: 'auth-store',
        partialize: (state: AuthState) => ({ isAuthenticated: state.isAuthenticated }),
      },
    ),

  ));

export default useAuthStore;