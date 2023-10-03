'use client';
import QuestionsTable from './questionsTable';
import QuestionAddModal from './modal/addQuestionModal';
import { useState } from 'react';
import { GET } from '../libs/axios/axios';
import { notifyError } from '../components/toast/notifications';
import useAuth from '../hooks/useAuth';

export default function Questions() {
  const [questions, setQuestions] = useState([]);
  const { userRole } = useAuth();
  const isAdmin = userRole === "admin";
  
  const fetchQuestions = async () => {
    try {
      const response = await GET('questions');
      setQuestions(response.data == null ? [] : response.data);
    } catch (error) {
      notifyError(error.message.data);
    }
  };

  return (
    <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="questions-header flex justify-between items-center mb-5">
        <span className="text-3xl">Question Bank</span>
        {isAdmin && <QuestionAddModal fetchQuestions={fetchQuestions} />}
      </div>
      <div className="table w-full">
        <QuestionsTable isAdmin={isAdmin} fetchQuestions={fetchQuestions} questions={questions} />
      </div>
    </div>
  );
}
