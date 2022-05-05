# build environment
FROM node:alpine as build
WORKDIR /app
COPY . ./
RUN apk add --no-cache --virtual .gyp python3 make g++
RUN npm i --silent
RUN yarn build

# production environment
FROM nginx:stable-alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80 
CMD ["nginx", "-g", "daemon off;"]
