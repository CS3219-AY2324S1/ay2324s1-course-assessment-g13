import { combineReducers, configureStore } from '@reduxjs/toolkit';
import {
  FLUSH,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
  REHYDRATE,
  persistReducer,
  persistStore,
} from 'redux-persist';
import preferenceReducer from './slices/matchPreferenceSlice';
import questionBankReducer, { QuestionBankState } from './slices/questionBankSlice';
import createWebStorage from 'redux-persist/lib/storage/createWebStorage';
import userReducer, { UserState } from './slices/userSlice';
import collabReducer from './slices/collabSlice';
import authReducer, { AuthState } from './slices/authSlice';

export interface RootState {
  user: UserState,
  questionBank: QuestionBankState,
  auth: AuthState,
}

const createNoopStorage = () => {
  return {
    getItem() {
      return Promise.resolve(null);
    },
    setItem(_key, value) {
      return Promise.resolve(value);
    },
    removeItem() {
      return Promise.resolve();
    },
  };
};

const storage = typeof window !== 'undefined' ? createWebStorage('local') : createNoopStorage();

const persistConfig = {
  key: 'root',
  storage,
  whitelist: ['questionBank', 'user', 'preference', 'collab', 'auth'],
};

const rootReducer = combineReducers({
  questionBank: questionBankReducer,
  user: userReducer,
  auth: authReducer,
  preference: preferenceReducer,
  collab: collabReducer,
});

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

export const persistor = persistStore(store);

export type AppState = ReturnType<typeof rootReducer>;
