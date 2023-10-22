'use client';
import { useState, useEffect } from 'react';
import { Chip } from '@nextui-org/chip';
import { Category, Complexity, ComplexityToColor, Question } from '../types/question';
import { GET } from "../libs/axios/axios";
import { notifyError } from '../components/toast/notifications';
import Editor from '@monaco-editor/react';
import { useSelector } from 'react-redux';
import { selectCollabChatState, selectCollabLeavingState, setIsLeaving } from '../libs/redux/slices/collabSlice';
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
import { useRef } from 'react';
import { selectUsername } from '../libs/redux/slices/userSlice';

export default function Collab() {
  const collabLeavingState = useSelector(selectCollabLeavingState);
  const chatOpenState = useSelector(selectCollabChatState);
  const userId = useSelector(selectUsername);
  const dispatch = useDispatch();
  const router = useRouter();
  const searchParams = useSearchParams();
  const roomId = searchParams.get('room_id');
  const defaultCode = "# Type answer here";
  const [code, setCode] = useState(defaultCode);
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');
  const ws = useRef(null);
  useEffect(() => {
    ws.current = new WebSocket(`ws://localhost:5005/ws/${roomId}`);
    fetchQuestion();
    // onmessage is for receiving messages
    ws.current.onmessage = function (event) {
      var message = JSON.parse(event.data);
      if (message["Type"] === "code") {
        setCode(message["Content"]);
      } else {
        console.log(message);
        setMessages((prevMessages) => [...prevMessages, message]);
      }
    }
  }, [])

  const [question, setQuestion] = useState<Question>({
    id: "",
    title: "",
    categories: [],
    complexity: Complexity.EASY,
    description: ""
  });

  const fetchQuestion = async () => {
    try {
      const idResponse = await GET(`ws/question/${roomId}`);
      const response = await GET(`questions/${idResponse.data}`);
      setQuestion(response.data as Question);
    } catch (error) {
      notifyError(error.message.data);
    }
  };

  const handleEditorChange = (value: string, event) => {
    const message = {
      content: value,
      type: "code",
    };
    ws.current.send(JSON.stringify(message));
  }

  const handleSendMessage = () => {
    const sendMessage = {
      content: newMessage,
      type: "chat",
    };
    ws.current.send(JSON.stringify(sendMessage));

    const message = {
      UserId: userId,
      Content: newMessage,
    }
    setNewMessage('');
    setMessages((prevMessages) => [...prevMessages, message]);
  };

  const editorOptions = {
    minimap: {
      enabled: false
    }
  };

  const exitRoom = () => {
    ws.current.close();
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
          <Editor
            height="80vh"
            theme="vs-dark"
            defaultLanguage="python"
            value={code}
            onChange={handleEditorChange}
            options={editorOptions}
          />
        </div>
      </div>
      <Modal size={'xl'} isOpen={collabLeavingState} onClose={() => dispatch(setIsLeaving(false))} placement="top-center">
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
      {chatOpenState && (
        <div className="bg-black w-1/4 h-2/5 absolute left-0 bottom-0 border border-white border-opacity-20">
        <div className="h-5/6 p-4 overflow-y-auto">
          {messages.map((message, index) => (
            <div key={index} className="mb-4 flex">
              <span>{message.UserId}</span>
              <span>{message.Content}</span>
            </div>
          ))}
        </div>
        <div className="mb-4 border-b border-gray-400"></div>
        <div className="flex justify-center p-4">
          <input
            type="text"
            placeholder="Type your message..."
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            className="w-3/4 p-2 rounded"
          />
          <Button onClick={handleSendMessage} color="primary" className='ml-5'>
            Send
          </Button>
        </div>
      </div>
      )}
    </>
  )
}
