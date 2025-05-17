import styled from "styled-components";
import type { Breakpoint } from "../../common/hooks/useBreakpoints";
import { FlexBox } from "../../common";

interface Props {
  breakpoint: Breakpoint;
}

export const HomeContainer = styled(FlexBox)<Props>`
  padding: ${({ breakpoint }) =>
    breakpoint === "xl" ? "100px" : breakpoint === "l" ? "50px" : "20px"};
`;
