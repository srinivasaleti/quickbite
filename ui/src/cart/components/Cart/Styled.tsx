import styled from "styled-components";

export const CART_CONTAINER_WIDTH = "400px";

export const CartContainer = styled.div<{ showFullWithContainer: boolean }>`
  height: ${({ showFullWithContainer }) =>
    showFullWithContainer ? "100%" : "300px"};
  min-width: ${({ showFullWithContainer }) =>
    showFullWithContainer ? "100%" : CART_CONTAINER_WIDTH};
  max-width: ${({ showFullWithContainer }) =>
    showFullWithContainer ? "100%" : CART_CONTAINER_WIDTH};
  box-sizing: border-box;
  gap: 12px;
  padding: 24px;
  background: ${({ theme }) => theme.colors.white};
`;
