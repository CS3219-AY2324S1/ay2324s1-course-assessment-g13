import { useSelector } from 'react-redux';
import { AppState } from '../redux/store';

const useAuth = () => {
  // TODO: Remove this, use access token to check isLoggedin state instead
  const { id, username } = useSelector((state: AppState) => state.user);

  const isLoggedIn = username !== '';

  return { isLoggedIn, username, id };
};

export default useAuth;
