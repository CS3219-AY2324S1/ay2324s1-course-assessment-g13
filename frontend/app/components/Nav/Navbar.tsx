'use client';
import { Link } from '@nextui-org/link';

import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from '@nextui-org/navbar';
import React from 'react';
import LoginButton from './LoginButton';
import { useSelector } from 'react-redux';
import { RootState } from '../../libs/redux/store';
import useAuth from '../../auth/hooks/useAuth';
import LogoutButton from './LogoutButton';

const Nav = () => {
  const user = useSelector((state: RootState) => state.user)
  const isLoggedIn = user.userId !== 0
  const {handleLogout} = useAuth();
  return (
    <Navbar
      isBordered
      maxWidth="xl"
    >
      <NavbarBrand>
        <Link className="font-bold text-inherit">
          PeerPrep
        </Link>
      </NavbarBrand>
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
      <NavbarContent justify="end">
        <NavbarItem>
          {isLoggedIn 
            ? <LogoutButton handleLogout={handleLogout}/>
            : <LoginButton /> 
          }
        </NavbarItem>
        <NavbarItem className="hidden lg:flex">
          
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
};

export default Nav;
