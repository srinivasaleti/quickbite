import styled from "styled-components";

export const CART_CONTAINER_WIDTH = "400px";

export const CartContainer = styled.div`
  height: 300px;
  min-width: ${CART_CONTAINER_WIDTH};
  max-width: ${CART_CONTAINER_WIDTH};
  box-sizing: border-box;
  gap: 12px;
  padding: 24px;
  background: ${({ theme }) => theme.colors.white};
`;
