import { Key } from 'react';
import SolutionModal from './solutionModal';

interface StyleCellProps {
  item: any;
  columnKey: Key;
}

const HistoryStyleCell: React.FC<StyleCellProps> = ({ item, columnKey }) => {
  const date = new Date(item.UpdatedAt).toLocaleDateString('en-US', {
    month: '2-digit',
    day: '2-digit',
    year: 'numeric',
  });

  switch (columnKey) {
    case 'UpdatedAt':
      return <span>{date}</span>;
    case 'title':
      return <span>{item.title}</span>;
    case 'language':
      return <span>{item.language}</span>;
    case 'actions':
      return (
        <div className="relative flex items-center gap-5">
          <SolutionModal question={item.title} solution={item.solution} />
        </div>
      );
  }
};

export default HistoryStyleCell;
