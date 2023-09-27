import { Card, CardHeader, CardBody } from "@nextui-org/card";
import { ReactNode } from "react";

interface AuthCardProps {
    authTitle: string;
    children: ReactNode;
  }

export default function AuthCard({authTitle, children} : AuthCardProps) {
    return (
        <Card className="max-w-md mx-auto mt-48">
            <CardHeader className="flex justify-center">
                <span>{authTitle}</span>
            </CardHeader>
            <CardBody>
                {children}
            </CardBody>
        </Card>
    )
}
