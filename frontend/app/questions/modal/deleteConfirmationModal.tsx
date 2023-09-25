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
import { DeleteIcon } from '../assets/DeleteIcon';
import { notifyWarning, notifyError } from '../../components/toast/notifications';
import { deleteEntry } from '../../libs/axios/axios';


const DeleteConfirmationModal = ({ title, id, fetchQuestions }) => {
  const { isOpen, onOpen, onClose, onOpenChange } = useDisclosure();

  const handleDelete = () => {
    deleteEntry(`questions/${id}`)
    .then(res => {
      if (res.message) {
        fetchQuestions();
        notifyWarning(res.message);
      } else {
        notifyError(res.error);
      }
    })
    .finally(() => onClose());
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
