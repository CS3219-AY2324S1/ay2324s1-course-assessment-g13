'use client';

import { Avatar, Divider, Input } from '@nextui-org/react';
import { Button } from '@nextui-org/button';
import useAuth from '../hook/useAuth';
import { Controller, useForm } from 'react-hook-form';
import ImageUpload from '../components/form/ImageUpload';

const Profile = () => {
  const { username } = useAuth();

  const { control, handleSubmit, setValue, watch } = useForm();

  const onSubmit = handleSubmit(data => {
    console.log(data);
  });
  const photo = watch('photo');
  return (
    <div className="flex flex-row justify-center items-center text-center px-4">
      <div className="flex flex-col my-4 max-w-[1280px] w-full">
        <div className="flex flex-col py-4 md:grid md:grid-cols-4">
          <Avatar
            showFallback
            isBordered
            color="primary"
            className="h-20 w-20 self-center md:h-32 md:w-32"
            radius="sm"
          />
          <p className="self-center py-4 text-2xl md:justify-self-start">{username}</p>
        </div>
        <Divider className="my-4" />
        <form>
          <div className="flex flex-col gap-2 md:grid md:grid-cols-4">
            <p className="text-start font-medium text-sm md:text-base">Username</p>
            <Controller
              name="username"
              control={control}
              render={({ field }) => (
                <Input {...field} value={username} radius="sm" size="lg" className="col-span-3" />
              )}
            />
          </div>
          <Divider className="my-4 " />
          <div className="flex flex-col gap-2 md:grid md:grid-cols-4">
            <p className="text-start font-medium text-sm md:text-base">Photo</p>
            <div className="grid grid-cols-3 col-span-3">
              <Avatar
                showFallback
                src={photo}
                isBordered
                color="primary"
                className="h-20 w-20 md:h-32 md:w-32 self-center justify-self-center"
              />
              <div className="col-span-2 col-start-2">
                <Controller
                  name="photo"
                  control={control}
                  render={({ field }) => (
                    <ImageUpload onChange={value => setValue('photo', value)} {...field} />
                  )}
                />
              </div>
            </div>
          </div>
        </form>
        <Button color="primary" onClick={onSubmit}>
          Add
        </Button>
      </div>
    </div>
  );
};

export default Profile;
