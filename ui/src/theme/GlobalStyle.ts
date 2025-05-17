import { createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    font-family: ${({ theme }) => {
      return theme.fonts.body;
    }};
  }
`;

export default GlobalStyle;
