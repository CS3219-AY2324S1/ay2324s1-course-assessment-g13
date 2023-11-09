import {
  Button,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  useDisclosure,
} from '@nextui-org/react';
import { DELETE } from '../../libs/axios/axios';
import { notifyError, notifySuccess } from '../toast/notifications';
import { useDispatch } from 'react-redux';
import { logout as UserLogout} from '../../libs/redux/slices/userSlice';
import { logout as AuthLogout} from '../../libs/redux/slices/authSlice';
import { useRouter } from 'next/navigation';
import { signOut } from 'next-auth/react';

const DeleteAccountModal = () => {
  const dispatch = useDispatch();
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
  const router = useRouter();

  const handleDelete = async () => {
    try {
      const response = await DELETE(`/auth/user`);
      dispatch(UserLogout());
      dispatch(AuthLogout());
      router.push('/');
      notifySuccess(response.data);
      onClose();
      signOut();
    } catch (error) {
      notifyError(error.message.data.message);
    }
  };

  return (
    <>
      <Button className="bg-red-500 " onPress={onOpen}>
        Delete Account
      </Button>
      <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
        <ModalContent>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">
                {' '}
                Are you sure you want to delete your account?
              </ModalHeader>
              <ModalBody>
                <p>This action is irreversible and all your data will be deleted.</p>
              </ModalBody>
              <ModalFooter>
                <Button color="default" onPress={onClose}>
                  Close
                </Button>
                <Button color="danger" onPress={handleDelete}>
                  Delete
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
};

export default DeleteAccountModal;
