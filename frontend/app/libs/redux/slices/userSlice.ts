import { createSlice } from '@reduxjs/toolkit';

export interface UserState {
    username: string;
    userId: number;
    userRole: string;
    photoUrl?: string;
}

const initialState: UserState = {
    username: '',
    userId: 0,
    userRole: '',
    photoUrl: ''
  };

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    login: (state, action) => {
      state.username = action.payload.username;
      state.userId = action.payload.id;
      state.userRole = action.payload.role;
      state.photoUrl = action.payload.picture ?? '';
    },
    logout: (state) => {
      state.username = '';
      state.userId = 0;
      state.userRole = '';
      state.photoUrl = '';
    },
    updateUser: (state, action) => {
      state.username = action.payload.username;
      state.userId = action.payload.id;
      state.userRole = action.payload.role;
      state.photoUrl = action.payload.picture ?? '';
    },
  },
});

export const { login, logout, updateUser } = userSlice.actions;
export default userSlice.reducer;
