'use client';

import { Avatar, Divider } from '@nextui-org/react';
import useAuth from '../hook/useAuth';
import { useSelector } from 'react-redux';
import { AppState } from '../redux/store';

const profileSideBar = [
  { name: 'Info', path: '/info' },
  { name: 'Account', path: '/account' },
];

const ProfileLayout = ({ children }) => {
  const photoUrl = useSelector((state: AppState) => state.user.photoUrl);
  const { username } = useAuth();

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
              <aside className="p-2 self-start rounded" key={item.name}>
                {item.name}
              </aside>
            ))}
          </div>
          <div className="p-4 auto-cols-fr">{children}</div>
        </div>
      </div>
    </div>
  );
};

export default ProfileLayout;
