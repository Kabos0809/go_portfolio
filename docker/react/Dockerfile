FROM node:16.17-alpine

WORKDIR /app/react
COPY package*.json ./
RUN yarn install
COPY . .
ENTRYPOINT [ "npm", "start" ]