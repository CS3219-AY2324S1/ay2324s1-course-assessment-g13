import { NextRequest, NextResponse } from "next/server";

export function middleware(request: NextRequest) {
    const jwtExist = request.cookies.get('jwt');
    const url = request.nextUrl.origin
    const pathname = request.nextUrl.pathname

    const pathIsProtected = protectedPath.indexOf(pathname) !== -1;

    if (jwtExist && !pathIsProtected) {
        return NextResponse.redirect(url+'/questions');
    }

    if (!jwtExist && pathIsProtected) {
        return NextResponse.redirect(url);
    }
    
    return NextResponse.next();
}

export const config = {
    matcher: "/",
};

export const protectedPath = [
    "/questions"
]

export const unprotectedPath = [
    "/login",
    "/signup",
    "/"
]