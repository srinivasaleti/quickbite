import styled from "styled-components";
import { Text } from "../../../common";

export const CART_CONTAINER_WIDTH = "400px";

export const CartContainer = styled.div<{ showFullWithContainer: boolean }>`
  min-width: ${({ showFullWithContainer }) =>
    showFullWithContainer ? "100%" : CART_CONTAINER_WIDTH};
  max-width: ${({ showFullWithContainer }) =>
    showFullWithContainer ? "100%" : CART_CONTAINER_WIDTH};
  box-sizing: border-box;
  gap: 12px;
  padding: 24px;
  background: ${({ theme }) => theme.colors.white};
`;

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

export const OrderSummaryContainer = styled.div<{
  showFullWidthContainer: boolean;
}>`
  width: ${({ showFullWidthContainer }) =>
    showFullWidthContainer ? "100%" : "350px"};
  font-family: Arial, sans-serif;
  margin-top: 60px;
`;

export const OrderSummaryIemRow = styled.div`
  margin-bottom: 15px;
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

export const OrderSummaryLeftPart = styled.div`
  flex-grow: 1;
`;

export const OrderSummaryItemName = styled(Text)`
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  margin-bottom: 6px;
`;

export const OrderSummaryQuantity = styled(Text)`
  color: ${({ theme }) => theme.colors.brown["500"]};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  margin-right: 8px;
`;

export const PriceInfo = styled(Text)`
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  margin-right: 6px;
`;

export const DeleteBtn = styled.button`
  background: none;
  border: 1px solid #b7a19b;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  color: #b7a19b;
  font-weight: 700;
  cursor: pointer;
  line-height: 18px;
`;

export const OrderTotal = styled(Text)`
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  color: ${({ theme }) => theme.colors.grey["500"]};
`;

export const ConfirmBtn = styled.button<{ disabled: boolean }>`
  margin-top: 25px;
  width: 100%;
  padding: 14px;
  background-color: ${({ theme, disabled }) =>
    disabled ? theme.colors.coral["500"] : theme.colors.coral["700"]};
  border: none;
  border-radius: 25px;
  color: ${({ theme, disabled }) =>
    disabled ? theme.colors.grey : theme.colors.white};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  font-size: ${({ theme }) => theme.fontSizes.body};
  cursor: pointer;
`;

export const HorizontalLine = styled.hr`
  border: none;
  height: 1px;
  background-color: ${({ theme }) => theme.colors.grey["500"]};
  margin: 10px 0;
`;
