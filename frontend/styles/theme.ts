import { ThemeConfig, theme as baseTheme } from "@chakra-ui/react";

export const customTheme: Object = {
  colors: {
    light: baseTheme.colors.yellow["300"],
    dark: baseTheme.colors.gray["700"],
    primary: {
      main: baseTheme.colors.red,
    },
    secondary: {
      main: baseTheme.colors.green,
    },
  },
  fontSizes: {
    xs: "0.75rem",
    sm: "0.875rem",
    md: "1rem",
    lg: "1.125rem",
    xl: "1.25rem",
    xxl: "2.5rem",
    xxxl: "6rem",
  },
  fonts: {
    paragraph: "avenir",
  },
  fontWeights: {
    light: 300,
    normal: 400,
    medium: 500,
    bold: 700,
    extrabold: 800,
  },
  lineHeights: {
    normal: "normal",
    none: 1,
    short: 1.375,
    base: 1.5,
    tall: 1.625,
  },
  letterSpacings: {
    tight: "-0.025em",
    normal: "0",
    wide: "0.025em",
  },
};

export const themeConfig: ThemeConfig = {
  initialColorMode: "dark",
  useSystemColorMode: false,
};
