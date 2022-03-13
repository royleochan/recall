import { ThemeConfig } from "@chakra-ui/react";
import { colors } from "./colors";
import { breakpoints } from "./breakpoints";
import { components } from "./components";
import { fonts, fontSizes, fontWeights } from "./typography";

export const customTheme: Object = {
  colors,
  fonts,
  fontSizes,
  fontWeights,
  components,
  breakpoints,
};

export const themeConfig: ThemeConfig = {
  initialColorMode: "dark",
  useSystemColorMode: false,
};
