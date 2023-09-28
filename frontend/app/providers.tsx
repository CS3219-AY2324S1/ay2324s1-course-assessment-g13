'use client';

import { NextUIProvider } from '@nextui-org/react';
import { Provider } from 'react-redux';
import { persistor, store } from './redux/store';
import { PersistGate } from 'redux-persist/integration/react';
import { ReactNode } from 'react';

export function Providers({ children }: { children: ReactNode }) {
  return (
    <Provider store={store}>
      <PersistGate persistor={persistor} loading={null}>
        <NextUIProvider>{children}</NextUIProvider>
      </PersistGate>
    </Provider>
  );
}
