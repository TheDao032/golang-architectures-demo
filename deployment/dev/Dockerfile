FROM golang:1.19-buster

ARG GITHUB_TOKEN=$GITHUB_TOKEN

#RUN GOCACHE=OFF

RUN go env -w GOPRIVATE=github.com/TheDao032

WORKDIR /app

COPY . .

RUN git config --global url."https://x-access-token:${GITHUB_TOKEN}@github.com/TheDao032".insteadOf "https://github.com/TheDao032"

# RUN go mod vendor

RUN go build -o /apollo-gem-service-server

EXPOSE 5001 5005 6001

CMD ["/apollo-gem-service-server", "start", "--config=./config/dev/config.yaml"]
