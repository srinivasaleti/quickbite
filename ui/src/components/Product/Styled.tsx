import styled from "styled-components";

export const Card = styled.div`
  width: 200px;
  border: 1px solid #ddd;
  border-radius: 12px;
  padding: 12px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
`;

export const Image = styled.img`
  width: 100%;
  height: 180px;
  object-fit: cover;
  border-radius: 8px;
`;

export const Name = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.grey[700]};
  font-size: ${({ theme }) => theme.fontSizes.productName};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  margin-top: ${({ theme }) => theme.margins.sm};
`;

export const Category = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.grey[500]};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  margin-top: ${({ theme }) => theme.margins.sm};
`;

export const Price = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.brown[500]};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  margin-top: ${({ theme }) => theme.margins.sm};
`;
