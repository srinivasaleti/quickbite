import React, { createContext, useContext, useState } from "react";
import type { CartItems } from "./types";
import type { ProductType } from "../product/types/product";

type CartContextType = {
  addToCart: (product: ProductType) => void;
  removeFromCart: (productId: string) => void;
  removeCompleteProduct: (productId: string) => void;
  getQuantity: (productId: string) => number;
};

const CartContext = createContext<CartContextType | undefined>(undefined);

export const CartProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [cart, setCart] = useState<CartItems>({});

  const addToCart = (product: ProductType) => {
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
    setCart((prev) => {
      const newCart = { ...prev };
      delete newCart[productId];
      return newCart;
    });
  };

  const getQuantity = (productId: string): number => {
    return cart?.[productId]?.quantity || 0;
  };
  return (
    <CartContext.Provider
      value={{
        addToCart,
        removeFromCart,
        removeCompleteProduct,
        getQuantity,
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
