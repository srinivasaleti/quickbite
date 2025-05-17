import { useApi, type ApiErrorResponse } from "../../common/hooks/useApi";
import type { OrderSummaryRequest, OrderSummaryResponse } from "../types";
import { useCart } from "../CartContext";
import { useModal } from "../../common/components/Modal";

export function useOrder() {
  const { data, loading, error, request } = useApi<OrderSummaryResponse>();
  const { cart, clearCart, setOrder, setError, coupon } = useCart();

  const orderRequestBody: OrderSummaryRequest = {
    couponCode: coupon?.length ? coupon : undefined,
    items: Object.keys(cart).map((productId) => {
      return {
        productId: productId,
        quantity: cart[productId].quantity,
      };
    }),
  };
  const { openModal } = useModal();

  const placeOrder = async () => {
    return request({
      config: { method: "POST", url: "/order", data: orderRequestBody },
      onError: (error: ApiErrorResponse | null) => {
        error && setError(error.message);
      },
      onSuccess: (response) => {
        clearCart();
        setOrder(response);
        openModal();
      },
    });
  };

  return {
    products: data,
    loading,
    error,
    placeOrder,
  };
}
