import type { NextPage } from "next";
import { HStack, Box, Text } from "@chakra-ui/react";

import Layout from "@/components/Layout";
import SiderMenu from "@/components/SiderMenu";
import Card from "@/components/Card";

const Flashcards: NextPage = () => {
  return (
    <Layout>
      <HStack alignSelf="flex-start">
        <SiderMenu />
        <Box pl="400px">
          <Text fontSize="xxl" fontWeight="bold">
            CS3230
          </Text>
          <Box>
            <Card />
          </Box>
        </Box>
      </HStack>
    </Layout>
  );
};

export default Flashcards;
