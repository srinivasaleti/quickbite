import React from "react";
import type { ProductType } from "../../types/product";
import {
  AddToCartButtonContainer,
  Card,
  Category,
  Image,
  InfoBox,
  Name,
  Price,
} from "./Styled";
import { useBreakpoint } from "../../../common/hooks/useBreakpoints";
import { AddToCartButton } from "../../../cart/components/AddToCartButton/AddToCartButton";
import { useCart } from "../../../cart/CartContext";

type ProductProps = {
  product: ProductType;
};

export const Product: React.FC<ProductProps> = ({ product }) => {
  const breakpoint = useBreakpoint();
  const { addToCart, removeFromCart, getQuantity } = useCart();

  return (
    <Card breakpoint={breakpoint}>
      <div style={{ position: "relative" }}>
        <Image
          src={product.imageUrl}
          alt={product.name}
          breakpoint={breakpoint}
        />
        <AddToCartButtonContainer>
          <AddToCartButton
            onAdd={() => addToCart(product)}
            onRemove={() => removeFromCart(product.id)}
            count={getQuantity(product.id)}
          />
        </AddToCartButtonContainer>
      </div>
      <InfoBox>
        <Category>{product.category}</Category>
        <Name>{product.name}</Name>
        <Price>${product.price.toFixed(2)}</Price>
      </InfoBox>
    </Card>
  );
};
