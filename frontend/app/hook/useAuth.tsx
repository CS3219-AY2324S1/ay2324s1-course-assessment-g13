import { useSelector } from 'react-redux';
import { AppState } from '../redux/store';

const useAuth = () => {
  const username = useSelector((state: AppState) => state.user.username);

  const isLoggedIn = username !== '';

  return { isLoggedIn, username };
};

export default useAuth;
