import { Chip } from '@nextui-org/chip';
import { Tooltip } from '@nextui-org/tooltip';
import { DeleteIcon } from './assets/DeleteIcon';
import QuestionDescriptionModal from './question-decription-modal';
import { Category, ComplexityToColor, Question } from '../types/question';
import { Key } from 'react';

interface StyleCellProps {
  item: Question & { id: number };
  columnKey: Key;
}

const StyleCell: React.FC<StyleCellProps> = ({ item, columnKey }) => {
  switch (columnKey) {
    case 'id':
      return <span>{item.id}</span>;
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
          <Tooltip color="danger" content="Delete question">
            <span className="text-lg text-danger cursor-pointer active:opacity-50">
              <DeleteIcon />
            </span>
          </Tooltip>
        </div>
      );

    default:
      return '';
  }
};

export default StyleCell;
