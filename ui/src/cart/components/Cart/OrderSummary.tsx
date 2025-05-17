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
import type { OrderSummaryResponse } from "../../types";
import { useModal } from "../../../common/components/Modal";
import { CouponCodeInput } from "./CouponCodeInput";

const DOLLAR = "$";

export const OrderSumary = ({
  orderSummary,
  isOrderPlaced = false,
}: {
  orderSummary?: OrderSummaryResponse;
  isOrderPlaced?: boolean;
}) => {
  const breakpoint = useBreakpoint();
  const isMobileView = breakpoint != "l" && breakpoint !== "xl";
  const theme = useTheme();
  const { placeOrder, loading } = useOrder();
  const { removeCompleteProduct, error, setOrder, setError } = useCart();
  const { closeModal } = useModal();

  const onConfirmOrder = async () => {
    await placeOrder();
  };

  const onStartNewOrder = () => {
    setOrder(undefined);
    closeModal();
  };

  return (
    <OrderSummaryContainer
      showFullWidthContainer={isMobileView || isOrderPlaced}
    >
      <FlexBox gap={"26px"}>
        {isOrderPlaced && (
          <FlexBox direction="column">
            <Text weight={theme.fontWeights.semiBold} size="40px">
              Order Confirmed
            </Text>
            <Text color={theme.colors.grey["500"]}>
              We hope you enjoy your food!
            </Text>
          </FlexBox>
        )}
        <FlexBox
          style={{ background: theme.colors.background, borderRadius: "20px" }}
          padding="16px"
        >
          <FlexBox
            style={{
              maxHeight: "40vh",
              overflow: "scroll",
            }}
            direction="column"
            gap={"16px"}
          >
            {orderSummary?.items.map((item) => {
              const productId = item.productId;
              const product = orderSummary.products.find(
                (product) => product.id === productId
              );
              if (!product) {
                return null;
              }
              const quantity = item.quantity;
              const totalPrice =
                ((product?.priceInCents || 0) * quantity) / 100;

              return (
                <>
                  <FlexBox
                    justifyContent="space-between"
                    alignItems="center"
                    direction="row"
                  >
                    <FlexBox>
                      <OrderSummaryItemName>
                        {product?.name}
                      </OrderSummaryItemName>
                      <FlexBox gap={"24px"} direction="row">
                        <OrderSummaryQuantity>{quantity}x</OrderSummaryQuantity>
                        <FlexBox gap={"12px"} direction="row">
                          <PriceInfo color={theme.colors.grey["500"]}>
                            @{DOLLAR}
                            {product?.price}
                          </PriceInfo>
                          {!isOrderPlaced && (
                            <PriceInfo color={theme.colors.grey["500"]}>
                              {DOLLAR}
                              {totalPrice}
                            </PriceInfo>
                          )}
                        </FlexBox>
                      </FlexBox>
                    </FlexBox>
                    {!isOrderPlaced ? (
                      <DeleteBtn
                        onClick={() => removeCompleteProduct(product.id)}
                      >
                        X
                      </DeleteBtn>
                    ) : (
                      <PriceInfo color={theme.colors.grey["700"]}>
                        {DOLLAR}
                        {totalPrice}
                      </PriceInfo>
                    )}
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
              ${(orderSummary?.totalPriceInCents || 0) / 100}
            </Text>
          </OrderTotal>
        </FlexBox>
        {!isOrderPlaced ? (
          <CouponCodeInput onCouponRemoved={() => setError(undefined)} />
        ) : orderSummary?.couponCode ? (
          <Text
            weight={theme.fontWeights.bold}
            color={theme.colors.coral["500"]}
          >
            Applied Coupon: {orderSummary?.couponCode}
          </Text>
        ) : null}
        {error && <Text color={theme.colors.coral["700"]}>{error}</Text>}
        <ConfirmBtn
          disabled={loading}
          onClick={isOrderPlaced ? onStartNewOrder : onConfirmOrder}
        >
          {isOrderPlaced ? "Start New Order" : "Confirm Order"}
        </ConfirmBtn>
      </FlexBox>
    </OrderSummaryContainer>
  );
};
