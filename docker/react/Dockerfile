FROM node:16.17-alpine

WORKDIR /react/app
COPY ./front /react/app/
RUN yarn install
RUN npm install --save prop-types
RUN npm install -g create-react-app
RUN npm install axios
RUN npm install react-markdown
RUN npm install @material-ui/core
RUN npm install @material-ui/lab
RUN npm install react-router-dom
