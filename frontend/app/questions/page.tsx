'use client';
import QuestionsTable from './questionsTable';
import QuestionAddModal from './addQuestionModal';
import { useState } from 'react';

export default function Questions() {
  const [update, setUpdate] = useState(true);
  return (
    <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="questions-header flex justify-between items-center mb-5">
        <span className="text-3xl">Question Bank</span>
        <QuestionAddModal setUpdate={setUpdate} />
      </div>
      <div className="table w-full">
        <QuestionsTable update={update} setUpdate={setUpdate} />
      </div>
    </div>
  );
}
