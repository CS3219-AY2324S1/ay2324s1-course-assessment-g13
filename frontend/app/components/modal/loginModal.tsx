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

const LoginModal = () => {
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = handleSubmit(async data => {
    // eslint-disable-next-line no-console
    console.log(data); //add in your api call to check login
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
                {...register('email', {
                  required: 'Email is required',
                })}
                label="Email"
                isRequired
                variant="bordered"
                placeholder="Enter your email"
                labelPlacement="outside"
                errorMessage={errors.email?.message as string}
              />
              <Input
                {...register('password', {
                  required: 'Password is required',
                })}
                label="Password"
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
