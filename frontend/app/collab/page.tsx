'use client';
import { useState, useEffect, ChangeEvent } from 'react';
import { Complexity, ComplexityToColor, Question } from '../types/question';
import { GET } from "../libs/axios/axios";
import { notifyError, notifySuccess, notifyWarning } from '../components/toast/notifications';
import Editor from '@monaco-editor/react';
import { useSelector } from 'react-redux';
import { selectCollabState, setIsChatOpen, setIsLeaving } from '../libs/redux/slices/collabSlice';
import {
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter
} from '@nextui-org/modal';
import { Button, Input, Chip, Select, SelectItem } from '@nextui-org/react';
import { useDispatch } from 'react-redux';
import { useRouter, useSearchParams } from 'next/navigation';
import { useRef } from 'react';
import { selectUsername } from '../libs/redux/slices/userSlice';
import { LANGUAGES, Language } from '../constants/languages';
import Markdown from 'react-markdown';

export default function Collab() {
  const collabState = useSelector(selectCollabState);
  const username = useSelector(selectUsername);
  const dispatch = useDispatch();
  const router = useRouter();
  const searchParams = useSearchParams();
  const roomId = searchParams.get('room_id');
  const defaultCode = "# Type answer here";
  const [code, setCode] = useState(defaultCode);
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');
  const ws = useRef(null);
  const languages = LANGUAGES.slice(1);
  const [currentLanguage, setCurrentLanguage] = useState(LANGUAGES[10]);
  const [isPartnerPresent, setIsPartnerPresent] = useState(true);
  const [question, setQuestion] = useState<Question>({
    id: "",
    title: "",
    categories: [],
    complexity: Complexity.EASY,
    description: ""
  });

  useEffect(() => {
    window.addEventListener('popstate', exitRoom);
    window.addEventListener('beforeunload', sendExitMessage);
    window.history.pushState(null, '', "");

    return () => {
      window.removeEventListener('popstate', exitRoom);
      window.removeEventListener('beforeunload', sendExitMessage);
    };
  }, []);

  useEffect(() => {
    ws.current = new WebSocket(`${process.env.NEXT_PUBLIC_COLLAB_SERVICE_URL}/ws/${roomId}/${username}`);
    // onmessage is for receiving messages
    ws.current.onmessage = function (event) {
      var message = JSON.parse(event.data);
      if (message.Type === "code") {
        setCode(message.Content);
      } else if (message.Type === "chat") {
        setMessages((prevMessages) => [...prevMessages, {
          content: message.Content,
          user: "Other",
        }]);
        notifyWarning("You have unread messages!");
      } else if (message.Type === "language") {
        setCurrentLanguage(message.Content as Language);
        notifyWarning(`Editor's language has been changed to ${message.Content}`);
      } else if (message.Type === "exit") {
        setIsPartnerPresent(false);
        notifyError(message.Content);
      } else {
        setIsPartnerPresent(true);
        notifySuccess(message.Content);
      }
    }

    ws.current.onerror = function (event) {
      router.push('/');
    }
  }, []);

  useEffect(() => {
    fetchQuestion();
    dispatch(setIsLeaving(false));
    dispatch(setIsChatOpen(false));
  }, []);

  useEffect(() => {
    if (isPartnerPresent && ws.current.readyState) {
      sendMessage(code, "code");
      sendMessage(currentLanguage, "language");
    } 
  }, [isPartnerPresent])

  const sendMessage = (value : string, type : string) => {
    const message = {
      content: value,
      roomId: roomId,
      username: username,
      type: type,
    };
    ws.current.send(JSON.stringify(message));
  }

  const fetchQuestion = async () => {
    try {
      const idResponse = await GET(`ws/${roomId}`);
      const response = await GET(`questions/${idResponse.data}`);
      setQuestion(response.data as Question);
    } catch (error) {
      notifyError(error.message.data);
    }
  };

  const handleEditorChange = (value: string, event) => {
    setCode(value);
    sendMessage(value, "code");
  }

  const sendChatMessage = () => {
    if (newMessage.trim() === "") {
      return;
    }

    sendMessage(newMessage, "chat");
    const message = {
      content: newMessage,
      user: "Current",
    }
    setNewMessage('');
    setMessages((prevMessages) => [...prevMessages, message]);
  };

  const editorOptions = {
    minimap: {
      enabled: false
    }
  };

  const sendExitMessage = () => {
    sendMessage(`${username} has left the room!`, "exit");
    ws.current.close();
  }

  const exitRoom = () => {
    sendExitMessage();
    router.push('/');
  }

  const handleLanguageChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const newLanguage = e.target.value as Language;
    setCurrentLanguage(newLanguage);
    sendMessage(newLanguage, "language");
  }

  return (
    <>
      <div className="flex">
        <div className="w-1/2 m-8 overflow-x-auto" style={{backgroundColor: '#1e1e1e'}}>
          <div className="p-3 flex flex-col justify-center">
            <div className="my-12 flex align-center">
              <p className="mr-12 text-lg">{question.title}</p>
              <Chip color={ComplexityToColor[question.complexity]} className="mx-2">
                {question.complexity}
              </Chip>
              <div className="flex flex-wrap">
                {question.categories && question.categories.map(category => (
                <Chip variant="bordered" key={category} className="mx-2 mb-2">
                  {category}
                </Chip>))}
              </div>
            </div>
            <div className="mb-4 border-b border-gray-400"></div>
            {question.description &&
              <Markdown className="whitespace-pre-line">
                {question.description}
              </Markdown>
            }
          </div>
        </div>
        <div className="w-1/2 m-6 flex flex-col">
          <div className="flex align-center justify-between font-bold">
            <h2 className='mb-2'>Editor</h2>
            <Select
                className='w-1/5 mb-2 items-center'
                classNames={{"label": "text-md"}}
                size='sm'
                label="Language" 
                labelPlacement='outside-left'
                selectedKeys={[currentLanguage]}
                onChange={handleLanguageChange}
                disallowEmptySelection={true}
            >
              {languages.map((language) => (
              <SelectItem key={language} value={language}>
                  {language}
              </SelectItem>
              ))}
            </Select>
          </div>
          <Editor
            height="80vh"
            theme="vs-dark"
            defaultLanguage={currentLanguage.toLowerCase()}
            language={currentLanguage.toLowerCase()}
            value={code}
            onChange={handleEditorChange}
            options={editorOptions}
          />
        </div>
      </div>
      <Modal size="xl" isOpen={collabState.isLeaving} onClose={() => dispatch(setIsLeaving(false))} placement="top-center">
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
      <Modal 
        size="full" 
        isOpen={collabState.isChatOpen} 
        onClose={() => dispatch(setIsChatOpen(false))} 
        scrollBehavior="inside" 
        placement="center"
      >
        <ModalContent>
          <>
            <ModalHeader className="flex flex-col gap-1">
              Chat Room
              <p className="text-sm"> <span className="text-red-600">Warning:</span> The chat resets upon refresh or leaving the room!</p>
            </ModalHeader>
            <ModalBody>
              {messages.map((message, index) => (
                <p
                  key={index}
                  className={"my-5 border rounded-lg p-2 max-w-md break-words" 
                    + (message.user === "Other" ? " ml-10 mr-auto text-cyan-100 border-cyan-100" 
                    : " mr-10 ml-auto text-violet-100 border-violet-100")}
                >
                  {message.content}
                </p>
              ))}
            </ModalBody>
            <ModalFooter>
              <div className="flex w-full item-center mt-10 p-5">
                <Input
                  autoFocus
                  isRequired
                  type="text"
                  variant="bordered"
                  placeholder="Type your message..."
                  value={newMessage}
                  onValueChange={setNewMessage}
                />
                <Button onClick={sendChatMessage} color="primary" className="ml-5">
                  Send
                </Button>
              </div>
            </ModalFooter>
          </>
        </ModalContent>
      </Modal>
    </>
  );
}
