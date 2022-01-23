import React, { ReactNode } from "react";
import { VStack } from "@chakra-ui/react";
import Navbar from "../Navbar";

interface PageProps {
  children: ReactNode;
}

const Page = ({ children }: PageProps) => {
  return (
    <VStack>
      <Navbar />
      {children}
    </VStack>
  );
};

export default Page;
