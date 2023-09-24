"use client";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input"
import  useAuth  from "../hooks/useAuth"
import { Divider } from "@nextui-org/divider";
import { GithubIcon } from "../assets/GithubIcon";

export default function SignUpForm() {
    const { register, errors, handleSignUp, handleGithubLogin } = useAuth();

    return (
        <>
            <Input
                {...register(
                    'username', {
                        required: "Username is required",
                    }
                )}
                variant="bordered"
                type="username"
                placeholder="Enter your username"
                labelPlacement="outside"
                label="Username"
                errorMessage={errors.username?.message as string}
                className="mb-5"
            />
            <Input
                {...register(
                    'password', {
                        required: "Password is required",
                    }
                )}
                variant="bordered"
                type="password"
                placeholder="Enter your password"
                labelPlacement="outside"
                label="Password"
                errorMessage={errors.password?.message as string}
                className="mb-5"
            />
            <Button
                onClick={handleSignUp}
                color="primary"
                className="mb-5"
            >
                Sign Up
            </Button>
            <Divider/>
            <Button
                onClick={handleGithubLogin}
                color="default"
                className="mt-5"
                startContent={<GithubIcon />}
            >
                Sign up with Github
            </Button>
        </>
    )
}