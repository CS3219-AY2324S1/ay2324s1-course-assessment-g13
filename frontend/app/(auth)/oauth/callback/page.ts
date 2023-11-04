"use client";
import { useRouter, useSearchParams } from "next/navigation";
import { GET, POST } from "../../../libs/axios/axios";
import { notifyError, notifySuccess } from "../../../components/toast/notifications";
import { login } from "../../../libs/redux/slices/userSlice";
import { useDispatch } from "react-redux";
import { v4 } from "uuid";

export default function OAuthCallback() {
    const param = useSearchParams();
    const code = param.get('code');
    const dispatch = useDispatch();
    const router = useRouter();

    const handleGithubLoginCallback = async (code: string) => {
        try {
            const response = await GET(`/auth/login/github?code=${code}`);
            const rUser = response.data.user;
            rUser.username = rUser.username !== "" ? rUser.username : v4();
            // Perform upsert of user login details in user service
            await POST(`users`, {
                "user_id": rUser.id,
                "username": rUser.username,
                "photo_url": rUser.picture,
                "password": "1234"
            });
            notifySuccess(response.data.message);
            dispatch(login(rUser));
            router.push('/questions');
        } catch (error) {
            notifyError(error.message.data.error);
            router.push("/login");
        }
    }

    handleGithubLoginCallback(code); 
}
