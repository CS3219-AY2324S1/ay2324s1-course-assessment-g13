'use client';

import { Controller, useForm } from 'react-hook-form';
import RootLayout from '../../layout';
import ProfileLayout from '../layout';
import useAuth from '../../hook/useAuth';
import { Avatar, Button, Divider, Input } from '@nextui-org/react';
import ImageUpload from '../../components/form/ImageUpload';
import { PUT } from '../../libs/axios/axios';
import { useDispatch, useSelector } from 'react-redux';
import { AppState } from '../../redux/store';
import { notifyError, notifySuccess } from '../../components/Notifications';
import { updateUser } from '../../redux/slices/userSlice';

interface UserInfo {
  username: string;
  photoUrl: string;
}

const InfoPage = () => {
  const dispatch = useDispatch();
  const { id, username } = useAuth();
  const photoUrl = useSelector((state: AppState) => state.user.photoUrl);

  const initialValues: UserInfo = {
    username,
    photoUrl,
  };
  const {
    control,
    handleSubmit,
    setValue,
    watch,
    formState: { isDirty },
  } = useForm({
    defaultValues: initialValues,
  });

  const onSubmit = handleSubmit(async (data: UserInfo) => {
    try {
      const response = await PUT(`/users/info/${id}`, data);
      dispatch(updateUser({ ...data }));
      notifySuccess(response.data);
    } catch (error) {
      notifyError(error.message.data);
    }
  });

  const photo = watch('photoUrl');

  return (
    <>
      <form>
        <div className="flex flex-col gap-2 md:grid md:grid-cols-4">
          <p className="text-start font-medium text-sm md:text-base">Username</p>
          <Controller
            name="username"
            control={control}
            render={({ field }) => (
              <Input {...field} radius="sm" size="lg" className="col-span-3" />
            )}
          />
        </div>
        <Divider className="my-4 " />
        <div className="flex flex-col gap-2 md:grid md:grid-cols-4">
          <p className="text-start font-medium text-sm md:text-base">Photo</p>
          <div className="flex flex-col md:grid md:grid-cols-3 md:col-span-3">
            <Avatar
              showFallback
              src={photo}
              isBordered
              color="primary"
              className="h-20 w-20 md:h-32 md:w-32 self-center justify-self-center md: mb-4"
            />
            <div className="col-span-2 col-start-2">
              <Controller
                name="photoUrl"
                control={control}
                render={({ field }) => (
                  <ImageUpload
                    setImage={value => {
                      return setValue('photoUrl', value, { shouldDirty: true });
                    }}
                    {...field}
                  />
                )}
              />
            </div>
          </div>
        </div>
      </form>
      <Divider className="my-4 " />
      <div className="flex flex-row self-end justify-end">
        <Button color="primary" onClick={onSubmit} className=" my-1 " isDisabled={!isDirty}>
          Save Changes
        </Button>
      </div>
    </>
  );
};

InfoPage.getLayout = page => {
  <RootLayout>
    <ProfileLayout>{page}</ProfileLayout>
  </RootLayout>;
};

export default InfoPage;
