import { useRouter } from "next/navigation";
import { useEffect } from "react"
import { useForm } from "react-hook-form"
import { GET, POST } from "../../libs/axios/axios";
import { login, logout } from "../../libs/redux/slices/userSlice";
import { useSelector, useDispatch } from "react-redux";
import { RootState } from '../../libs/redux/store';
import { notifyError, notifySuccess } from "../../components/toast/notifications";

let resetInteveralId: NodeJS.Timeout;

export default function useAuth() {
    const { userId, userRole } =  useSelector((state: RootState) => state.user);
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
        try {
            const response = await POST('/auth/login', data);
            console.log(response);
            dispatch(login(response.data.user));
            notifySuccess(response.data.message);
            router.push('/questions');
            reset();
        } catch (error) {
            notifyError(error.data.error);
        }
    })

    const handleLogout = async () => {
        try {
            dispatch(logout());
            router.push('/');
            await GET('/auth/logout')
        } catch (error) {
            notifyError(error.data.error);
        }
    }

    const handleRefresh = async () => {
        try {
            await GET('/auth/refresh')
        } catch (error) {
            notifyError("Error Refreshing Token");
        }
    }

    const handleSignUp = handleSubmit(async (data) => {
        try {
            const response = await POST('/auth/register', data)
            notifySuccess(response.data.message);
            router.push('/login');
            reset();
        } catch (error) {
            notifyError(error.data.error);
        }
    })

    const handleGithubLogin = async () => {
        const clientId = "e2d4b8fe671589d0d378" // Should move to .env
        const redirectUrl = "http://localhost:3000/oauth/callback"
        const github_authorize_url = `https://github.com/login/oauth/authorize?client_id=${clientId}&redirect_uri=${redirectUrl}`
        window.location.href = github_authorize_url;
    }

    const handleGithubLoginCallback = async (code: string) => {
        try {
            const response = await GET(`/auth/login/github?code=${code}`)
            notifySuccess(response.data.message);
            dispatch(login(response.data.user));
            router.push('/questions');
        } catch (error) {
            notifyError(error.data.error);
            router.push("/login");
        }
    } 

    return {
        register, 
        errors, 
        handleLogin, 
        handleLogout, 
        handleSignUp, 
        handleGithubLogin, 
        handleGithubLoginCallback, 
        isAuthenticated, 
        userRole
    }

}