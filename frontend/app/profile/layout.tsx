'use client';

import { Avatar, Divider } from '@nextui-org/react';
import useAuth from '../hooks/useAuth';
import { useSelector } from 'react-redux';
import { AppState } from '../libs/redux/store';
import { usePathname } from 'next/navigation';
import Link from 'next/link';

const profileSideBar = [
  { name: 'Info', path: '/profile/info' },
  { name: 'Account', path: '/profile/account' },
];

const ProfileLayout = ({ children }) => {
  const photoUrl = useSelector((state: AppState) => state.user.photoUrl);
  const { username } = useAuth();
  const pathname = usePathname();

  return (
    <div className="flex flex-row justify-center px-2">
      <div className="flex flex-col max-w-[1280px] w-full ">
        <div className="flex flex-row items-center p-4">
          <Avatar
            showFallback
            isBordered
            color="primary"
            className="h-12 w-12 self-center"
            src={photoUrl}
          />
          <p className="self-center py-2 px-4 text-2xl md:justify-self-start">{username}</p>
        </div>
        <Divider className="my-2" />
        <div className="grid grid-cols-[max-content_1fr]">
          <div className="flex flex-col auto-cols-max py-2 mr-12 ml-3">
            {profileSideBar.map(item => (
              <Link
                className={`px-2 py-1 self-start rounded w-full ${
                  pathname === item.path ? 'bg-gray-800' : ''
                }`}
                key={item.name}
                href={item.path}
              >
                {item.name}
              </Link>
            ))}
          </div>
          <div className="p-4 auto-cols-fr">{children}</div>
        </div>
      </div>
    </div>
  );
};

export default ProfileLayout;
