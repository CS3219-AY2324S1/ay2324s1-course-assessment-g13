import QuestionsTable from './questionsTable';
import QuestionAddModal from './addQuestionModal';

export default function Questions() {
  return (
    <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="questions-header flex justify-between items-center mb-5">
        <span className="text-3xl">Question Bank</span>
        <QuestionAddModal />
      </div>
      <div className="table w-full">
        <QuestionsTable />
      </div>
    </div>
  );
}
