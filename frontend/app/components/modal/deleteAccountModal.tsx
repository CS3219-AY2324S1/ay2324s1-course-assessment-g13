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
import useAuth from '../../hooks/useAuth';
import { notifyError, notifySuccess } from '../Notifications';
import { useDispatch } from 'react-redux';
import { logout } from '../../libs/redux/slices/userSlice';
import { useRouter } from 'next/navigation';

const DeleteAccountModal = () => {
  const { userId:id } = useAuth();
  const dispatch = useDispatch();
  const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
  const router = useRouter();

  const handleDelete = async () => {
    try {
      const response = await DELETE(`/users/delete/${id}`);
      dispatch(logout());
      router.push('/');
      notifySuccess(response.data);
      onClose();
    } catch (error) {
      notifyError(error.message.data);
    }
  };

  return (
    <>
      <Button className="bg-red-500 rounded-md " onPress={onOpen}>
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
