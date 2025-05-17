import React from "react";
import type { ProductType } from "../../types/product";
import { Card, Category, Image, Name, Price } from "./Styled";

type ProductProps = {
  product: ProductType;
};

export const Product: React.FC<ProductProps> = ({ product }) => {
  return (
    <Card>
      <Image src={product.imageUrl} alt={product.name} />
      <Category>{product.category}</Category>
      <Name>{product.name}</Name>
      <Price>${product.price.toFixed(2)}</Price>
    </Card>
  );
};
