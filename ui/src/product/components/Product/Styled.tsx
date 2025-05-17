import styled from "styled-components";
import { FlexBox, Text } from "../../../common";
import type { Breakpoint } from "../../../common/hooks/useBreakpoints";

export const Card = styled.div<{ breakpoint: Breakpoint }>`
  position: relative;
  border-radius: 12px;
  padding: ${({ theme }) => theme.margins.sm};
  width: ${({ breakpoint }) => (breakpoint === "xs" ? "100%" : "270px")};
  height: ${({ breakpoint }) => (breakpoint === "xs" ? "auto" : "350px")};
`;

export const Image = styled.img<{ breakpoint: Breakpoint }>`
  width: ${({ breakpoint }) => (breakpoint === "xs" ? "100%" : "250px")};
  height: ${({ breakpoint }) => (breakpoint === "xs" ? "300px" : "250px")};
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

export const AddToCartButtonContainer = styled.div`
  position: absolute;
  left: 50%;
  transform: translateX(-55%) translateY(-50%);
`;
