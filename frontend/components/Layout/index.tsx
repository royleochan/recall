import React, { ReactNode } from "react";
import Head from "next/head";
import { VStack } from "@chakra-ui/react";

import Navbar from "../Navbar";

interface PageProps {
  children: ReactNode;
}

const Layout = ({ children }: PageProps) => {
  return (
    <>
      <Head>
        <title>Recall | Learning with flashcards</title>
      </Head>
      <main>
        <VStack>
          <Navbar />
          {children}
        </VStack>
      </main>
    </>
  );
};

export default Layout;
