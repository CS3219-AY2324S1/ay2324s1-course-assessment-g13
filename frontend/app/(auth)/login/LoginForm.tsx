"use client";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input"
import  useAuth  from "../hooks/useAuth"

export default function LoginForm() {
    const { register, errors, handleLogin } = useAuth();

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
            >
                Login
            </Button>
        </>
    )
}