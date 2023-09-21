'use client';

import { Link } from '@nextui-org/link';

import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from '@nextui-org/navbar';
import React from 'react';
import LoginModal from './modal/loginModal';
import SignupModal from './modal/signupModal';
import useAuth from '../hook/useAuth';
import { Avatar, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger } from '@nextui-org/react';
import { useDispatch } from 'react-redux';
import { logout } from '../redux/slices/userSlice';
import { usePathname, useRouter } from 'next/navigation';

const Nav = () => {
  const { isLoggedIn } = useAuth();
  const dispatch = useDispatch();
  const router = useRouter();
  const pathname = usePathname();

  const handleLogout = () => {
    dispatch(logout());
    router.push('/');
  };

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
        <>
          <NavbarContent className="flex gap-4" justify="center">
            <NavbarItem isActive={pathname === '/questions'}>
              <Link
                href="/questions"
                color={pathname === '/questions' ? 'primary' : 'foreground'}
                aria-current="page"
              >
                Questions
              </Link>
            </NavbarItem>
            <NavbarItem>
              <Link color={pathname === '/interviews' ? 'primary' : 'foreground'} href="#">
                Interviews
              </Link>
            </NavbarItem>
          </NavbarContent>
          <NavbarContent justify="end">
            <NavbarItem>
              <Dropdown placement="bottom-end">
                <DropdownTrigger>
                  <Avatar
                    showFallback
                    src="https://images.unsplash.com/broken"
                    isBordered
                    as="button"
                    color="primary"
                  />
                </DropdownTrigger>
                <DropdownMenu aria-label="Profile Actions" variant="flat">
                  <DropdownItem key="profile" color="primary">
                    <Link href="/profile" className="text-white text-sm w-full">
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
        </>
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
