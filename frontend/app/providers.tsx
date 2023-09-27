'use client';

import { NextUIProvider } from '@nextui-org/react';
import { ReactNode } from 'react';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Provider } from 'react-redux';
import { persistor, store } from './libs/redux/store';
import { PersistGate } from 'redux-persist/integration/react';
import ProtectedRoute from './libs/_protected/ProtectedRoute';

export function Providers({ children }: { children: ReactNode }) {
  return (
    <>
      <ToastContainer />
      <NextUIProvider>
        <Provider store={store}>
          <PersistGate loading={null} persistor={persistor}>
            <ProtectedRoute>
              {children}
            </ProtectedRoute>
          </PersistGate>
        </Provider>
      </NextUIProvider>
    </>
  );
}
