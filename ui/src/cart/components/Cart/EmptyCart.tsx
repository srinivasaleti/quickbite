import { useTheme } from "styled-components";
import { FlexBox, Text } from "../../../common";
import emptyCart from "../../../assets/empty-cart.svg";
import { Image, ImageContainer } from "./Styled";

export const EmptyCart = () => {
  const { fontWeights, colors } = useTheme();
  return (
    <ImageContainer>
      <FlexBox>
        <Image src={emptyCart} />
        <Text color={colors.grey[500]} weight={fontWeights.semiBold} size="12">
          Your added items will appear here
        </Text>
      </FlexBox>
    </ImageContainer>
  );
};
