'use client';
import { Link } from '@nextui-org/link';

import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from '@nextui-org/navbar';
import React, { useEffect, useState } from 'react';
import LoginButton from './LoginButton';
import { useSelector } from 'react-redux';
import { RootState } from '../../libs/redux/store';
import useAuth from '../../(auth)/hooks/useAuth';
import LogoutButton from './LogoutButton';
import SignUpButton from './SignUpButton';

const Nav = () => {
  const user = useSelector((state: RootState) => state.user)
  const [isLoggedIn, setIsLoggedIn] = useState(true)
  const {handleLogout} = useAuth();

  useEffect(() => {
    setIsLoggedIn(user.userId !== 0);
  })

  return (
    <Navbar
      isBordered
      maxWidth="xl"
    >
      <NavbarBrand>
        <Link href={isLoggedIn ? "/questions" : "/"} className="font-bold text-inherit">
          PeerPrep
        </Link>
      </NavbarBrand>
      {isLoggedIn &&
      <NavbarContent justify="center">
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
      }
      <NavbarContent justify="end">
        <NavbarItem>
        {isLoggedIn
          ? <LogoutButton  handleLogout={handleLogout}/>
          : <LoginButton />
        }
        </NavbarItem>
        {!isLoggedIn &&  
          <NavbarItem className="hidden lg:flex">
            <SignUpButton />
          </NavbarItem>
        }
        </NavbarContent>
    </Navbar>
  );
};

export default Nav;
