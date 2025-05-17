import { useState } from "react";
import styled from "styled-components";
import { FlexBox, Text } from "../../../common";
import { useCart } from "../../CartContext";

const LinkText = styled.p`
  color: ${({ theme }) => theme.colors.coral["500"]};
  cursor: pointer;
  text-decoration: underline;
  margin: 0;
`;

const Input = styled.input`
  padding: 10px 12px;
  flex: 1;
  font-size: ${({ theme }) => theme.fontSizes.body};
  font-weight: ${({ theme }) => theme.fontWeights.semiBold};
  border: 1px solid #ccc;
  border-radius: 5px;
`;

const AddButton = styled.button`
  padding: 8px 12px;
  margin: auto;
  background-color: ${({ theme }) => theme.colors.coral["500"]};
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  width: 100px;
`;

export const CouponCodeInput = ({
  onCouponRemoved,
}: {
  onCouponRemoved: () => void;
}) => {
  const [showInput, setShowInput] = useState(false);
  const { coupon, setCoupon } = useCart();

  const handleApply = () => {
    if (!coupon.trim()) return;
    setCoupon(coupon.trim());
    setShowInput(false);
  };

  const handleRemove = () => {
    setCoupon("");
    setShowInput(false);
    onCouponRemoved();
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleApply();
    }
  };

  return (
    <div>
      {!showInput && !coupon && (
        <LinkText onClick={() => setShowInput(true)}>
          You have coupon code?
        </LinkText>
      )}

      {showInput && (
        <FlexBox gap={"16px"}>
          <Input
            type="text"
            placeholder="Enter coupon code"
            value={coupon}
            onChange={(e) => setCoupon(e.target.value)}
            onKeyDown={handleKeyDown}
          />
          <AddButton onClick={handleApply}>Add</AddButton>
        </FlexBox>
      )}

      {coupon && !showInput && (
        <FlexBox gap={"16px"}>
          <Text>
            Coupon Added to order: <strong>{coupon}</strong>
          </Text>
          <LinkText onClick={handleRemove}>Remove coupon code?</LinkText>
        </FlexBox>
      )}
    </div>
  );
};
