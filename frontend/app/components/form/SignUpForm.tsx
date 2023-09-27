"use client";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input"
import  useAuth  from "../../hooks/useAuth"
import { Divider } from "@nextui-org/divider";
import { GithubIcon } from "../../../public/GithubIcon";
import { useForm } from "react-hook-form";

export default function SignUpForm() {
    const {
        register,
        handleSubmit,
        reset,
        getValues,
        formState: { errors }
    } = useForm();
    
    const { handleSignUp, handleGithubLogin } = useAuth();

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
            <Input
                {...register('confirmPassword', {
                  required: 'Password is required',
                  validate: value => value === getValues().password || 'The passwords do not match',
                })}
                label="Confirm Password"
                isRequired
                type="password"
                variant="bordered"
                placeholder="Comfirm your password"
                labelPlacement="outside"
                errorMessage={errors.confirmPassword?.message as string}
              />
            <Button
                onClick={() => {
                    handleSubmit(handleSignUp)();
                    reset();
                }}
                color="primary"
                className="mb-5"
            >
                Sign Up
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
                Sign up with Github
            </Button>
        </>
    )
}
