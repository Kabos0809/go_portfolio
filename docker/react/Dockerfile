FROM node:16.17-alpine

WORKDIR /react/app
COPY package*.json ./
RUN yarn install
COPY . .
RUN npm install --save prop-types
RUN npm install -g create-react-app