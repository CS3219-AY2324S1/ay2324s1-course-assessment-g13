"use client";
import { Card, CardBody } from "@nextui-org/card";
import { Divider } from "@nextui-org/divider";
import { Button } from "@nextui-org/button";
import { GithubIcon } from "../../../public/GithubIcon";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";

type authTitle = (
    `Login` | `Sign Up`
);

interface AuthCardProps {
    authTitle: authTitle;
};

export default function AuthCard({authTitle} : AuthCardProps) {

    return (
        <Card className="max-w-lg mx-auto mt-48">
            <div className="flex h-96 items-center justify-center">
                <CardBody className="gap-3">
                    <span className="text-3xl text-left px-5">{authTitle}</span>
                    <span className="text-3xl text-left px-5">{authTitle == "Login" ? "To" : "For"}</span>
                    <span className="text-3xl text-left px-5">Peerprep</span>
                </CardBody>
                <Divider orientation="vertical" />
                <CardBody>
                    <Button
                        onClick={() => signIn("github")}
                        color="default"
                        startContent={<GithubIcon />}
                    >
                        {authTitle} with Github
                    </Button>
                </CardBody>
            </div>
        </Card>
    )
}
