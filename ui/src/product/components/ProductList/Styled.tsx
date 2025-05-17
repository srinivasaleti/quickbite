import styled from "styled-components";
import { Text } from "../../../common";

export const Grid = styled.div`
  display: flex;
  flex-wrap: wrap;
  col-gap: 12px;
  row-gap: 20px;
  margin-top: 30px;
`;

export const NoProductsMessage = styled(Text)`
  text-align: center;
  margin-top: 50px;
  font-size: 18px;
  color: #555;
`;

export const NoProductsContainer = styled.div`
  text-align: center;
  font-size: 18px;
  color: #555;
  display: flex;
  justify-content: center;
  height: 100vh;
  flex-direction: column;
  align-items: center;
  gap: 10px;
`;

export const Icon = styled.div`
  font-size: 50px;
`;
