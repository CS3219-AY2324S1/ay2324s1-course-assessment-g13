"use client";

import AuthCard from "../../components/card/AuthCard";
import { signOut, useSession } from "next-auth/react";
import { GET, POST } from "../../libs/axios/axios";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { notifyError, notifySuccess } from "../../components/toast/notifications";
import { useDispatch } from "react-redux";
import { AxiosResponse } from "axios";
import { login as UserLogin } from "../../libs/redux/slices/userSlice";
import { login as AuthLogin } from "../../libs/redux/slices/authSlice";

interface LoginRequest {
    oauth_id: number,
    oauth_provider: "GitHub"
}

export interface LoginResponse {
    message: string
    user?: AuthResponse
}

export interface UserResponse {
    message: string,
    user?: {
        authId: number;
        username: string,
        photo_url?: string,
        preferred_language?: string,
        auth_user_id: number,
        ID: number
    }
}

export interface AuthResponse {
    CreatedAt: string,
    UpdatedAt: string,
    DeleteAt?: string, 
    ID: number,
    role: "user" | "admin" | "super admin",
    oauth_id: number,
    oauth_provider: string
}



export default function LoginPage() {
    const {data: session, status} = useSession();
    const router = useRouter();
    const dispatch = useDispatch();

    const handleLogin = async () => {
        const id = Number(session.user.id)
        const loginRequest: LoginRequest = {
            oauth_id: id,
            oauth_provider: "GitHub"
        }
        try {
            const authResponse: AxiosResponse<LoginResponse> = await POST(`auth/login`, loginRequest);
            const { message: authMessage, user: auth } = authResponse.data
            dispatch(AuthLogin(auth));
            const authId = auth.ID
            const userResponse: AxiosResponse<UserResponse> = await GET(`/users/${authId}`)
            const { user } = userResponse.data
            dispatch(UserLogin(user));
            notifySuccess(authMessage);
            router.push('/questions');
        } catch (error) {
            const message = error.message.data.message;
            notifyError(message);
            signOut();
        }
    }

    useEffect(()=>{
        if (status === "authenticated") {
            handleLogin()
        }
    }, [status, session])

    return (
        <AuthCard authTitle={"Login"} />
    )
}
