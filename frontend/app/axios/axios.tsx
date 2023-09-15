import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080/',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
});

export const getData = async (url : string) => {
  try {
    const response = await axiosInstance.get(url);
    return {
      status: response.status, 
      data: response.data == null ? [] : response.data
    };
  } catch (error) {
    return {
      status: error.status,
      error: error.response ? error.response.data.error : error.message 
    };
  }
}

export const createEntry = async (url : string, data: any) => {
  try {
    const response = await axiosInstance.post(url, data);
    return {
      status: response.status, 
      message: response.data.message
    };
  } catch (error) {
    return {
      status: error.status,
      error: error.response ? error.response.data.error : error.message 
    };
  }
}

export const deleteEntry = async (url : string) => {
  try {
    const response = await axiosInstance.delete(url);
    return {
      status: response.status, 
      message: response.data.message
    };
  } catch (error) {
    return {
      status: error.status,
      error: error.response ? error.response.data.error : error.message 
    };
  }
}

