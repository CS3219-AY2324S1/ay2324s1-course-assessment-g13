import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_URL,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
});

const ErrorResponse = {
  status: 500,
  data: "Internal Server Error",
};

export const GET = async (url : string) => {
  try {
    const response = await axiosInstance.get(url);
    return response;
  } catch (error) {
    return ErrorResponse;
  }
}

export const POST = async (url : string, data: any) => {
  try {
    const response = await axiosInstance.post(url, data);
    return response;
  } catch (error) {
    return ErrorResponse;
  }
}

export const DELETE = async (url : string) => {
  try {
    const response = await axiosInstance.delete(url);
    return response;
  } catch (error) {
    return ErrorResponse;
  }
}

export const getData = async (url : string) => {
  try {
    const response = await axiosInstance.get(url);
    return {
      data: response.data == null ? [] : response.data
    };
  } catch (error) {
    return {
      error: error.response ? error.response.data.error : error.message 
    };
  }
}

export const createEntry = async (url : string, data: any) => {
  try {
    const response = await axiosInstance.post(url, data);
    return {
      message: response.data.message
    };
  } catch (error) {
    return {
      error: error.response ? error.response.data.error : error.message 
    };
  }
}

export const deleteEntry = async (url : string) => {
  try {
    const response = await axiosInstance.delete(url);
    return {
      message: response.data.message
    };
  } catch (error) {
    return {
      error: error.response ? error.response.data.error : error.message 
    };
  }
}
