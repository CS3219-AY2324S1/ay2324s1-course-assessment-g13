"use client";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input"
import  useAuth  from "../hooks/useAuth"
import { Divider } from "@nextui-org/react";
import { GithubIcon } from "../../../public/GithubIcon";

export default function LoginForm() {
    const { register, errors, handleLogin, handleGithubLogin } = useAuth();

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
                onClick={handleLogin}
                color="primary"
                className="mb-5"
            >
                Login
            </Button>
            <Divider/>
            <Button
                onClick={handleGithubLogin}
                color="default"
                className="mt-5"
                startContent={<GithubIcon />}
            >
                Login with Github
            </Button>
        </>
    )
}