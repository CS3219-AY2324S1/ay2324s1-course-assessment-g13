import { useRouter } from "next/navigation";
import { useEffect, useMemo, useState } from "react"
import { useForm } from "react-hook-form"
import { useDispatch } from "react-redux";
import { GET, POST } from "../../libs/axios/axios";
import { login, logout } from "../../libs/redux/slices/userSlice";
import { useSelector } from "react-redux";
import { RootState } from '../../libs/redux/store';
import { notifyError, notifySuccess } from "../../components/toast/notifications";

export default function useAuth() {
    const [refreshInterval, setRefreshInterval] = useState(null);
    const userId =  useSelector((state: RootState) => state.user.userId)
    const isLogin = userId !== 0
    const dispatch = useDispatch();
    const router = useRouter();
    const oneMinute = 60000;
    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm()

    useEffect(()=>{
        if (isLogin) {
            const initialRefreshInterval = setInterval(() => {
                handleRefresh();
            }, oneMinute);
            setRefreshInterval(initialRefreshInterval);
        }

        return () => {
            clearInterval(refreshInterval)
        }
    }, [isLogin])

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

    const handleRefresh = async() => {
        const response = await GET('/auth/refresh')
        if (response.status != 200) {
            return;
        }
    }

    return {register, errors, handleLogin, handleLogout}

}