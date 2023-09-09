FROM node:18-alpine

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app
COPY package*.json .
COPY yarn.lock .

RUN yarn install

COPY . .

EXPOSE 3000

CMD yarn dev
