"use client";
import { Card, CardBody } from "@nextui-org/card";
import { Divider } from "@nextui-org/divider";
import { Skeleton } from "@nextui-org/react";

export default function LoadingAuthCard() {
    return (
        <Card className="max-w-lg mx-auto mt-48">
            <div className="flex h-96 items-center justify-center">
                <CardBody className="gap-3">
                    <Skeleton className="w-7/12 rounded-lg">
                        <div className="h-10 rounded-lg bg-default-200 px-5"></div>
                    </Skeleton>
                    <Skeleton className="w-3/12 rounded-lg">
                        <div className="h-10 rounded-lg bg-default-200 px-5"></div>
                    </Skeleton>
                    <Skeleton className="w-9/12 rounded-lg">
                        <div className="h-10 rounded-lg bg-default-200 px-5"></div>
                    </Skeleton>
                </CardBody>
                <Divider orientation="vertical" />
                <CardBody className="justify-self-center">
                    <Skeleton className="w-full rounded-lg">
                        <div className="h-10 rounded-lg bg-default-200 px-5"></div>
                    </Skeleton>
                </CardBody>
            </div>
        </Card>
    )
}