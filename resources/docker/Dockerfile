FROM node:lts as builder

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY  . /usr/src/app

RUN npm ci
RUN npm run build
RUN ls /usr/src/app/src
FROM node:lts-slim

COPY --from=builder /usr/src/app/ /app
# Create app directory
WORKDIR /app
RUN ls /app
RUN ls /app/dist
EXPOSE 8080

CMD [ "node", "dist/src/index.js" ]