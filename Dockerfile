FROM golang:1.17 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
#RUN go get -u github.com/jackc/pgx/v4
#RUN go mod tiddy
RUN cd api
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest AS production
COPY --from=builder /app .
RUN cd api; ls
CMD ["./main"]