'use client';

import {
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  useDisclosure,
} from '@nextui-org/modal';
import { Button } from '@nextui-org/button';
import { Input } from '@nextui-org/input';
import { Select, SelectItem } from '@nextui-org/select';
import { Chip } from '@nextui-org/chip';
import { Textarea } from '@nextui-org/react';
import { Complexity, categories } from './data';

export default function QuestionAddModal() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <Button color="primary" variant="ghost" className="text-lg py-5" onPress={onOpen}>
        Add Question
      </Button>
      <Modal size={'2xl'} isOpen={isOpen} onOpenChange={onOpenChange} placement="top-center">
        <ModalContent>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">Add Question</ModalHeader>
              <ModalBody>
                <Input
                  autoFocus
                  label="Title"
                  placeholder="Enter Question Title"
                  variant="bordered"
                  labelPlacement="outside"
                  className="mb-6"
                />
                <Select
                  label="Complexity"
                  isRequired
                  placeholder="Select Complexity"
                  variant="bordered"
                  labelPlacement="outside"
                >
                  {Object.values(Complexity).map(c => (
                    <SelectItem key={c} value={c}>
                      {c.toUpperCase()}
                    </SelectItem>
                  ))}
                </Select>
                <Select
                  items={categories}
                  label="Category"
                  isRequired
                  variant="bordered"
                  labelPlacement="outside"
                  isMultiline
                  selectionMode="multiple"
                  placeholder="Select Categories"
                  renderValue={items => {
                    return (
                      <div className="flex flex-wrap gap-2">
                        {items.map(item => (
                          <Chip key={item.key} variant="bordered">
                            {item.key}
                          </Chip>
                        ))}
                      </div>
                    );
                  }}
                >
                  {categories.map(category => (
                    <SelectItem key={category} value={category}>
                      {category.toUpperCase()}
                    </SelectItem>
                  ))}
                </Select>
                <Textarea
                  label="Description"
                  labelPlacement="outside"
                  placeholder="Enter Question Description"
                />
              </ModalBody>
              <ModalFooter>
                <Button color="danger" variant="flat" onPress={onClose}>
                  Close
                </Button>
                <Button color="primary" onPress={onClose}>
                  Add
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}
