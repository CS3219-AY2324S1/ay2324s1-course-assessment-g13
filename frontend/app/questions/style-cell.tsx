import { Chip } from '@nextui-org/chip';
import QuestionDescriptionModal from './modal/descriptionModal';
import { Category, ComplexityToColor, Question } from '../types/question';
import { Key } from 'react';
import DeleteConfirmationModal from './modal/deleteConfirmationModal';

interface StyleCellProps {
  isAdmin: boolean,
  item: Question & { listId: number };
  columnKey: Key;
  fetchQuestions: () => void;
}

const StyleCell: React.FC<StyleCellProps> = ({ isAdmin, item, columnKey, fetchQuestions }) => {
  switch (columnKey) {
    case 'id':
      return <span>{item.listId}</span>;
    case 'title':
      return <span>{item.title}</span>;
    case 'category':
      return (
        <div className="relative flex items-center">
          {(item.categories as Category[]).map(category => (
            <Chip variant="bordered" key={category}>
              {category}
            </Chip>
          ))}
        </div>
      );
    case 'complexity':
      return <Chip color={ComplexityToColor[item.complexity]}>{item.complexity}</Chip>;
    case 'actions':
      return (
        <div className="relative flex items-center gap-5">
          <QuestionDescriptionModal title={item.title} description={item.description} />
          {isAdmin && 
            <DeleteConfirmationModal 
              title={item.title} 
              id={item.id} 
              fetchQuestions={fetchQuestions}
            />}
        </div>
      );

    default:
      return '';
  }
};

export default StyleCell;
