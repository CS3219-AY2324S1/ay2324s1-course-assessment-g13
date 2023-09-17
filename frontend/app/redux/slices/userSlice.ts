import { createSlice } from '@reduxjs/toolkit';

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
        console.log(action.payload);
    //   state.username = action.payload;
    },
    logout: (state) => {
      state.usernmae = null;
    },
  },
});

export const { login, logout } = userSlice.actions;
export default userSlice.reducer;
