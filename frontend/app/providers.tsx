'use client';

import { NextUIProvider } from '@nextui-org/react';
import { ReactNode } from 'react';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Provider } from 'react-redux';
import { store } from './libs/redux/store';

export function Providers({ children }: { children: ReactNode }) {
  return (
    <>
      <ToastContainer />
      <NextUIProvider>
        <Provider store={store}>
          {children}
        </Provider>
      </NextUIProvider>
    </>
  );
}
