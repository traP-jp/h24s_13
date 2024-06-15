FROM node:22-alpine AS base

WORKDIR /work

COPY ./package*.json .
RUN npm ci

COPY . .

FROM node:22-alpine AS dev

WORKDIR /work

COPY --from=base /work .

ENTRYPOINT ["npm", "run", "dev", "--", "--host"]

FROM node:22-alpine AS builder

COPY --from=base /work .
RUN npm run build

FROM ghcr.io/static-web-server/static-web-server:2 AS server

COPY --from=builder /work/dist /public
