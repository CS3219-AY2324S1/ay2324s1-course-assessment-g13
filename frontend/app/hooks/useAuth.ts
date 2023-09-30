import { useEffect } from "react"
import { GET } from "../libs/axios/axios";
import { useSelector } from "react-redux";
import { RootState } from '../libs/redux/store';
import { notifyError } from "../components/toast/notifications";

let resetInteveralId: NodeJS.Timeout;

export default function useAuth() {
    const { userId, userRole, username } =  useSelector((state: RootState) => state.user);
    const oneMinute = 60000;

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

    const handleRefresh = async () => {
        try {
            await GET('/auth/refresh')
        } catch (error) {
            notifyError("Error Refreshing Token");
        }
    }

    return {
        isAuthenticated, 
        userRole,
        userId,
        username
    }

}
