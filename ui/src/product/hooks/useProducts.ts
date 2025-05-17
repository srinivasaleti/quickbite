import { useEffect } from "react";
import type { ProductType } from "../types/product";
import { useApi } from "../../common/hooks/useApi";
import { toast } from "react-toastify";

export function useProducts() {
  const { data, loading, error, request } = useApi<ProductType[]>();

  useEffect(() => {
    request({
      config: { method: "GET", url: "/product" },
      onError: () => {
        toast("Unable to get products", {
          type: "error",
        });
      },
    });
  }, []);

  return {
    products: data,
    loading,
    error,
    request,
  };
}
