'use client';
import QuestionsTable from './questionsTable';
import QuestionAddModal from './addQuestionModal';
import { useEffect, useState } from 'react';
import { getData } from '../libs/axios/axios';
import { notifyError } from '../components/toast/notifications';
import ProtectedRoute from '../libs/_protected/ProtectedRoute';
import { useRouter } from 'next/navigation';
import useAuth from '../(auth)/hooks/useAuth';


export default function Questions() {
  const [questions, setQuestions] = useState([]);
  const [isAdmin, setIsAdmin] = useState(false);
  const { userRole } = useAuth();
  const router = useRouter();

  useEffect(() => {
    setIsAdmin(userRole === "admin");
  }, [isAdmin])

  const fetchQuestions = () => {
    getData('questions').then(res => {
      if (res.data) {
        setQuestions(res.data);
      } else {
        notifyError(res.error);
      }
    });
  }

  return (
    <ProtectedRoute router={router}>
      <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
        <div className="questions-header flex justify-between items-center mb-5">
          <span className="text-3xl">Question Bank</span>
          {isAdmin && <QuestionAddModal fetchQuestions={fetchQuestions} />}
        </div>
        <div className="table w-full">
          <QuestionsTable isAdmin={isAdmin} fetchQuestions={fetchQuestions} questions={questions} />
        </div>
      </div>
    </ProtectedRoute>
  );
}
