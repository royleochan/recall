import React from "react";
import {
  Box,
  IconButton,
  Text,
  useDisclosure,
  Drawer,
  DrawerContent,
  DrawerHeader,
  DrawerBody,
  DrawerFooter,
} from "@chakra-ui/react";
import { ArrowRightIcon, ArrowLeftIcon } from "@chakra-ui/icons";

const SiderMenu = () => {
  const { isOpen, onOpen, onClose } = useDisclosure({ defaultIsOpen: true });

  return (
    <>
      {!isOpen && (
        <IconButton
          position="absolute"
          fontSize={18}
          aria-label="Toggle Drawer"
          _focus={{ boxShadow: "none" }}
          onClick={onOpen}
          icon={<ArrowRightIcon />}
        />
      )}
      <Drawer
        isFullHeight={false}
        variant="permanent"
        placement="left"
        isOpen={isOpen}
        onClose={onClose}
      >
        <DrawerContent h="100vh" bottom="0 !important" top="revert !important">
          <DrawerHeader> Categories</DrawerHeader>
          <DrawerBody>
            <Text>CS3230</Text>
          </DrawerBody>
          <DrawerFooter>
            <IconButton
              w="100%"
              fontSize={18}
              aria-label="Toggle Drawer"
              variant="solid"
              bg="transparent"
              _focus={{ boxShadow: "none" }}
              onClick={onClose}
              icon={<ArrowLeftIcon />}
            />
          </DrawerFooter>
        </DrawerContent>
      </Drawer>
    </>
  );
};

export default SiderMenu;
