import {
  CircleButton,
  CountText,
  QuantityContainer,
  StyledButton,
} from "./Styled";
import cartIcon from "../../../assets/cart-icon.svg";
import { FlexBox, Text } from "../../../common";
import { useTheme } from "styled-components";
import { useState } from "react";

export const AddToCartButton = () => {
  const { fontWeights } = useTheme();
  const [added, setAdded] = useState(false);
  const [count, setCount] = useState(1);

  if (added) {
    return (
      <QuantityContainer>
        <CircleButton
          onClick={() => {
            const newCount = Math.max(1, count - 1);
            setCount(newCount);
            if (newCount === 1) setAdded(false);
          }}
        >
          −
        </CircleButton>
        <CountText>{count}</CountText>
        <CircleButton onClick={() => setCount((c) => c + 1)}>＋</CircleButton>
      </QuantityContainer>
    );
  }

  return (
    <StyledButton onClick={() => setAdded(true)}>
      <FlexBox gap={"8px"} direction="row">
        <img src={cartIcon} />
        <Text weight={fontWeights.semiBold}>Add To Cart</Text>
      </FlexBox>
    </StyledButton>
  );
};
