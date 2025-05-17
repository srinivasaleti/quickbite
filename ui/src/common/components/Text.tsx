import styled from "styled-components";

interface TextProps {
  size?: string;
  weight?: number;
  color?: string;
  children?: React.ReactNode;
}

export const Text = styled.span<TextProps>`
  font-size: ${({ size, theme }) => size || theme.fontSizes.body};
  font-weight: ${({ weight, theme }) => weight || theme.fontWeights.regular};
  color: ${({ color, theme }) => color || theme.colors.grey[700]};
  font-family: ${({ theme }) => theme.fonts.body};
  display: block;
`;

export default Text;
