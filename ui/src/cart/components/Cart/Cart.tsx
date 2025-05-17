import { useTheme } from "styled-components";
import { Text } from "../../../common";
import { CartContainer } from "./Styled";
import { useBreakpoint } from "../../../common/hooks/useBreakpoints";
import { useCart } from "../../CartContext";
import { EmptyCart } from "./EmptyCart";
import { OrderSumary } from "./OrderSummary";
import Modal from "../../../common/components/Modal";

export const Cart = () => {
  const { fontWeights, colors } = useTheme();
  const breakpoint = useBreakpoint();
  const showFullWithContainer = breakpoint != "l" && breakpoint !== "xl";
  const { getTotalItems } = useCart();

  const totalItems = getTotalItems();

  return (
    <CartContainer showFullWithContainer={showFullWithContainer}>
      <Text size="24px" weight={fontWeights.semiBold} color={colors.coral[500]}>
        Your Cart ({totalItems})
      </Text>
      {totalItems === 0 ? <EmptyCart /> : <OrderSumary />}
      <Modal>
        <OrderSumary />
      </Modal>
    </CartContainer>
  );
};
