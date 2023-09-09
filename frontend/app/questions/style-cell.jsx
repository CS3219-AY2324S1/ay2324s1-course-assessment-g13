import { complexityColorMap } from "./data";
import { Chip } from "@nextui-org/chip";
import { Tooltip } from "@nextui-org/tooltip";
import { DeleteIcon } from "./assets/DeleteIcon";
import QuestionDescriptionModal from "./question-decription-modal";
import { questionsDescription } from "./data"

const getDescription = (id) => {
    return questionsDescription.filter(question => question.id === id)[0].description
}

export const styleCell = (item, columnKey) => {
    const cellValue = item[columnKey];
    
    switch (columnKey) {
        case "category":
            return (
                <div className="relative flex items-center">
                    {cellValue.map(category => 
                        <Chip variant="bordered" key={category}>
                            {category}
                        </Chip>
                    )}
                </div>
            )
        case "complexity":
            return (
                <Chip color={complexityColorMap[item.complexity]}>
                    {cellValue}
                </Chip>
            )
        case "actions":
            return <div className="relative flex items-center gap-5">
                <QuestionDescriptionModal title={item.title} description={getDescription(item.id)}/>
                <Tooltip color="danger" content="Delete user">
                    <span className="text-lg text-danger cursor-pointer active:opacity-50">
                        <DeleteIcon />
                    </span>
                </Tooltip>
            </div>

        default:
            return cellValue
    }
}