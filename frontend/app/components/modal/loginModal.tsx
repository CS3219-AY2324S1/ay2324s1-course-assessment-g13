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
import { useForm } from 'react-hook-form';
import { useDispatch } from 'react-redux';
import { useRouter } from 'next/navigation';
import { POST } from '../../axios/axios';
import { login } from '../../redux/slices/userSlice';
import { notifyError, notifySuccess } from '../Notifications';

const LoginModal = () => {
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
  const dispatch = useDispatch();
  const router = useRouter();

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = handleSubmit(async data => {
    try {
      const response = await POST('/login', data);
      dispatch(
        login({
          ...response.data,
        }),
      );
      router.push('/questions');
      notifySuccess(response.data.statusText);
      reset();
      onClose();
    } catch (error) {
      notifyError(error.message.data);
    }
  });

  return (
    <>
      <Button onPress={onOpen} variant="light" color="default">
        Login
      </Button>
      <Modal
        isOpen={isOpen}
        onOpenChange={() => {
          onOpenChange();
          reset();
        }}
      >
        <ModalContent>
          <ModalHeader className="flex flex-col gap-1 text-center">Login</ModalHeader>
          <form className="flex flex-col gap-8">
            <ModalBody>
              <Input
                {...register('username', {
                  required: 'Username is required',
                })}
                label="Username"
                isRequired
                variant="bordered"
                placeholder="Enter your username"
                labelPlacement="outside"
                errorMessage={errors.username?.message as string}
              />
              <Input
                {...register('password', {
                  required: 'Password is required',
                })}
                label="Password"
                isRequired
                type="password"
                variant="bordered"
                placeholder="Enter your password"
                labelPlacement="outside"
                errorMessage={errors.password?.message as string}
              />
            </ModalBody>
          </form>
          <ModalFooter>
            <Button color="primary" onClick={onSubmit}>
              Login
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default LoginModal;
