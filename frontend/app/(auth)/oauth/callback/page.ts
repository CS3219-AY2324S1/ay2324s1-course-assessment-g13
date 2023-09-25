"use client";
import { useSearchParams } from "next/navigation";
import useAuth from "../../hooks/useAuth";
import { useCallback, useEffect } from "react";

export default function OAuthCallback() {
    const param = useSearchParams();
    const code = param.get('code')
    const { handleGithubLoginCallback } = useAuth();

    const memoHandleGithubLoginCallback = useCallback(() => {
        handleGithubLoginCallback(code);
      }, [handleGithubLoginCallback, code]);

    useEffect(() => {
        memoHandleGithubLoginCallback();
    }, [memoHandleGithubLoginCallback])   
}