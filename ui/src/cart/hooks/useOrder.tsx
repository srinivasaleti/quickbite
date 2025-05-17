import { useApi } from "../../common/hooks/useApi";
import { toast } from "react-toastify";
import type { OrderSummaryRequest, OrderSummaryResponse } from "../types";
import { useCart } from "../CartContext";
import { useModal } from "../../common/components/Modal";

export function useOrder() {
  const { data, loading, error, request } = useApi<OrderSummaryResponse>();
  const { cart, clearCart, setOrder } = useCart();

  const orderRequestBody: OrderSummaryRequest = {
    items: Object.keys(cart).map((productId) => {
      return { productId: productId, quantity: cart[productId].quantity };
    }),
  };
  const { openModal } = useModal();

  const placeOrder = async () => {
    request({
      config: { method: "POST", url: "/order", data: orderRequestBody },
      onError: () => {
        toast("Unable to place order", {
          type: "error",
        });
      },
      onSuccess: (response) => {
        clearCart();
        setOrder(response);
        openModal()
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
