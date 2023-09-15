'use client';
import QuestionsTable from './questionsTable';
import QuestionAddModal from './addQuestionModal';
import { useState } from 'react';
import axiosInstance from '../axios/axios';
import { notifyError } from '../components/notifications';


export default function Questions() {
  const [questions, setQuestions] = useState([]);
  const fetchData = async () => {
    try {
      const response = await axiosInstance.get('');
      setQuestions(response.data == null ? [] : response.data);
    } catch (error) {
      if (error.response) {
        notifyError(error.response.data.error);
      } else {
        notifyError(error.message);
      }
    } 
  }

  return (
    <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="questions-header flex justify-between items-center mb-5">
        <span className="text-3xl">Question Bank</span>
        <QuestionAddModal fetchData={fetchData} />
      </div>
      <div className="table w-full">
        <QuestionsTable fetchData={fetchData} questions={questions} />
      </div>
    </div>
  );
}
