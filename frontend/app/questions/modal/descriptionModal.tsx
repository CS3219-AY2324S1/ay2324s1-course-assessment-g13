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
import { EyeIcon } from '../../../public/EyeIcon';
import Markdown from 'react-markdown';

const QuestionDescriptionModal = ({ title, description }) => {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <Tooltip content="Question Description">
        <Button onPress={onOpen} className="h-fit min-w-0 px-0 bg-transparent">
          <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
            <EyeIcon />
          </span>
        </Button>
      </Tooltip>
      <Modal isOpen={isOpen} onOpenChange={onOpenChange} size="3xl" scrollBehavior='inside' className='h-4/5'>
        <ModalContent className='fit-content'>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">{title}</ModalHeader>
              <ModalBody className="whitespace-pre-line">
                <Markdown>
                  {description}
                </Markdown>
              </ModalBody>
              <ModalFooter>
                <Button color="danger" variant="light" onPress={onClose}>
                  Close
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
};

export default QuestionDescriptionModal;
