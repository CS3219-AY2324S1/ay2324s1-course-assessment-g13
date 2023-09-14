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
import { ToastContainer, toast } from 'react-toastify';
import axiosInstance from '../requests';
import 'react-toastify/dist/ReactToastify.css';


const DeleteConfirmationModal = ({ title, id, setUpdate}) => {
  const { isOpen, onOpen, onClose, onOpenChange } = useDisclosure();
  const notifyDelete = () => toast.error("Question Deleted Successfully", {
    theme:"dark"
  });
  const notifyError = (err : string) => toast.warn(err, {
    theme: "dark"
  });

  const handleDelete = async () => {
    try {
      axiosInstance.delete(`/${id}`);
      setUpdate(true);
      notifyDelete();
    } catch(error) {
      const status = error.response.status;
      if (status === 404) {
        notifyError("Not Found: Specified question has already been deleted");
      }
    } finally {
      onClose();
    }
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
      <ToastContainer/>
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
