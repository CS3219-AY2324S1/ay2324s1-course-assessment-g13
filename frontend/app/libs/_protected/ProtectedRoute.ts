import { useEffect } from 'react';
import useAuth from '../../(auth)/hooks/useAuth';
import { usePathname, useRouter } from 'next/navigation';

const isBrowser = () => typeof window !== "undefined";

export default function ProtectedRoute({children}) {
    const pathName = usePathname();
    const router = useRouter()
    const { isAuthenticated } = useAuth();

    const protectedRoutes = [
        "/questions"
    ]

    const pathIsProtected = protectedRoutes.indexOf(pathName) !== -1;

    useEffect(() => {
        if (isBrowser() && !isAuthenticated && pathIsProtected) {
            router.replace("/");
        }
        if (isBrowser() && isAuthenticated && !pathIsProtected) {
            router.replace("/questions");
        }
    }, [pathIsProtected, isAuthenticated, router])

    return children;
}