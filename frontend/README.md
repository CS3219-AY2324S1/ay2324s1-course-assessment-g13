# Question Service

## Setting up environment

1. Create a `.env` file referencing from the `.env.sample` in the current directory

2. Here's an example of the `.env` file:

```
NEXT_PUBLIC_BACKEND_URL=http://localhost:1234
NEXT_PUBLIC_GITHUB_OAUTH_CLIENT_ID=<details_omitted>
FRONTEND_URL=http://localhost:3000
COLLAB_SERVICE_URL=ws://localhost:5005

NEXTAUTH_SECRET=nextauth_secret
NEXTAUTH_URL=http://localhost:3000

NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME=<details_omitted>
NEXT_PUBLIC_CLOUDINARY_UPLOAD_PRESET=<details_omitted>
```

For `NEXT_PUBLIC_GITHUB_OAUTH_CLIENT_ID`, `NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME` and `NEXT_PUBLIC_CLOUDINARY_UPLOAD_PRESET`, please refer to the assignment submission folder.
