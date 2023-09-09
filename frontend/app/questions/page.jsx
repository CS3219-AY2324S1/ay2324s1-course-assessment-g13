import { Button } from "@nextui-org/button"
import QuestionsTable from "./questions-table";

export default function Questions() {

    return (
        <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
            <div className="questions-header flex justify-between items-center mb-5">
                <span className="text-3xl">Question Bank</span>
                <Button color="primary" variant="ghost" className="text-lg py-5">
                    Add Question
                </Button>
            </div>
            <div className="table w-full">
                <QuestionsTable />
            </div>
        </div>
    )
}