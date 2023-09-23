import axios, { AxiosError } from 'axios';

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_USER_URL,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
});

export const GET = async (url: string) => {
  try {
    const response = await axiosInstance.get(url);
    return response;
  } catch (error) {
    return error;
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

export const getData = async (url: string) => {
  try {
    const response = await axiosInstance.get(url);
    return {
      data: response.data == null ? [] : response.data,
    };
  } catch (error) {
    return {
      error: error.response ? error.response.data.error : error.message,
    };
  }
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const createEntry = async (url: string, data: any) => {
  try {
    const response = await axiosInstance.post(url, data);
    return {
      message: response.data.message,
    };
  } catch (error) {
    return {
      error: error.response ? error.response.data.error : error.message,
    };
  }
};

export const deleteEntry = async (url: string) => {
  try {
    const response = await axiosInstance.delete(url);
    return {
      message: response.data.message,
    };
  } catch (error) {
    return {
      error: error.response ? error.response.data.error : error.message,
    };
  }
};
