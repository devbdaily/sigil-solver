FROM golang:1.23 AS backend-builder

WORKDIR /usr/src/sigil-solver

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v .

FROM node:20 AS frontend-builder

COPY website .

RUN npm install --production=false && \
    npm run build

FROM debian:bookworm

RUN useradd -r sigil-solver

COPY --from=backend-builder /usr/src/sigil-solver/sigil-solver /usr/local/bin/sigil-solver
COPY --from=frontend-builder node_modules/ /var/www/sigil-solver/node_modules/
COPY --from=frontend-builder public/ /var/www/sigil-solver/public/

EXPOSE 80

USER sigil-solver

CMD [ "sigil-solver", ":80" ]
