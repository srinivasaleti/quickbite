import { ProductList } from "../../product/components/ProductList";
import { Cart } from "../../cart/components/Cart";
import { useBreakpoint } from "../../common/hooks/useBreakpoints";
import {
  HomeContainer,
  ProductListContainer,
  StickyCartWrapper,
} from "./Styled";
import { CartIcon } from "../../common/components/CartIcon";
import Modal, { useModal } from "../../common/components/Modal";
import { useCart } from "../../cart/CartContext";
import { OrderSumary } from "../../cart/components/Cart/OrderSummary";

export const Home = () => {
  const breakpoint = useBreakpoint();
  const { openModal } = useModal();
  const { getTotalItems, order } = useCart();

  return (
    <>
      <HomeContainer direction="row" breakpoint={breakpoint}>
        <ProductListContainer breakpoint={breakpoint}>
          <ProductList />
        </ProductListContainer>
        {breakpoint == "l" || breakpoint == "xl" ? (
          <StickyCartWrapper breakpoint={breakpoint}>
            <Cart />
          </StickyCartWrapper>
        ) : (
          <CartIcon onClick={openModal} />
        )}
        <Modal>
          {getTotalItems() !== 0 ? (
            <Cart />
          ) : (
            <OrderSumary orderSummary={order} isOrderPlaced={true} />
          )}
        </Modal>
      </HomeContainer>
    </>
  );
};
