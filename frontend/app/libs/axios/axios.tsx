import axios, { AxiosError } from 'axios';
import { refreshToken } from './helper';
import { store } from '../redux/store';
import { logout as AuthLogout } from '../redux/slices/authSlice';
import { logout as UserLogout } from '../redux/slices/userSlice';
import { redirect } from 'next/navigation';

type TMessage = {
  message: string;
}

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_URL,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
  withCredentials: true
});

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error: AxiosError) => {
    if (error.response && error.response.status === 401) {
      await refreshToken();
      return axiosInstance(error.config);
    }
    if (error.response && error.response.status === 404) {
      console.log(error.response.data);
      const { message } = error.response.data as TMessage;
      if (message === "User Not Found!") {
        store.dispatch(AuthLogout());
        store.dispatch(UserLogout());
        window.location.href = "/";
      }
    }
    return Promise.reject(error);
  }
);

export const GET = async (url: string) => {
  try {
    const response = await axiosInstance.get(url);
    return response;
  } catch (error) {
    throw new AxiosError({ ...error.response });
  }
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const POST = async (url: string, data: any) => {
  try {
    const response = await axiosInstance.post(url, data);
    return response;
  } catch (error) {
    throw new AxiosError({ ...error.response });
  }
};

export const PUT = async (url: string, data) => {
  try {
    const response = await axiosInstance.put(url, data);
    return response;
  } catch (error) {
    throw new AxiosError({ ...error.response });
  }
};

export const DELETE = async (url: string) => {
  try {
    const response = await axiosInstance.delete(url);
    return response;
  } catch (error) {
    throw new AxiosError({ ...error.response });
  }
};
