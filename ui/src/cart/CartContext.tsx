import React, { createContext, useContext, useState } from "react";
import type { CartItems, OrderSummaryResponse } from "./types";
import type { ProductType } from "../product/types/product";

type CartContextType = {
  cart: CartItems;
  addToCart: (product: ProductType) => void;
  removeFromCart: (productId: string) => void;
  removeCompleteProduct: (productId: string) => void;
  getQuantity: (productId: string) => number;
  getTotalItems: () => number;
  getTotal: () => number;
  clearCart: () => void;
  order?: OrderSummaryResponse;
  cartOrderSummary?: OrderSummaryResponse;
  setOrder: (order: OrderSummaryResponse | undefined) => void;
};

const CartContext = createContext<CartContextType | undefined>(undefined);

export const CartProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [cart, setCart] = useState<CartItems>({});
  const [order, setOrder] = useState<OrderSummaryResponse | undefined>();

  const addToCart = (product: ProductType) => {
    setOrder(undefined);
    setCart((prev) => {
      if (prev[product.id]) {
        return {
          ...prev,
          [product.id]: {
            ...prev[product.id],
            quantity: prev[product.id].quantity + 1,
          },
        };
      }
      return {
        ...prev,
        [product.id]: {
          product,
          quantity: 1,
        },
      };
    });
  };

  const removeFromCart = (productId: string) => {
    setOrder(undefined);
    setCart((prev) => {
      if (!prev[productId]) return prev;

      const currentQuantity = prev[productId].quantity;

      if (currentQuantity <= 1) {
        const newCart = { ...prev };
        delete newCart[productId];
        return newCart;
      }

      return {
        ...prev,
        [productId]: {
          ...prev[productId],
          quantity: currentQuantity - 1,
        },
      };
    });
  };

  const removeCompleteProduct = (productId: string) => {
    setOrder(undefined);
    setCart((prev) => {
      const newCart = { ...prev };
      delete newCart[productId];
      return newCart;
    });
  };

  const getQuantity = (productId: string): number => {
    return cart?.[productId]?.quantity || 0;
  };

  const getTotalPriceInCents = (): number => {
    const total = Object.keys(cart).reduce(
      (prev, curr) =>
        prev + cart[curr].product.priceInCents * cart[curr].quantity,
      0
    );
    return total;
  };

  const getTotalItems = (): number => {
    return Object.keys(cart).reduce(
      (prev, curr) => prev + cart[curr].quantity,
      0
    );
  };

  const clearCart = () => setCart({});

  const cartOrderSummary: OrderSummaryResponse = {
    totalPriceInCents: getTotalPriceInCents(),
    products: Object.keys(cart).map((productID) => {
      return cart[productID].product;
    }),
    items: Object.keys(cart).map((productID) => {
      return {
        productId: productID,
        priceInCents: cart[productID].product.priceInCents,
        quantity: cart[productID].quantity,
      };
    }),
  };

  return (
    <CartContext.Provider
      value={{
        cart,
        addToCart,
        removeFromCart,
        removeCompleteProduct,
        getQuantity,
        getTotalItems,
        getTotal: getTotalPriceInCents,
        clearCart,
        setOrder,
        order,
        cartOrderSummary,
      }}
    >
      {children}
    </CartContext.Provider>
  );
};

export const useCart = () => {
  const context = useContext(CartContext);
  if (!context) throw new Error("useCart must be used within CartProvider");
  return context;
};
