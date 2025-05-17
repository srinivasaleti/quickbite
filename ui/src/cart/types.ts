import type { ProductType } from "../product/types/product";

export type CartItem = {
  product: ProductType;
  quantity: number;
};

export type CartItems = {
  [id: string]: CartItem;
};

export type OrderSummaryRequest = {
  couponCode?: string;
  items: {
    productId: string;
    quantity: number;
  }[];
};

export type OrderSummaryResponse = {
  id: string;
  totalPriceInCents: number;
  couponCode?: string;
  items: {
    productId: string;
    priceInCents: number;
    quantity: number;
  }[];
  products: ProductType[];
};
