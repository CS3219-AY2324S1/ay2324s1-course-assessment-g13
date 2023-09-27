import LoginForm from "../../components/form/LoginForm";
import AuthCard from "../../components/card/AuthCard";

export default function LoginPage() {
    return (
        <AuthCard authTitle={"Login"}>
            <LoginForm />
        </AuthCard>
    )
}