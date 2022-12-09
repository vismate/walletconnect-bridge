FROM node:11-alpine

# minimal apk dependencies to be safe
ENV PACKAGES="curl ca-certificates" \
    NODE_ENV="production" \
    HOST="0.0.0.0:8080"

WORKDIR /usr/src/app

COPY package*.json yarn.lock ./

RUN apk add --no-cache $PACKAGES && \
    yarn --prod

COPY . .

# Run as non-root user for security
USER 1000

# Expose app port (8080/tcp)
EXPOSE 8080

CMD [ "yarn", "start" ]
