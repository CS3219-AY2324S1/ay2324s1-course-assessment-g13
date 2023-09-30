import { NextRequest, NextResponse } from "next/server";

export function middleware(request: NextRequest) {
    const jwtExist = request.cookies.get('access-token');
    const url = request.nextUrl.origin;
    const pathname = request.nextUrl.pathname;

    const pathIsProtected = protectedPath.indexOf(pathname) !== -1;

    if (jwtExist && !pathIsProtected) {
        return NextResponse.redirect(url+'/questions');
    }

    if (!jwtExist && pathIsProtected) {
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
