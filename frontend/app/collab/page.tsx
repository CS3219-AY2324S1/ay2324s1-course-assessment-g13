'use client';
import { useState, useEffect } from 'react';
import { Chip } from '@nextui-org/chip';
import { Category, Complexity, ComplexityToColor, Question } from '../types/question';
import { GET } from "../libs/axios/axios";
import { notifyError } from '../components/toast/notifications';
import Editor from '@monaco-editor/react';
import { useSelector } from 'react-redux';
import { selectCollabState, setIsLeaving } from '../libs/redux/slices/collabSlice';
import {
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter
} from '@nextui-org/modal';
import { Button } from '@nextui-org/react';
import { useDispatch } from 'react-redux';
import { useRouter, useSearchParams } from 'next/navigation';

export default function Collab() {
  const collabState = useSelector(selectCollabState);
  const dispatch = useDispatch();
  const router = useRouter();
  const searchParams = useSearchParams();
  const roomId = searchParams.get('room_id');
  const [question, setQuestion] = useState<Question>({
    id: "",
    title: "",
    categories: [],
    complexity: Complexity.EASY,
    description: ""
  });

  const fetchQuestion = async (complexity : string) => {
    try {
      const id_response = await GET(`questions/complexity/${complexity}`);
      const response = await GET(`questions/${id_response.data}`);
      setQuestion(response.data as Question);
    } catch (error) {
      notifyError(error.message.data);
    }
  };

  useEffect(() => {
    fetchQuestion('Easy');
  }, []);

  const editorOptions = {
    minimap: {
      enabled: false
    }
  };

  const exitRoom = () => {
    dispatch(setIsLeaving(false));
    router.push('/');
  }

  return (
    <>
      <div className="flex">
        <div className="w-1/2 m-8" style={{backgroundColor: '#1e1e1e'}}>
          <div className="p-3 flex flex-col justify-center">
            <div className="my-12 flex align-center">
              <p className="mr-12 text-lg">{question.title}</p>
              <Chip color={ComplexityToColor[question.complexity]} className="mx-2">
                {question.complexity}
              </Chip>
              {question.categories && (question.categories as Category[]).map(category => (
              <Chip variant="bordered" key={category} className="mx-2">
                {category}
              </Chip>
              ))}
            </div>
            <div className="mb-4 border-b border-gray-400"></div>
            {question.description && question.description.split('\n').map((line : string, index : number) => (
              <p className="my-2 test-md" key={index}>{line}</p>
            ))}
          </div>
        </div>
        <div className="w-1/2 m-6 flex flex-col">
          <div className="flex align-center justify-between font-bold">
            <h2 className='mb-2'>Editor</h2>
            <p>Current Language: Python</p>
          </div>
          <Editor height="80vh" theme="vs-dark" defaultLanguage="python"
          defaultValue="# Type answer here" options={editorOptions}/>
        </div>
      </div>
      <Modal size={'xl'} isOpen={collabState} onClose={() => dispatch(setIsLeaving(false))} placement="top-center">
        <ModalContent>
          {onClose => (
            <>
              <ModalHeader className="flex flex-col gap-1">Exit Collaboration Room</ModalHeader>
              <ModalBody>
                {"Are you sure you want to leave the collaboration room?"}
              </ModalBody>
              <ModalFooter>
                <Button color="danger" onClick={() => {
                   onClose();
                   exitRoom();
                }}>
                  Confirm
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  )
}
