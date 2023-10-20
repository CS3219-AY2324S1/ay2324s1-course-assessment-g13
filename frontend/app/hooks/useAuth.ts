import { useSelector } from "react-redux";
import { RootState } from '../libs/redux/store';
import { useSession } from "next-auth/react";

export default function useAuth() {
    const { authId, oauthId, oauthProvider, role } = useSelector((state: RootState) => state.auth)
    const { status } = useSession();
    const isLoggedIn = authId !== 0;

    return {
        isLoggedIn,
        authId,
        oauthId,
        oauthProvider,
        role,
        status
    }

}
