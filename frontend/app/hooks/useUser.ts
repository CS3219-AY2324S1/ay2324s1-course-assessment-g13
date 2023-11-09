import { useSelector } from "react-redux";
import { RootState } from '../libs/redux/store';

export default function useUser() {
    const { username, photoUrl, preferredLanguage } = useSelector((state: RootState) => state.user)

    return {
        username,
        photoUrl,
        preferredLanguage
    }
}
