import { useSelector } from "react-redux";
import { RootState } from '../libs/redux/store';

export default function useAuth() {
    const { userId, userRole, username } =  useSelector((state: RootState) => state.user);

    const isAuthenticated = userId !== 0;

    return {
        isAuthenticated, 
        userRole,
        userId,
        username,
    }

}
