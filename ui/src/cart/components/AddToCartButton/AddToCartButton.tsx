import {
  CircleButton,
  CountText,
  QuantityContainer,
  StyledButton,
} from "./Styled";
import cartIcon from "../../../assets/cart-icon.svg";
import { FlexBox, Text } from "../../../common";
import { useTheme } from "styled-components";

export const AddToCartButton = ({
  onAdd,
  count,
  onRemove,
}: {
  onAdd: () => void;
  onRemove: () => void;
  count: number;
}) => {
  const { fontWeights } = useTheme();

  if (count) {
    return (
      <QuantityContainer>
        <CircleButton onClick={onRemove}>−</CircleButton>
        <CountText>{count}</CountText>
        <CircleButton onClick={onAdd}>＋</CircleButton>
      </QuantityContainer>
    );
  }

  return (
    <StyledButton onClick={onAdd}>
      <FlexBox gap={"8px"} direction="row">
        <img src={cartIcon} />
        <Text weight={fontWeights.semiBold}>Add To Cart</Text>
      </FlexBox>
    </StyledButton>
  );
};
