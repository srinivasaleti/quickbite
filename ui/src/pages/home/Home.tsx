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

export const Home = () => {
  const breakpoint = useBreakpoint();
  const { openModal } = useModal();

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
          <Cart />
        </Modal>
      </HomeContainer>
    </>
  );
};
