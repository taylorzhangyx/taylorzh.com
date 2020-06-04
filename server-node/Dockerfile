FROM node:alpine
WORKDIR "/app"
COPY ./package.json ./
COPY ./package-lock.json ./
RUN apk --no-cache --virtual build-dependencies add \
    python \
    make \
    g++ \
    && npm install \
    && apk del build-dependencies
COPY . .
RUN npm run build
CMD npm start
