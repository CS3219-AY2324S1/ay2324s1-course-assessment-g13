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

interface SignUpProps {
  isNav?: boolean;
}

const SignupModal: React.FC<SignUpProps> = ({ isNav = false }) => {
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = handleSubmit(async data => {
    // eslint-disable-next-line no-console
    POST('/users', data);
    // console.log(data); //add in your axios call to add user at backend
    reset();
    onClose();
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
              {/* <Input
                {...register('email', {
                  required: 'Email is required',
                })}
                label="Email"
                isRequired
                variant="bordered"
                placeholder="Enter your email"
                labelPlacement="outside"
                errorMessage={errors.email?.message as string}
              /> */}
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
              Sign Up
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default SignupModal;
