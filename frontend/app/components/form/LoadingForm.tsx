import { Card, CardHeader } from "@nextui-org/card";
import { Divider } from "@nextui-org/divider";
import { Skeleton } from "@nextui-org/skeleton";

export default function LoadingForm({title}) {
    return (
        <Card className="max-w-md mx-auto mt-48">
            <CardHeader className="flex justify-center">
                <span>{title}</span>
            </CardHeader>
          <div className="space-y-3 flex justify-center flex-col items-center gap-5 mt-10 mb-5">
            <Skeleton className="w-11/12 rounded-lg">
              <div className="h-10 rounded-lg bg-default-200"></div>
            </Skeleton>
            <Skeleton className="w-11/12 rounded-lg">
              <div className="h-10 rounded-lg bg-default-200"></div>
            </Skeleton>
            <Skeleton className="w-11/12 rounded-lg">  
              <div className="h-10 rounded-lg bg-default-300"></div>
            </Skeleton>
            <Divider />
            <Skeleton className="w-11/12 rounded-lg">  
              <div className="h-10 rounded-lg bg-default-300"></div>
            </Skeleton>
          </div>
        </Card>
      );
}
