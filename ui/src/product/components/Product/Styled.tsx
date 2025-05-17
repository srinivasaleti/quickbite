import styled from "styled-components";
import { FlexBox, Text } from "../../../common";

export const Card = styled.div`
  border-radius: 12px;
  padding: ${({ theme }) => theme.margins.sm};
`;

export const Image = styled.img`
  width: 250px;
  height: 250px;
  object-fit: cover;
  border-radius: 8px;
`;

export const Name = styled(Text)`
  color: ${({ theme }) => theme.colors.grey[700]};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
`;

export const Category = styled(Text)`
  color: ${({ theme }) => theme.colors.grey[500]};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
`;

export const Price = styled(Text)`
  color: ${({ theme }) => theme.colors.brown[500]};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
`;

export const InfoBox = styled(FlexBox)`
  margin-top: 36px;
  gap: ${({ theme }) => theme.margins.sm};
`;
