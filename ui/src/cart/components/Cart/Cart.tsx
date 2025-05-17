import { useTheme } from "styled-components";
import { Text } from "../../../common";
import { EmptyCart } from "../EmptyCart";
import { CartContainer } from "./Styled";

export const Cart = () => {
  const { fontWeights, colors } = useTheme();

  return (
    <CartContainer>
      <Text size="24px" weight={fontWeights.semiBold} color={colors.coral[500]}>
        Your Cart (0)
      </Text>
      <EmptyCart />
    </CartContainer>
  );
};
