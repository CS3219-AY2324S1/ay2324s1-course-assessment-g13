'use client';

import { Button, Divider, Input } from '@nextui-org/react';
import RootLayout from '../../layout';
import ProfileLayout from '../layout';
import { Controller, useForm } from 'react-hook-form';
import { GET, PUT } from '../../libs/axios/axios';
import useAuth from '../../hooks/useAuth';
import { notifyError, notifySuccess } from '../../components/toast/notifications';
import DeleteAccountModal from '../../components/modal/deleteAccountModal';
import { useDispatch } from 'react-redux';
import { updateUser } from '../../libs/redux/slices/userSlice';

interface ChangePassword {
  oldPassword: string;
  newPassword: string;
  confirmPassword: string;
}

const AccountPage = () => {
  const dispatch = useDispatch();
  const { userId:id, userRole } = useAuth();
  const initialValues: ChangePassword = {
    oldPassword: '',
    newPassword: '',
    confirmPassword: '',
  };
  const {
    control,
    handleSubmit,
    getValues,
    reset,
    formState: { dirtyFields, errors },
  } = useForm({
    defaultValues: initialValues,
  });

  const allFieldsDirty = Object.keys(dirtyFields).length === Object.keys(initialValues).length;

  const onSubmit = handleSubmit(async (data: ChangePassword) => {
    try {
      const response = await PUT(`/users/password/${id}`, {
        oldPassword: data.oldPassword,
        newPassword: data.newPassword,
      });

      notifySuccess(response.data);
      reset();
    } catch (error) {
      notifyError(error.message.data);
    }
  });

  const handleChangeRole = async () => {
    try {
      const response = await GET(`/auth/user/${userRole === 'user' ? "upgrade" : "downgrade"}`);
      dispatch(updateUser(response.data.user))
      notifySuccess(response.data.message)
    } catch (error) {
      notifyError(error.message.data);
    }
  }

  return (
    <div className="flex flex-col">
      <h1 className="text-2xl">Change Password</h1>
      <Divider className="my-2 " />
      <form>
        <p className="text-start font-medium text-sm py-2 md:text-base">Old password</p>
        <Controller
          name="oldPassword"
          control={control}
          render={({ field }) => (
            <Input
              {...field}
              radius="sm"
              size="lg"
              type="password"
              className="col-span-3 pb-2"
              errorMessage={errors.oldPassword?.message as string}
            />
          )}
        />
        <p className="text-start font-medium text-sm py-2 md:text-base">New password</p>
        <Controller
          name="newPassword"
          control={control}
          rules={{
            minLength: {
              value: 8,
              message: 'Password must be at least 8 characters long',
            },
          }}
          render={({ field }) => (
            <Input
              {...field}
              radius="sm"
              size="lg"
              type="password"
              className="col-span-3 pb-2"
              errorMessage={errors.newPassword?.message as string}
            />
          )}
        />
        <p className="text-start font-medium text-sm py-2 md:text-base">Confirm new password</p>
        <Controller
          name="confirmPassword"
          control={control}
          rules={{
            validate: value => value === getValues().newPassword || 'The passwords do not match',
          }}
          render={({ field }) => (
            <Input
              {...field}
              radius="sm"
              size="lg"
              type="password"
              className="col-span-3 pb-2"
              errorMessage={errors.confirmPassword?.message as string}
            />
          )}
        />
      </form>

      <Divider className="my-4 " />
      <div className="flex flex-row self-end justify-end pb-4">
        <Button color="primary" onClick={onSubmit} className=" my-1 " isDisabled={!allFieldsDirty}>
          Save Changes
        </Button>
      </div>
      <h1 className="text-2xl font-bold text-blue-500">{userRole === 'user' ? "Upgrade" : "Downgrade"} Role</h1>
      <Divider className="my-2 " />
      <p className="text-start font-medium text-sm py-2  md:text-base">
        {userRole === 'user' 
          ? "Upgrade account to gain access to add/edit/delete questions"
          : "Downgrade account only allow you to veiw the questions"
        }
      </p>
      <div className="flex flex-row self-start pb-4">
        <Button color='primary' onClick={handleChangeRole}>
          {userRole === 'user' ? "Upgrade" : "Downgrade"} Role
        </Button>
      </div>
      <h1 className="text-2xl font-bold text-red-500">Delete account</h1>
      <Divider className="my-2 " />
      <p className="text-start font-medium text-sm py-2  md:text-base">
        Deleting your account will remove all your information from our database. This cannot be
        undone.
      </p>
      <div className="flex flex-row self-start pb-4">
        <DeleteAccountModal />
      </div>
    </div>
  );
};

AccountPage.getLayout = page => {
  <RootLayout>
    <ProfileLayout>{page}</ProfileLayout>
  </RootLayout>;
};

export default AccountPage;
