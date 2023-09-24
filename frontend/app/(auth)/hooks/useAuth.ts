import { useRouter } from "next/navigation";
import { useEffect, useMemo, useState } from "react"
import { useForm } from "react-hook-form"
import { useDispatch } from "react-redux";
import { GET, POST } from "../../libs/axios/axios";
import { login, logout } from "../../libs/redux/slices/userSlice";
import { useSelector } from "react-redux";
import { RootState } from '../../libs/redux/store';
import { notifyError, notifySuccess } from "../../components/toast/notifications";

let resetInteveralId: NodeJS.Timeout;

export default function useAuth() {
    const user =  useSelector((state: RootState) => state.user);
    const userId = user.userId;
    const userRole = user.userRole;
    const dispatch = useDispatch();
    const router = useRouter();
    const oneMinute = 60000;
    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm()

    const isAuthenticated = userId !== 0;

    useEffect(()=>{
        if (isAuthenticated && !resetInteveralId) {
            resetInteveralId = setInterval(handleRefresh, oneMinute);
        }
        return () => {
            if (resetInteveralId) {
              clearInterval(resetInteveralId);
              resetInteveralId = null;
            }
        };
    }, [isAuthenticated])

    const handleLogin = handleSubmit(async data => {
        const response = await POST('/auth/login', data);
        if (response.status != 200) {
            notifyError(response.data.error);
            return;
        }
        dispatch(login(response.data.user));
        notifySuccess(response.data.message);
        router.push('/questions');
        reset();
    })

    const handleLogout = async () => {
        const response = await GET('/auth/logout')
        if (response.status != 200) {
            return;
        }
        dispatch(logout());
        router.push('/');
    }

    const handleRefresh = async () => {
        const response = await GET('/auth/refresh')
        if (response.status != 200) {
            return;
        }
    }

    const handleSignUp = handleSubmit(async (data) => {
        const response = await POST('/auth/register', data)
        if (response.status != 201) {
            notifyError(response.data.error);
            return;
        }
        notifySuccess(response.data.message);
        router.push('/login');
        reset();
    })

    return {register, errors, handleLogin, handleLogout, handleSignUp, isAuthenticated, userRole}

}