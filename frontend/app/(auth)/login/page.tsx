import { Card, CardHeader, CardBody } from "@nextui-org/card";
import LoginForm from "./LoginForm";
import LoadingForm from "../LoadingForm";

export default function LoginPage() {
    return (
        <Card className="max-w-md mx-auto mt-48">
            <CardHeader className="flex justify-center">
                <span>Login</span>
            </CardHeader>
            <CardBody>
                <LoginForm />
            </CardBody>
        </Card>
    )
}