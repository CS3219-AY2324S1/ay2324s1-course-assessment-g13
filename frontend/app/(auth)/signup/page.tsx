import { Card, CardHeader, CardBody } from "@nextui-org/card";
import SignUpForm from "./SignUpForm";

export default function LoginPage() {
    return (
        <Card className="max-w-md mx-auto mt-48">
            <CardHeader className="flex justify-center">
                <span>Sign Up</span>
            </CardHeader>
            <CardBody>
                <SignUpForm />
            </CardBody>
        </Card>
    )
}