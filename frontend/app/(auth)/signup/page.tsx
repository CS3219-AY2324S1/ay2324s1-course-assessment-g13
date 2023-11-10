'use client';

import { signOut, useSession } from 'next-auth/react';
import AuthCard from '../../components/card/AuthCard';
import { useDisclosure } from '@nextui-org/react';
import { Modal, ModalContent, ModalBody, ModalHeader, ModalFooter } from '@nextui-org/modal';
import { Button } from '@nextui-org/button';
import { Input } from '@nextui-org/input';
import { Select, SelectItem } from '@nextui-org/select';
import { LANGUAGES } from '../../constants/languages';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { AuthResponse } from '../login/page';
import { POST } from '../../libs/axios/axios';
import { AxiosError, AxiosResponse } from 'axios';
import { notifyError, notifySuccess } from '../../components/toast/notifications';
import { useRouter } from 'next/navigation';

interface CreateUserRequest {
  username: string;
  preferred_language: string;
  photo_url: string;
  oauth_id: number;
  oauth_provider: 'GitHub';
}

interface SignUpResponse {
  message: string;
  user?: AuthResponse;
}

export default function SignUpPage() {
  const { data: session, status } = useSession();
  const router = useRouter();
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const languages = LANGUAGES;

  const onSubmit = handleSubmit(async (data: CreateUserRequest) => {
    data.oauth_id = Number(session.user.id);
    data.oauth_provider = 'GitHub';
    data.photo_url = session.user.image;
    try {
      const response: AxiosResponse<SignUpResponse> = await POST(`/auth/signup`, data);
      const { message } = response.data;
      notifySuccess(message);
      reset();
      onClose();
      const signOutData = await signOut({ redirect: false, callbackUrl: '/login' });
      router.push(signOutData.url);
    } catch (error) {
      const message = error.message.data ? error.message.data.message : 'Server Error';
      notifyError(message);
      onClose();
      await signOut({ redirect: false });
    }
  });

  useEffect(() => {
    status === 'authenticated' && onOpen();
  }, [status]);

  return (
    <>
      <Modal
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        placement="top-center"
        isDismissable={false}
      >
        <ModalContent>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">Sign Up Form</ModalHeader>
              <ModalBody>
                <Input
                  {...register('username', {
                    required: 'Username is Required',
                    validate: {
                      notEmpty: value =>
                        value.trim() !== '' ||
                        'Username cannot be empty or contain only whitespace',
                    },
                  })}
                  autoFocus
                  isRequired
                  label="Username"
                  placeholder="Enter your username"
                  variant="bordered"
                  defaultValue={session?.user.name}
                  errorMessage={errors.username?.message as string}
                />
                <Select
                  {...register('preferred_language')}
                  label="Preferred Language"
                  disallowEmptySelection={true}
                  isRequired={true}
                >
                  {languages.map(language => (
                    <SelectItem key={language} value={language}>
                      {language}
                    </SelectItem>
                  ))}
                </Select>
              </ModalBody>
              <ModalFooter>
                <Button
                  color="danger"
                  variant="flat"
                  onPress={() => {
                    signOut({ redirect: false });
                    onClose();
                  }}
                >
                  Close
                </Button>
                <Button color="primary" onClick={onSubmit}>
                  Submit
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
      <AuthCard authTitle={'Sign Up'} />
    </>
  );
}
