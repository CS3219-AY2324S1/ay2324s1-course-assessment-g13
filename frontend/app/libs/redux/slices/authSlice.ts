import { createSlice } from '@reduxjs/toolkit';
import { AuthResponse } from '../../../(auth)/login/page';

export interface AuthState {
    authId: number;
    oauthId: number;
    oauthProvider: string;
    role: "user" | "admin" | "super admin"
}

const initialState: AuthState = {
    authId: 0,
    oauthId: 0,
    oauthProvider: '',
    role: "user"
}

interface Action {
    payload: AuthResponse,
    type: string
}

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        login: (state, action: Action) => {
            state.authId = action.payload.ID
            state.oauthId = action.payload.oauth_id
            state.oauthProvider= action.payload.oauth_provider
            state.role= action.payload.role
        },
        logout: (state) => {
            state.authId = 0
            state.oauthId = 0
            state.oauthProvider = ''
            state.role= 'user'      
        },
        update: (state, action: Action) => {
            state.authId = action.payload.ID
            state.oauthId = action.payload.oauth_id
            state.oauthProvider= action.payload.oauth_provider
            state.role= action.payload.role
        }
    }
})

export const { login, logout, update } = authSlice.actions;
export default authSlice.reducer;
