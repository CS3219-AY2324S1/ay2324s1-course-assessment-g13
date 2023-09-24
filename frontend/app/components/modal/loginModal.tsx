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
import { POST } from '../../libs/axios/axios';
import { login } from '../../libs/redux/slices/userSlice';

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
    // eslint-disable-next-line no-console
    const response = await POST('/auth/login', data);
    if (response.status != 200) {
      return;
    }

    dispatch(login(response.data.username));
    router.push('/questions');

    reset();
    onClose();
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
                type='password'
                isRequired
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
