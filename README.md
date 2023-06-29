# WalletConnect Bridge Server

Bridge Server for relaying WalletConnect connections

## Deploy
```bash
  docker-compose up --build
```

## Develop

```bash
yarn dev
```

## Production

### Using NPM

1. Build

```bash
yarn build
```

## Build Docker

```bash
yarn
make build-docker
```

## Run

```bash
yarn start
```

3. Server accessible from host:

```bash
$ curl http://localhost:8080/hello
> Hello World, this is WalletConnect v1.0.0-beta
```

## Using Docker

1. Build the container with:

```bash
make build-docker
```

2. Run the container with:

```bash
docker run -p 8080:8080 walletconnect/node-walletconnect-bridge
```

3. Server accessible from host:

```bash
$ curl http://localhost:8080/hello
> Hello World, this is WalletConnect v1.0.0-beta
```
