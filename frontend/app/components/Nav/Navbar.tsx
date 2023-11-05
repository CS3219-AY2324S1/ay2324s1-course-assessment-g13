'use client';
import { Link } from '@nextui-org/link';
import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from '@nextui-org/navbar';
import React, { useEffect } from 'react';
import useAuth from '../../hooks/useAuth';
import { Dropdown, DropdownItem, DropdownMenu, DropdownTrigger } from '@nextui-org/dropdown'
import { Avatar } from '@nextui-org/avatar';
import { Button } from '@nextui-org/button';
import { useDispatch, useSelector } from 'react-redux';
import { AppState } from '../../libs/redux/store';
import { useRouter } from 'next/navigation';
import { logout as UserLogout } from '../../libs/redux/slices/userSlice';
import { logout as AuthLogout } from '../../libs/redux/slices/authSlice';
import { GET } from '../../libs/axios/axios';
import { usePathname } from 'next/navigation';
import { setIsLeaving, setIsChatOpen, selectCollabState } from '../../libs/redux/slices/collabSlice';
import { ChatIcon } from '../../../public/ChatIcon';
import { signOut } from 'next-auth/react';
import { notifyError } from '../toast/notifications';

const Nav = () => {
  const collabState = useSelector(selectCollabState);
  const dispatch = useDispatch();
  const router = useRouter();
  const { status, isLoggedIn } = useAuth();
  const photoUrl = useSelector((state: AppState) => state.user.photoUrl);
  const pathname = usePathname();

  const handleLogout = async () => {
    if (!isLoggedIn) return;
    try {
        dispatch(UserLogout());
        dispatch(AuthLogout());
        await GET('/auth/logout');
        router.push('/');
        signOut();
    } catch (error) {
      const message =  error.message.data.message;
      notifyError(message);
    }
  }

  useEffect(() => {
    status === "unauthenticated" && handleLogout();
  }, [status])

  const checkPath = (url : string) => {
    return pathname === url;
  }

  return (
    <Navbar
      isBordered
      maxWidth="xl"
    >
      <NavbarBrand>
        <Link href={isLoggedIn && !checkPath("/collab") ? "/questions" : "#"} className="font-bold text-inherit">
          PeerPrep
        </Link>
      </NavbarBrand>
      {(isLoggedIn && !checkPath("/collab")) &&
      <>
        <NavbarContent justify="center">
          <NavbarItem isActive>
            <Link color={checkPath("/questions") ? 'primary' : 'foreground'} href="/questions" aria-current="page">
              Questions
            </Link>
          </NavbarItem>
          <NavbarItem>
            <Link color={checkPath("/interviews") ? 'primary' : 'foreground'} href="/interviews">
              Interviews
            </Link>
          </NavbarItem>
        </NavbarContent>
      </>
      }
      {checkPath("/collab") && (
        <NavbarContent justify="center">
          <NavbarItem>
            <Button
              color="danger" 
              variant="solid" 
              className="text-lg" 
              onPress={() => dispatch(setIsLeaving(true))}
            >
              End Collaboration
            </Button>
          </NavbarItem>
        </NavbarContent>
        
      )}

      {isLoggedIn ? 
        <NavbarContent justify="end">
          {checkPath("/collab") && (
            <NavbarItem>
              <Button onPress={() => dispatch(setIsChatOpen(!collabState.isChatOpen))} className="h-fit min-w-0 px-0 bg-transparent flex item-center">
                <span className="cursor-pointer active:opacity-50">
                  <ChatIcon />
                </span>
              </Button>
            </NavbarItem>
          )}
          <NavbarItem>
            <Dropdown placement="bottom-end">
              <DropdownTrigger>
                <Avatar src={photoUrl} showFallback isBordered as="button" color="primary" />
              </DropdownTrigger>
              {!checkPath("/collab") && (
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
              )}
            </Dropdown>
          </NavbarItem>
        </NavbarContent>
        :
        <NavbarContent justify="end">
          <NavbarItem className="hidden lg:flex gap-3">
            <Button  
                variant="bordered" 
                color="default"
                as={Link}
                href="/login"
            >
                Login
            </Button>
            <Button  
                color="primary"
                as={Link}
                href="/signup"
            >
                Sign Up
            </Button>
          </NavbarItem>
        </NavbarContent>
      }
    </Navbar>
  );
};

export default Nav;
