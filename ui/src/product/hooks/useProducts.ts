import { useEffect } from "react";
import type { ProductType } from "../types/product";
import { useApi } from "../../common/hooks/useApi";

export function useProducts() {
  const { data, loading, error, request } = useApi<ProductType[]>();

  useEffect(() => {
    request({ method: "GET", url: "/product" });
  }, []);

  return {
    products: data,
    loading,
    error,
    request,
  };
}
