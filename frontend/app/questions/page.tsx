'use client';
import QuestionsTable from './questionsTable';
import QuestionAddModal from './modal/addQuestionModal';
import { useEffect, useState } from 'react';
import { getData } from '../libs/axios/axios';
import { notifyError } from '../components/toast/notifications';
import useAuth from '../(auth)/hooks/useAuth';
import LoadingTable from './Loading/LoadingTable';


export default function Questions() {
  const [questions, setQuestions] = useState([]);
  const [isAdmin, setIsAdmin] = useState(false);
  const { userRole } = useAuth();

  useEffect(() => {
    setIsAdmin(userRole === "admin");
  }, [])

  const fetchQuestions = () => {
    setTimeout(()=>{}, 10000)
    getData('questions').then(res => {
      if (res.data) {
        setQuestions(res.data);
      } else {
        notifyError(res.error);
      }
    });
  }

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
