'use client';
import { useState, useEffect } from 'react';
import { Chip } from '@nextui-org/chip';
import { Category, ComplexityToColor, Question } from '../types/question';
import { GET } from "../libs/axios/axios";
import { notifyError } from '../components/toast/notifications';
import Editor from '@monaco-editor/react';
import { Button } from '@nextui-org/react';

export default function Collab() {
  const [question, setQuestion] = useState({});

  const fetchQuestion = async (complexity : string) => {
    try {
      const response = await GET(`questions/${complexity}`);
      setQuestion(response.data);
    } catch (error) {
      notifyError(error.message.data);
    }
  };

  useEffect(() => {
    fetchQuestion('Medium');
  }, []);

  const editorOptions = {
    minimap: {
      enabled: false
    }
  };

  return (
    <div className="flex">
      <div className="w-1/2 h-1/2 m-8">
        <div className="p-3 flex flex-col justify-center text-md" style={{backgroundColor: '#1e1e1e'}}>
          <div className="my-12 flex align-center">
            <p className="mr-12">{question.title}</p>
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
            <p className="my-2"key={index}>{line}</p>
          ))}
        </div>
        <div className="flex justify-end">
          <Button
            color="danger" 
            variant="solid" 
            className="text-lg mt-4" 
          >
            End Collaboration
          </Button>
        </div>
      </div>
      <div className="w-1/2 m-6">
        <div className="flex align-center justify-between font-bold">
          <h2 className='mb-2'>Editor</h2>
          <p>Current Language: Python</p>
        </div>
        <Editor height="85vh" theme="vs-dark" defaultLanguage="python"
         defaultValue="# Type answer here" options={editorOptions}/>
      </div>
    </div>
  )
}
