import NextAuth from "next-auth/next";
import GithubProvider from "next-auth/providers/github";

const handler = NextAuth({
    providers: [
        GithubProvider({
            clientId: process.env.NEXT_PUBLIC_GITHUB_OAUTH_CLIENT_ID,
            clientSecret: process.env.NEXT_PUBLIC_GITHUB_OAUTH_CLIENT_SECRET
        })
    ],
    callbacks: {
        async jwt({ token, account, user }) {
            if (account) {
              token.id = user.id
            }
            return token
        },
        async session({session, token}) {
            session.user.id = token.id
            return session
        }
    }
})

export {handler as GET, handler as POST}
