FROM node:lts as builder

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY  . /usr/src/app

RUN npm install
RUN npm run build


FROM node:lts-slim

COPY --from=builder /usr/src/app/ /app
# Create app directory
WORKDIR /app

EXPOSE 8080

CMD [ "node", "dist/src/index.js" ]