import type { NextPage } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import { Box, Button, Text, Stack, Container } from "@chakra-ui/react";
import { AiFillGithub } from "react-icons/ai";

import Layout from "@/components/Layout";

// Constants //
const TITLE = "Recall";
const SUB_TITLE = "Learn better with flashcards";
const BODY =
  "Active recall is a proven learning technique which allows you to remember and understand what you have learnt effectively. One simple way to practice active recall is through the use of flashcards!";
const GITHUB_LINK = "https://github.com/royleochan/recall";
const BUTTON_TEXTS = ["Get Started", "View Source"];

// Other Components //
const UndrawSvg = () => (
  <Box w={[200, 300, 300, 500]} h={[200, 300, 300, 500]} position="relative">
    <Image alt="about" src="/about.svg" layout="fill" />
  </Box>
);

// Page Component //
const About: NextPage = () => {
  const router = useRouter();

  const navigateToFlashcards = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    router.push("/flashcards");
  };

  const openSourceInNewTab = (e: { preventDefault: () => void }) => {
    const newWindow = window.open(GITHUB_LINK, "_blank", "noopener,noreferrer");
    if (newWindow) newWindow.opener = null;
  };

  return (
    <Layout>
      <Stack
        direction={["column", "row"]}
        paddingX={20}
        paddingTop={10}
        mt={30}
        justify="center"
      >
        <Box>
          <Text fontWeight="extrabold" fontSize="xxxl" color="primary.main.400">
            {TITLE}
          </Text>
          <Text fontWeight="bold" fontSize="lg">
            {SUB_TITLE}
          </Text>
          <Container padding={0} marginTop={4} fontFamily="paragraph">
            {BODY}
          </Container>
          <Stack spacing={4} direction="row" align="center" marginTop={10}>
            <Button
              colorScheme="green"
              size="md"
              onClick={navigateToFlashcards}
            >
              {BUTTON_TEXTS[0]}
            </Button>
            <Button
              leftIcon={<AiFillGithub />}
              colorScheme="purple"
              variant="outline"
              onClick={openSourceInNewTab}
            >
              {BUTTON_TEXTS[1]}
            </Button>
          </Stack>
        </Box>
        <UndrawSvg />
      </Stack>
    </Layout>
  );
};

export default About;
