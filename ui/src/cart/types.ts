import type { ProductType } from "../product/types/product";

export type CartItem = {
  product: ProductType;
  quantity: number;
};

export type CartItems = {
  [id: string]: CartItem;
};
