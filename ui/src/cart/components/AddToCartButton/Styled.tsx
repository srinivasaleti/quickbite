import styled from "styled-components";
import { Text } from "../../../common";

export const StyledButton = styled.button`
  width: 170px;
  height: 50px;
  border-radius: 20px;
  padding-left: 24px;
  cursor: pointer;
  background: ${({ theme }) => theme.colors.background};
  border: 2px solid ${({ theme }) => theme.colors.grey[500]};
`;

export const QuantityContainer = styled.div`
  background-color: ${({ theme }) => theme.colors.coral[700]};
  border-radius: 30px;
  padding: 10px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 140px;
`;

export const CircleButton = styled.button`
  width: 24px;
  height: 24px;
  border-radius: 50%;
  border: 2px solid white;
  color: white;
  font-size: 20px;
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
`;

export const CountText = styled(Text)`
  color: white;
  font-weight: 600;
  font-size: 18px;
`;
