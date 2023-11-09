import { createSlice } from '@reduxjs/toolkit';
import {AppState} from "../store";
import { UserResponse } from '../../../(auth)/login/page';

export interface UserState {
    username: string;
    userId: number;
    authId: number;
    photoUrl?: string;
    preferredLanguage?: string;
}

interface Action {
    payload: UserResponse["user"]
}

const initialState: UserState = {
    username: '',
    userId: 0,
    authId: 0,
    photoUrl: '',
    preferredLanguage: '',
  };

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    login: (state, action: Action) => {
      state.username = action.payload.username;
      state.userId = action.payload.ID;
      state.authId = action.payload.auth_user_id;
      state.photoUrl = action.payload.photo_url ?? state.photoUrl;
      state.preferredLanguage = action.payload.preferred_language ?? state.preferredLanguage;
    },
    logout: (state) => {
      state.username = '';
      state.userId = 0;
      state.authId = 0;
      state.photoUrl = '';
      state.preferredLanguage = '';
    },
    updateUser: (state, action: Action) => {
      state.username = action.payload.username;
      state.userId = action.payload.ID;
      state.photoUrl = action.payload.photo_url ?? state.photoUrl;
      state.preferredLanguage = action.payload.preferred_language ?? state.preferredLanguage;
    },
  },
});

export const { login, logout, updateUser } = userSlice.actions;
export default userSlice.reducer;
export const selectUsername = (state : AppState) => state.user.username;
