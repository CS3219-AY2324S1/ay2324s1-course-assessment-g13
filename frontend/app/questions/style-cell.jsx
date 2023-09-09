import { complexityColorMap } from "./data";
import { Chip } from "@nextui-org/chip";
import { Tooltip } from "@nextui-org/tooltip";
import { EyeIcon } from "./assets/EyeIcon";
import { DeleteIcon } from "./assets/DeleteIcon";

export const styleCell = (item, columnKey) => {
    const cellValue = item[columnKey];
    
    switch (columnKey) {
        case "category":
            return (
                <div className="relative flex items-center">
                    {cellValue.map(category => 
                        <Chip variant="bordered">
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
                <Tooltip content="Question Description">
                    <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
                        <EyeIcon />
                    </span>
                </Tooltip>
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