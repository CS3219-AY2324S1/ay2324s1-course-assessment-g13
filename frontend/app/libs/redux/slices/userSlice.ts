import { createSlice } from '@reduxjs/toolkit';

export interface UserState {
    username: string;
    userId: number;
    userRole: string;
}

const initialState: UserState = {
    username: null,
    userId: 0,
    userRole: null
  };

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    login: (state, action) => {
      state.username = action.payload.username
      state.userId = action.payload.ID
      state.userRole = action.payload.role
    },
    logout: (state) => {
      state.username = null;
      state.userId = 0;
      state.userRole = null;
    },
  },
});

export const { login, logout } = userSlice.actions;
export default userSlice.reducer;
