'use client';
import QuestionsTable from './questionsTable';
import QuestionAddModal from './addQuestionModal';
import { useState } from 'react';
import { getData } from '../libs/axios/axios';
import { notifyError } from '../components/toast/notifications';


export default function Questions() {
  const [questions, setQuestions] = useState([]);
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
    <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="questions-header flex justify-between items-center mb-5">
        <span className="text-3xl">Question Bank</span>
        <QuestionAddModal fetchQuestions={fetchQuestions} />
      </div>
      <div className="table w-full">
        <QuestionsTable fetchQuestions={fetchQuestions} questions={questions} />
      </div>
    </div>
  );
}
