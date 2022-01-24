import React from "react";
import { Box, Text, useColorModeValue } from "@chakra-ui/react";

const SiderMenu = () => {
  const bgColor = useColorModeValue("#EDF2F7", "#2D3748");

  return (
    <Box
      background={bgColor}
      height="92vh"
      width="20vh"
      position="fixed"
      left={0}
      bottom={0}
      padding={6}
    >
      <Text>CS3230</Text>
    </Box>
  );
};

export default SiderMenu;
