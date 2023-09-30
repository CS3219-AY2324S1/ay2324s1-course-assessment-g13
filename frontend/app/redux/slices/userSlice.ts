import { createSlice } from '@reduxjs/toolkit';
import {AppState} from "../store";

interface UserState {
    username: string;
}

const initialState: UserState = {
    username: null,
  };

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    login: (state, action) => {
        // console.log(action.payload);
      state.username = action.payload;
    },
    logout: (state) => {
      state.username = null;
    },
  },
});

export const { login, logout } = userSlice.actions;
export const selectUsername = (state : AppState) => state.user.username;
export default userSlice.reducer;