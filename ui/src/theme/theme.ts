export const theme = {
  fonts: {
    body: "'Red Hat Text', sans-serif",
  },
  fontSizes: {
    productName: "16px",
  },
  margins: {
    sm: "8px",
  },
  colors: {
    grey: {
      500: "#BCB0AD",
      700: "#544946",
    },
    brown: {
      500: "#BC7863",
    },
  },
  fontWeights: {
    regular: 400,
    semiBold: 600,
    bold: 700,
  },
};

export type ThemeType = typeof theme;
