import { useSelector } from 'react-redux';
import { RootState } from '../redux/store';

const isBrowser = () => typeof window !== "undefined";

export default function ProtectedRoute({router, children}) {
    const user = useSelector((state: RootState) => state.user);
    const isLoggedIn = user.userId !== 0;

    let protectedRoutes = [
        "/questions"
    ]

    let pathIsProtected = protectedRoutes.indexOf(router.pathname) !== -1;

    if (isBrowser() && !isLoggedIn && !pathIsProtected) {
        router.push("/");
    }

    return children;
}