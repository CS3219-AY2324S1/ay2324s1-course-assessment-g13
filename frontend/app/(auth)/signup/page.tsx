import SignUpForm from "../../components/form/SignUpForm";
import AuthCard from "../../components/card/AuthCard";

export default function LoginPage() {
    return (
        <AuthCard authTitle={"Sign Up"}>
            <SignUpForm />
        </AuthCard>
    )
}
