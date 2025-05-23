export const theme = {
  fonts: {
    body: "'Red Hat Text', sans-serif",
  },
  fontSizes: {
    body: "16px",
  },
  margins: {
    sm: "8px",
  },
  colors: {
    grey: {
      500: "#BCB0AD",
      700: "#544946",
    },
    coral: {
      500: "#C66C50",
      700: "#C63B0F",
    },
    brown: {
      500: "#BC7863",
    },
    background: "#fbf3f1",
    white: "white",
  },
  fontWeights: {
    regular: 400,
    semiBold: 600,
    bold: 700,
  },
};

export type ThemeType = typeof theme;
