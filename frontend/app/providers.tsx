'use client';

import { NextUIProvider } from '@nextui-org/react';
import { SessionProvider } from "next-auth/react";
import { ReactNode } from 'react';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Provider } from 'react-redux';
import { persistor, store } from './libs/redux/store';
import { PersistGate } from 'redux-persist/integration/react';

export function Providers({ children }: { children: ReactNode }) {
  return (
    <>
      <ToastContainer />
      <NextUIProvider>
        <SessionProvider>
          <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
              {children}
            </PersistGate>
          </Provider>
        </SessionProvider>
      </NextUIProvider>
    </>
  );
}
