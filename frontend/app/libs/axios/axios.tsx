import axios, { AxiosError } from 'axios';

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_URL,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
  withCredentials: true
});

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
