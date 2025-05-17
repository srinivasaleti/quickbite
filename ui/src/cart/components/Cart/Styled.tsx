import styled from "styled-components";

export const CartContainer = styled.div`
  height: 300px;
  min-width: 400px;
  max-width: 400px;
  box-sizing: border-box;
  gap: 12px;
  padding: 24px;
  background: ${({ theme }) => theme.colors.white};
`;
