import styled from "styled-components";
import type { Breakpoint } from "../../common/hooks/useBreakpoints";
import { FlexBox } from "../../common";
import { CART_CONTAINER_WIDTH } from "../../cart/components/Cart/Styled";

const PADDING_XL = "80px";
const PADDING_L = "50px";
const PADDING_S = "20px";

const RIGHT_GAP = "20px";

interface Props {
  breakpoint: Breakpoint;
}

export const HomeContainer = styled(FlexBox)<Props>`
  padding: ${({ breakpoint }) =>
    breakpoint === "xl"
      ? PADDING_XL
      : breakpoint === "l"
      ? PADDING_L
      : PADDING_S};
`;

export const ProductListContainer = styled.div<Props>`
  width: ${({ breakpoint }) =>
    breakpoint === "xl" || breakpoint === "l"
      ? `calc(100% - ${CART_CONTAINER_WIDTH})`
      : "100%"};
`;

export const StickyCartWrapper = styled.div<Props>`
  position: fixed;
  top: ${({ breakpoint }) =>
    breakpoint === "xl"
      ? PADDING_XL
      : breakpoint === "l"
      ? PADDING_L
      : PADDING_S};
  padding-right: ${({ breakpoint }) =>
    breakpoint === "xl"
      ? PADDING_XL
      : breakpoint === "l"
      ? PADDING_L
      : PADDING_S};
  right: ${RIGHT_GAP};
  width: ${CART_CONTAINER_WIDTH};
`;

export const StickyCartIcon = styled.div<Props>`
  position: fixed;
  top: 20px;
  right: 20px;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  color: white;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  z-index: 1000;
  background: white;

  ${({ breakpoint }) =>
    (breakpoint === "xl" || breakpoint === "l") &&
    `
    display: none;
  `}
`;
