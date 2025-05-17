import { useTheme } from "styled-components";
import { FlexBox, Text } from "../../../common";
import { useBreakpoint } from "../../../common/hooks/useBreakpoints";
import { useCart } from "../../CartContext";
import {
  ConfirmBtn,
  DeleteBtn,
  HorizontalLine,
  OrderSummaryContainer,
  OrderSummaryItemName,
  OrderSummaryQuantity,
  OrderTotal,
  PriceInfo,
} from "./Styled";
import { useOrder } from "../../hooks/useOrder";

const DOLLAR = "$";

export const OrderSumary = ({}: {}) => {
  const { cart, getTotal, removeCompleteProduct } = useCart();
  const breakpoint = useBreakpoint();
  const showFullWidthContainer = breakpoint != "l" && breakpoint !== "xl";
  const theme = useTheme();
  const { placeOrder, loading } = useOrder();

  const onConfirmOrder = () => {
    placeOrder();
  };

  return (
    <OrderSummaryContainer showFullWidthContainer={showFullWidthContainer}>
      <FlexBox direction="column" gap={"16px"}>
        {Object.keys(cart).map((productId) => {
          const product = cart[productId].product;
          const quantity = cart[productId].quantity;
          const totalPrice = (product.priceInCents * quantity) / 100;

          return (
            <>
              <FlexBox
                justifyContent="space-between"
                alignItems="center"
                direction="row"
              >
                <FlexBox>
                  <OrderSummaryItemName>{product.name}</OrderSummaryItemName>
                  <FlexBox gap={"24px"} direction="row">
                    <OrderSummaryQuantity>${quantity}x</OrderSummaryQuantity>
                    <FlexBox gap={"12px"} direction="row">
                      <PriceInfo>
                        @{DOLLAR}
                        {product.price}
                      </PriceInfo>
                      <PriceInfo>
                        {DOLLAR}
                        {totalPrice}
                      </PriceInfo>
                    </FlexBox>
                  </FlexBox>
                </FlexBox>
                <DeleteBtn onClick={() => removeCompleteProduct(product.id)}>
                  X
                </DeleteBtn>
              </FlexBox>
              <HorizontalLine />
            </>
          );
        })}
      </FlexBox>

      <OrderTotal>
        <div>Order Total</div>
        <Text
          color={theme.colors.grey["700"]}
          weight={theme.fontWeights.semiBold}
          size={"28px"}
        >
          ${getTotal()}
        </Text>
      </OrderTotal>

      <ConfirmBtn disabled={loading} onClick={onConfirmOrder}>
        Confirm Order
      </ConfirmBtn>
    </OrderSummaryContainer>
  );
};
