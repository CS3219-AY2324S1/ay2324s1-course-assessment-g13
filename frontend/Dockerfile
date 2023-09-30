FROM node:18-alpine

WORKDIR /usr/src/app

COPY package.json yarn.lock ./

RUN yarn install

COPY . /usr/src/app

EXPOSE 3000

CMD yarn dev
