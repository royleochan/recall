import React, { useState } from "react";
import { Box, Text } from "@chakra-ui/react";

const Card = () => {
  const [isFlipped, setIsFlipped] = useState(false);

  const flipCard = () => setIsFlipped(!isFlipped);

  return (
    <Box
      my={6}
      p={8}
      w="400px"
      borderWidth={1}
      borderRadius={10}
      borderColor={isFlipped ? "secondary.main.300" : "primary.main.300"}
      justifyContent="center"
      alignItems="center"
      cursor="pointer"
      onClick={flipCard}
    >
      {!isFlipped ? (
        <Text
          fontSize="xl"
          fontWeight="bold"
          color="primary.main.400"
          align="center"
        >
          How to identify dynamic programming problems?
        </Text>
      ) : (
        <Text fontSize="md" fontWeight="med" align="center">
          Dynamic programming problems are problems with overlapping subproblems
          and optimal substructure property
        </Text>
      )}
    </Box>
  );
};

export default Card;
