import { createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    background: ${({ theme }) => theme.colors.background};
    font-family: ${({ theme }) => {
      return theme.fonts.body;
    }};
  }
`;

export default GlobalStyle;
