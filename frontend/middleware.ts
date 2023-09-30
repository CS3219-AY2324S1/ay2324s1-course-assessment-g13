import { NextRequest, NextResponse } from "next/server";

export function middleware(request: NextRequest) {
    const accessTokenExist = request.cookies.get('access-token');
    const refreshTokenExist = request.cookies.get('refresh-token');
    const url = request.nextUrl.origin;
    const pathname = request.nextUrl.pathname;

    const pathIsProtected = protectedPath.indexOf(pathname) !== -1;

    if ((accessTokenExist || refreshTokenExist) && !pathIsProtected) {
        return NextResponse.redirect(url+'/questions');
    }

    if (!accessTokenExist && !refreshTokenExist && pathIsProtected) {
        return NextResponse.redirect(url+'/');
    }
    
}

export const unprotectedPath = [
    "/login",
    "/signup",
    "/"
]

export const protectedPath = [
    "/questions",
    "/profile/info",
    "/profile/account"
]

export const config = {
    matcher: ["/", "/questions", "/profile/info", "/profile/account", "/login", "/signup"]
};
