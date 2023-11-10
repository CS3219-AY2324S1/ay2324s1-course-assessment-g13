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
import { EyeIcon } from '../../public/EyeIcon';

const SolutionModal = ({ question, solution }) => {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  const content =
    solution === '' || solution === '# Type answer here' ? 'No solution provided.' : solution;
  return (
    <>
      <Tooltip content="Solution">
        <Button onPress={onOpen} className="h-fit min-w-0 px-0 bg-transparent">
          <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
            <EyeIcon />
          </span>
        </Button>
      </Tooltip>
      <Modal
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        size="3xl"
        scrollBehavior="inside"
        className="h-4/5"
      >
        <ModalContent className="fit-content">
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">{question}</ModalHeader>
              <ModalBody className="whitespace-pre-line">{content}</ModalBody>
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

export default SolutionModal;
