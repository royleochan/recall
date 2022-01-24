import React from "react";
import { Box, Text } from "@chakra-ui/react";
import { useRouter } from "next/router";

const Logo = () => {
  const router = useRouter();

  const navigateToHome = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    router.push("/");
  };

  return (
    <Box w="100px">
      <Text
        onClick={navigateToHome}
        fontSize="xl"
        fontWeight="bold"
        color="primary.main.400"
        cursor="pointer"
      >
        Recall
      </Text>
    </Box>
  );
};

export default Logo;
