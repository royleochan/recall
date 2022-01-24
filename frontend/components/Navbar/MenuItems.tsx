import React, { ReactNode } from "react";
import { Box, Link, Stack, useColorModeValue } from "@chakra-ui/react";
import NextLink from "next/link";

import ToggleColor from "./ToggleColor";
import Profile from "./Profile";

interface MenuItemProps {
  children: ReactNode;
  target: string;
}

const MenuItem = ({ children, target = "/" }: MenuItemProps) => {
  const color = useColorModeValue("black", "white");

  return (
    <NextLink href={target} passHref>
      <Link fontSize="lg" fontWeight="bold" color={color}>
        {children}
      </Link>
    </NextLink>
  );
};

interface MenuItemsProps {
  isOpen: boolean;
}

const MenuItems = ({ isOpen }: MenuItemsProps) => {
  return (
    <Box
      display={{ base: isOpen ? "block" : "none", md: "block" }}
      flexBasis={{ base: "100%", md: "auto" }}
    >
      <Stack
        spacing={8}
        align="center"
        justify="center"
        direction={["column", "column", "row", "row"]}
        pt={[4, 4, 0, 0]}
      >
        <MenuItem target="/flashcards">Flashcards</MenuItem>
        <ToggleColor />
        <Profile />
      </Stack>
    </Box>
  );
};

export default MenuItems;
