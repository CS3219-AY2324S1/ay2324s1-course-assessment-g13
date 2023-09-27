'use client';
import {
  Button,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  useDisclosure,
} from '@nextui-org/react';
import React from 'react';
import { useForm } from 'react-hook-form';
import { POST } from '../../axios/axios';
import { notifyError, notifySuccess } from '../Notifications';
import { useRouter } from 'next/navigation';
import { useDispatch } from 'react-redux';
import { login } from '../../redux/slices/userSlice';

interface SignUpData {
  username: string;
  password: string;
  confirmPassword: string;
}
interface SignUpProps {
  isNav?: boolean;
}

const SignupModal: React.FC<SignUpProps> = ({ isNav = false }) => {
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
  const router = useRouter();
  const dispatch = useDispatch();

  const {
    register,
    handleSubmit,
    reset,
    getValues,
    formState: { errors },
  } = useForm({
    defaultValues: {
      username: '',
      password: '',
      confirmPassword: '',
    },
  });

  const onSubmit = handleSubmit(async (data: SignUpData) => {
    try {
      const response = await POST('/users', {
        username: data.username,
        password: data.password,
      });
      router.push('/questions');
      dispatch(
        login({
          ...response.data,
        }),
      );
      notifySuccess(response.data);
      reset();
      onClose();
    } catch (error) {
      notifyError(error.message.data);
    }
  });
  return (
    <>
      <div className="w-full md:w-4/5 lg:w-1/2 px-4 max-w-lg">
        <Button
          onPress={onOpen}
          color="primary"
          variant={isNav ? 'flat' : 'solid'}
          fullWidth={true}
        >
          Sign Up
        </Button>
      </div>
      <Modal
        isOpen={isOpen}
        onOpenChange={() => {
          onOpenChange();
          reset();
        }}
      >
        <ModalContent>
          <ModalHeader className="flex flex-col gap-1 text-center">Sign Up Now</ModalHeader>
          <form className="flex flex-col gap-8">
            <ModalBody>
              <Input
                {...register('username', {
                  required: 'Username is required',
                })}
                label="Username"
                autoFocus
                isRequired
                variant="bordered"
                placeholder="Enter your username"
                labelPlacement="outside"
                errorMessage={errors.username?.message as string}
              />
              <Input
                {...register('password', {
                  required: 'Password is required',
                  minLength: {
                    value: 8,
                    message: 'Password must be at least 8 characters long',
                  },
                })}
                label="Password"
                isRequired
                type="password"
                variant="bordered"
                placeholder="Enter your password"
                labelPlacement="outside"
                errorMessage={errors.password?.message as string}
              />
              <Input
                {...register('confirmPassword', {
                  required: 'Password is required',
                  validate: value => value === getValues().password || 'The passwords do not match',
                })}
                label="Confirm Password"
                isRequired
                type="password"
                variant="bordered"
                placeholder="Comfirm your password"
                labelPlacement="outside"
                errorMessage={errors.confirmPassword?.message as string}
              />
            </ModalBody>
          </form>
          <ModalFooter>
            <Button color="primary" onClick={onSubmit}>
              Sign Up
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default SignupModal;
