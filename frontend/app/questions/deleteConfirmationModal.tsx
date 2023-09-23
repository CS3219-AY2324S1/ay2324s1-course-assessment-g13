import {
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Button,
  useDisclosure,
  Tooltip,
} from '@nextui-org/react';
import { DeleteIcon } from './assets/DeleteIcon';
import { useDispatch } from 'react-redux';
import { deleteQuestion } from '../redux/slices/questionBankSlice';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const DeleteConfirmationModal = ({ title }: { title: string }) => {
  const dispatch = useDispatch();
  const { isOpen, onOpen, onClose, onOpenChange } = useDisclosure();
  const notifyDelete = () =>
    toast.success('Question Deleted Successfully', {
      theme: 'dark',
    });

  const handleDelete = () => {
    dispatch(deleteQuestion(title));
    notifyDelete();
    onClose();
  };
  return (
    <>
      <Tooltip content="Delete question">
        <Button onPress={onOpen} className="h-fit min-w-0 px-0 bg-transparent">
          <span className="text-lg text-danger cursor-pointer active:opacity-50">
            <DeleteIcon />
          </span>
        </Button>
      </Tooltip>
      <ToastContainer />
      <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
        <ModalContent>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">{title}</ModalHeader>
              <ModalBody>{'Are you sure you want to delete this question?'}</ModalBody>
              <ModalFooter>
                <Button color="primary" variant="light" onPress={onClose}>
                  Close
                </Button>
                <Button color="danger" variant="light" onPress={handleDelete}>
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

export default DeleteConfirmationModal;
