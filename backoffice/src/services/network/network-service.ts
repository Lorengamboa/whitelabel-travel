import axios, { AxiosInstance, AxiosResponse, AxiosError } from 'axios';

// Define the base URL for your API
const BASE_URL = '/api';

// Create an Axios instance with default configuration
const axiosInstance: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  timeout: 10000, // 10 seconds timeout
});

// Define response and error types
export interface ResponseData extends AxiosResponse {
  // Define your response data structure
}

interface ErrorResponse {
  // Define your error response structure
  message: string;
}

// Define a generic function to make GET requests
export const get = async <T>(url: string): Promise<T> => {
  try {
    const response: AxiosResponse<T> = await axiosInstance.get<T>(url);
    return response.data;
  } catch (error) {
    handleAxiosError(error);
    throw error;
  }
};

// Define a generic function to make POST requests
export const post = async <T>(url: string, body: unknown): Promise<T> => {
  try {
    const response: AxiosResponse<T> = await axiosInstance.post<T>(url, body);
    return response.data;
  } catch (error) {
    handleAxiosError(error);
    throw error;
  }
};

// Define a generic function to make DELETE requests
export const del = async <T>(url: string): Promise<T> => {
  try {
    const response: AxiosResponse<T> = await axiosInstance.delete<T>(url);
    return response.data;
  } catch (error) {
    handleAxiosError(error);
    throw error;
  }
};

// Define a function to handle Axios errors
const handleAxiosError = (error: AxiosError<ErrorResponse>): void => {
  if (error.response) {
    // The request was made and the server responded with a status code
    console.error('Response Error:', error.response.data);
    console.error('Status:', error.response.status);
  } else if (error.request) {
    // The request was made but no response was received
    console.error('Request Error:', error.request);
  } else {
    // Something happened in setting up the request that triggered an Error
    console.error('Error:', error.message);
  }
};