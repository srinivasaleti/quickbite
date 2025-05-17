import { ProductList } from "../../product/components/ProductList";
import { Cart } from "../../cart/components/Cart";
import { useBreakpoint } from "../../common/hooks/useBreakpoints";
import {
  HomeContainer,
  ProductListContainer,
  StickyCartWrapper,
} from "./Styled";

export const Home = () => {
  const breakpoint = useBreakpoint();

  return (
    <>
      <HomeContainer direction="row" breakpoint={breakpoint}>
        <ProductListContainer breakpoint={breakpoint}>
          <ProductList />
        </ProductListContainer>
        <StickyCartWrapper breakpoint={breakpoint}>
          <Cart />
        </StickyCartWrapper>
      </HomeContainer>
    </>
  );
};
