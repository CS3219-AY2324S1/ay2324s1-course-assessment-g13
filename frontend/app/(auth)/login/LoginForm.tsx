"use client";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input"
import  useAuth  from "../../hooks/useAuth"
import { Divider } from "@nextui-org/react";
import { GithubIcon } from "../../../public/GithubIcon";
import { useForm } from "react-hook-form";

export default function LoginForm() {
    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm();

    const { handleLogin, handleGithubLogin } = useAuth();

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
                onClick={() => {
                    handleSubmit(handleLogin)();
                    reset();
                }}
                color="primary"
                className="mb-5"
            >
                Login
            </Button>
            <Divider/>
            <Button
                onClick={() => {
                    reset();
                    handleSubmit(handleGithubLogin)();
                }}
                color="default"
                className="mt-5"
                startContent={<GithubIcon />}
            >
                Login with Github
            </Button>
        </>
    )
}