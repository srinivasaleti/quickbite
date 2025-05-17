import styled from "styled-components";

export const EmptyCartContainer = styled.div`
  height: 300px;
  min-width: 400px;
  max-width: 400px;
  box-sizing: border-box;
  gap: 12px;
  padding: 24px;
  background: ${({ theme }) => theme.colors.white};
`;

export const ImageContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 50px;
`;

export const Image = styled.img`
  width: 120px;
  height: 120px;
  margin: auto;
  border-radius: 8px;
`;
