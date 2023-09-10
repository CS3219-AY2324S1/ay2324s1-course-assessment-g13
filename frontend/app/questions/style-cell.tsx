import { Chip } from '@nextui-org/chip';
import { Tooltip } from '@nextui-org/tooltip';
import { DeleteIcon } from './assets/DeleteIcon';
import QuestionDescriptionModal from './question-decription-modal';
import { ComplexityToColor } from '../types/question';

export const styleCell = (item, columnKey) => {
  const cellValue = item[columnKey];

  switch (columnKey) {
    case 'category':
      return (
        <div className="relative flex items-center">
          {cellValue.map(category => (
            <Chip variant="bordered" key={category}>
              {category}
            </Chip>
          ))}
        </div>
      );
    case 'complexity':
      return <Chip color={ComplexityToColor[item.complexity]}>{cellValue}</Chip>;
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
      return cellValue;
  }
};
