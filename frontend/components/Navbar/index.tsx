import { Flex } from "@chakra-ui/react";
import React, { useState } from "react";
import Logo from "./Logo";
import MenuItems from "./MenuItems";
import MobileMenu from "./MobileMenu";

const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => setIsOpen(!isOpen);

  return (
    <Flex
      as="nav"
      align="center"
      justify="space-between"
      wrap="wrap"
      width="100%"
      py={4}
      px={12}
    >
      <Logo />
      <MobileMenu toggle={toggle} isOpen={isOpen} />
      <MenuItems isOpen={isOpen} />
    </Flex>
  );
};

export default Navbar;
