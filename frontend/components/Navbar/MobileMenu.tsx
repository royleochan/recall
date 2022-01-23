import React, { MouseEventHandler } from "react";
import { Icon, IconButton } from "@chakra-ui/react";
import { AiOutlineMenu, AiOutlineClose } from "react-icons/ai";

interface MenuProps {
  toggle: MouseEventHandler<HTMLButtonElement>;
  isOpen: boolean;
}

const MobileMenu = ({ toggle, isOpen }: MenuProps) => {
  return (
    <IconButton
      aria-label="menu"
      display={{ base: "block", md: "none" }}
      onClick={toggle}
    >
      {isOpen ? (
        <Icon as={AiOutlineClose} fontSize={22} />
      ) : (
        <Icon as={AiOutlineMenu} fontSize={22} />
      )}
    </IconButton>
  );
};

export default MobileMenu;
