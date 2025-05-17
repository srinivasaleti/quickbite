import styled from "styled-components";

interface FlexBoxProps {
  direction?: "row" | "column";
  justifyContent?: string;
  alignItems?: string;
  gap?: string | number;
  wrap?: "nowrap" | "wrap" | "wrap-reverse";
  children?: React.ReactNode;
}

export const FlexBox = styled.div<FlexBoxProps>`
  display: flex;
  flex-direction: ${({ direction = "column" }) => direction};
  justify-content: ${({ justifyContent = "flex-start" }) => justifyContent};
  align-items: ${({ alignItems = "stretch" }) => alignItems};
  gap: ${({ gap = 0 }) => (typeof gap === "number" ? `${gap}px` : gap)};
  flex-wrap: ${({ wrap = "nowrap" }) => wrap};
`;
