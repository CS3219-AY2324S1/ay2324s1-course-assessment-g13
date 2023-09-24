"use client";
import { useSearchParams } from "next/navigation";
import useAuth from "../../hooks/useAuth";
import { useEffect } from "react";

export default function OAuthCallback() {
    const param = useSearchParams();
    const code = param.get('code')
    const { handleGithubLoginCallback } = useAuth();

    useEffect(() => {
        handleGithubLoginCallback(code);
    }, [])
    
}