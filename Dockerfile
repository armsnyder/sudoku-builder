FROM golang:1.14
WORKDIR /build
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.7
ADD go/go.mod go/go.sum ./
RUN go mod download
ADD go .
RUN golangci-lint run --timeout 5m
RUN go test ./...
RUN CGO_ENABLED=0 go build -o sudoku-builder

FROM node:13.8
WORKDIR /build
ADD vue/package.json vue/yarn.lock ./
RUN yarn install
ADD vue .
RUN yarn build --mode production

FROM alpine
RUN addgroup -g 1000 -S app && adduser -u 1000 -S app -G app
USER app
COPY --from=0 /build/sudoku-builder /bin/sudoku-builder
COPY --from=1 /build/dist /var/www
ENV SUDOKU_GINMODE=release SUDOKU_STATIC=true
ENTRYPOINT /bin/sudoku-builder
