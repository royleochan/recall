import type { NextPage } from "next";
import { Box, Text } from "@chakra-ui/react";

import Layout from "@/components/Layout";
import SiderMenu from "@/components/SiderMenu";
import Card from "@/components/Card";

const Flashcards: NextPage = () => {
  return (
    <Layout>
      <SiderMenu />
      <Box alignSelf="flex-start" pl={[12, 300]}>
        <Text fontSize="xxl" fontWeight="bold">
          CS3230
        </Text>
        <Box>
          <Card />
        </Box>
      </Box>
    </Layout>
  );
};

export default Flashcards;
