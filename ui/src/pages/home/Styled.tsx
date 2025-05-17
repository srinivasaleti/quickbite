import styled from "styled-components";
import type { Breakpoint } from "../../common/hooks/useBreakpoints";
import { FlexBox } from "../../common";

const PADDING_XL = "80px";
const PADDING_L = "50px";
const PADDING_S = "20px";

const CART_CONTAINER_WIDTH = "400px";
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
  max-width: 100%;
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
