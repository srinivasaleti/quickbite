import { useTheme } from "styled-components";
import { Text } from "../../../common";
import { EmptyCart } from "../EmptyCart";
import { CartContainer } from "./Styled";
import { useBreakpoint } from "../../../common/hooks/useBreakpoints";

export const Cart = () => {
  const { fontWeights, colors } = useTheme();
  const breakpoint = useBreakpoint();
  const showFullWithContainer = breakpoint != "l" && breakpoint !== "xl";

  return (
    <CartContainer showFullWithContainer={showFullWithContainer}>
      <Text size="24px" weight={fontWeights.semiBold} color={colors.coral[500]}>
        Your Cart (0)
      </Text>
      <EmptyCart />
    </CartContainer>
  );
};
