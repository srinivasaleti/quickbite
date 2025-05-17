import styled from "styled-components";
import cartIcon from "../../assets/cart-icon.svg";

export const StickyCartIcon = styled.div`
  position: fixed;
  top: 20px;
  right: 20px;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  color: white;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  z-index: 1000;
  background: white;
`;

export const CartIcon = ({ onClick }: { onClick: () => void }) => (
  <StickyCartIcon onClick={onClick}>{<img src={cartIcon} />}</StickyCartIcon>
);
