import { useState } from "react";
import axios, { type AxiosRequestConfig } from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080/api",
  headers: {
    "Content-Type": "application/json",
  },
});

interface ApiErrorResponse {
  code: string;
  message: string;
}
export function useApi<T = unknown>() {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<ApiErrorResponse | null>(null);

  const request = async ({
    config,
    onError,
    onSuccess,
  }: {
    config: AxiosRequestConfig;
    onSuccess?: (data: T) => void;
    onError?: (error: ApiErrorResponse | null) => void;
  }) => {
    setLoading(true);
    setError(null);
    try {
      const response = await api.request<T>(config);
      setData(response.data);
      if (onSuccess) {
        onSuccess(response.data);
      }
    } catch (err: any) {
      const apiError = err?.response?.data as ApiErrorResponse;
      setError(apiError);
      if (onError) {
        onError(apiError);
      }
    } finally {
      setLoading(false);
    }
  };

  return { data, loading, error, request };
}
