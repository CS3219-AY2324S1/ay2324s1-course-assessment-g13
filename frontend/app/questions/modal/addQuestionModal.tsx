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
import { Category, Complexity, Question } from '../../types/question';
import { useForm } from 'react-hook-form';
import { notifySuccess, notifyError } from '../../components/toast/notifications';
import { POST } from '../../libs/axios/axios';


export default function QuestionAddModal({fetchQuestions}) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const categories = Object.values(Category);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  const onSubmit = handleSubmit(async (data: Question) => {
    const modifiedData = {
      ...data,
      categories: (data.categories as string).split(',') as Category[],
    };

    try {
      const response = await POST('questions', modifiedData);
      fetchQuestions();
      notifySuccess(response.data);
      onOpenChange();
      reset();
    } catch (error) {
      notifyError(error.message.data);
    }
  });

  return (
    <>
      <Button color="primary" variant="ghost" className="text-lg py-5" onPress={onOpen}>
        Add Question
      </Button>
      <Modal
        size={'2xl'}
        isOpen={isOpen}
        onOpenChange={() => {
          onOpenChange();
          reset();
        }}
        placement="top-center"
      >
        <ModalContent>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">Add Question</ModalHeader>
              <form>
                <ModalBody>
                  <Input
                    {...register('title', 
                      { required: 'Title is required',
                        validate: { 
                          noNumbers : (value) => !/\d/.test(value) 
                          || 'Title should not contain numbers',
                        },
                      }
                    )}
                    autoFocus
                    label="Title"
                    placeholder="Enter Question Title"
                    isRequired
                    variant="bordered"
                    labelPlacement="outside"
                    errorMessage={errors.title?.message as string}
                  />
                  <Select
                    {...register('complexity', { required: 'Complexity is required' })}
                    label="Complexity"
                    isRequired
                    placeholder="Select Complexity"
                    variant="bordered"
                    labelPlacement="outside"
                    errorMessage={errors.complexity?.message as string}
                  >
                    {Object.values(Complexity).map(c => (
                      <SelectItem key={c} value={c}>
                        {c.toUpperCase()}
                      </SelectItem>
                    ))}
                  </Select>
                  <Select
                    {...register('categories', { required: 'Category is required' })}
                    items={categories}
                    label="Category"
                    isRequired
                    variant="bordered"
                    labelPlacement="outside"
                    isMultiline
                    selectionMode="multiple"
                    placeholder="Select Categories"
                    errorMessage={errors.categories?.message as string}
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
                    {...register('description', 
                      { 
                        required: 'Description is required',
                        validate: {
                          notEmpty: (value) => value.trim() !== '' 
                          || 'Description cannot be empty or contain only whitespace',
                        }, 
                      }
                    )}
                    label="Description"
                    isRequired
                    labelPlacement="outside"
                    placeholder="Enter Question Description"
                    errorMessage={errors.description?.message as string}
                  />
                </ModalBody>
              </form>
              <ModalFooter>
                <Button
                  color="danger"
                  variant="flat"
                  onClick={() => {
                    onClose();
                    reset();
                  }}
                >
                  Close
                </Button>
                <Button color="primary" onClick={onSubmit}>
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
