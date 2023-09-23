import { Link } from '@nextui-org/link';

import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from '@nextui-org/navbar';
import React from 'react';
import LoginModal from './modal/loginModal';
import SignupModal from './modal/signupModal';

const Nav = () => {
  // TODO: Check login status here
  const isLoggedIn = false;
  return (
    <Navbar
      isBordered
      maxWidth="xl"
      height="10vh"
      classNames={{
        item: [
          'flex',
          'relative',
          'h-full',
          'items-center',
          "data-[active=true]:after:content-['']",
          'data-[active=true]:after:absolute',
          'data-[active=true]:after:bottom-0',
          'data-[active=true]:after:left-0',
          'data-[active=true]:after:right-0',
          'data-[active=true]:after:h-[3px]',
          'data-[active=true]:after:rounded-[2px]',
          'data-[active=true]:after:bg-primary',
        ],
      }}
    >
      <NavbarBrand>
        <Link href={isLoggedIn ? '/questions' : '/'} className="font-bold text-inherit">
          PeerPrep
        </Link>
      </NavbarBrand>
      {isLoggedIn && (
        <NavbarContent className="hidden sm:flex gap-4" justify="center">
          <NavbarItem isActive>
            <Link href="/questions" aria-current="page">
              Questions
            </Link>
          </NavbarItem>
          <NavbarItem>
            <Link color="foreground" href="#">
              Interviews
            </Link>
          </NavbarItem>
        </NavbarContent>
      )}
      {!isLoggedIn && (
        <NavbarContent justify="end">
          <NavbarItem>
            <LoginModal />
          </NavbarItem>
          <NavbarItem className="hidden lg:flex">
            <SignupModal isNav={true} />
          </NavbarItem>
        </NavbarContent>
      )}
    </Navbar>
  );
};

export default Nav;
