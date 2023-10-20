import { useSelector } from "react-redux";
import { RootState } from '../libs/redux/store';
import { useSession } from "next-auth/react";

export default function useAuth() {
    const { authId, oauthId, oauthProvider, role } = useSelector((state: RootState) => state.auth)
    const { status } = useSession();
    const previouslyLoggedIn = authId !== 0;

    return {
        previouslyLoggedIn,
        authId,
        oauthId,
        oauthProvider,
        role,
        status
    }

}
