import { Chip } from '@nextui-org/chip';
import QuestionDescriptionModal from './descriptionModal';
import { Category, ComplexityToColor, Question } from '../types/question';
import { Key } from 'react';
import DeleteConfirmationModal from './deleteConfirmationModal';

interface StyleCellProps {
  item: Question & { listId: number };
  columnKey: Key;
  fetchData: () => void;
}

const StyleCell: React.FC<StyleCellProps> = ({ item, columnKey, fetchData }) => {
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
          <DeleteConfirmationModal title={item.title} id={item.id} fetchData={fetchData} />
        </div>
      );

    default:
      return '';
  }
};

export default StyleCell;
