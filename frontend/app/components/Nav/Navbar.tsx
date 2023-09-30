'use client';
import { Link } from '@nextui-org/link';

import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from '@nextui-org/navbar';
import React, { useEffect, useState } from 'react';
import LoginButton from './LoginButton';
import useAuth from '../../hooks/useAuth';
import SignUpButton from './SignUpButton';
import { Avatar, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger } from '@nextui-org/react';
import { useDispatch } from 'react-redux';
import { useRouter } from 'next/navigation';
import { logout } from '../../libs/redux/slices/userSlice';
import { GET } from '../../libs/axios/axios';

const Nav = () => {
  const dispatch = useDispatch();
  const router = useRouter();
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const { isAuthenticated } = useAuth();

  const handleLogout = async () => {
    try {
        dispatch(logout());
        router.push('/');
        await GET('/auth/logout');
    } catch (error) {
        console.error(error)
    }
  }

  useEffect(() => {
    setIsLoggedIn(isAuthenticated);
  }, [isAuthenticated])

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
      <>
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
      </>
      }
      {isLoggedIn ? 
        <NavbarContent justify="end">
          <NavbarItem>
            <Dropdown placement="bottom-end">
              <DropdownTrigger>
                <Avatar showFallback isBordered as="button" color="primary" />
              </DropdownTrigger>
              <DropdownMenu aria-label="Profile Actions" variant="flat">
                <DropdownItem key="profile" color="primary">
                  <Link href="/profile/info" className="text-white text-sm w-full">
                    Profile
                  </Link>
                </DropdownItem>
                <DropdownItem key="logout" color="danger" onClick={handleLogout}>
                  Log Out
                </DropdownItem>
              </DropdownMenu>
            </Dropdown>
          </NavbarItem>
        </NavbarContent>
        :
        <NavbarContent justify="end">
          <NavbarItem className="hidden lg:flex gap-3">
            <LoginButton />
            <SignUpButton />
          </NavbarItem>
        </NavbarContent>
      }
    </Navbar>
  );
};

export default Nav;
