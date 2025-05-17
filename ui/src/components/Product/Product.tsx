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
      <Name>{product.name}</Name>
      <Category>{product.category}</Category>
      <Price>${product.price.toFixed(2)}</Price>
    </Card>
  );
};
