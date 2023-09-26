import { createSlice } from '@reduxjs/toolkit';

interface UserState {
  id: number;
  username: string;
  photoUrl?: string;
}

const initialState: UserState = {
  id: 0,
  username: '',
  photoUrl: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    login: (state, action) => {
      state.id = action.payload.id;
      state.username = action.payload.username;
      state.photoUrl = action.payload.photoUrl ?? '';
    },
    logout: state => {
      state.id = 0;
      state.username = '';
      state.photoUrl = '';
    },
    updateUser: (state, action) => {
      state.username = action.payload.username;
      state.photoUrl = action.payload.photoUrl;
    },
  },
});

export const { login, logout, updateUser } = userSlice.actions;
export default userSlice.reducer;
