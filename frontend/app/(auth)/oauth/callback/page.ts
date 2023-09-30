"use client";
import { useRouter, useSearchParams } from "next/navigation";
import { GET } from "../../../libs/axios/axios";
import { notifyError, notifySuccess } from "../../../components/Notifications";
import { login } from "../../../libs/redux/slices/userSlice";
import { useDispatch } from "react-redux";

export default function OAuthCallback() {
    const param = useSearchParams();
    const code = param.get('code');
    const dispatch = useDispatch();
    const router = useRouter();

    const handleGithubLoginCallback = async (code: string) => {
        try {
            const response = await GET(`/auth/login/github?code=${code}`);
            notifySuccess(response.data.message);
            dispatch(login(response.data.user));
            router.push('/questions');
        } catch (error) {
            notifyError(error.message.data.error);
            router.push("/login");
        }
    }

    handleGithubLoginCallback(code); 
}
