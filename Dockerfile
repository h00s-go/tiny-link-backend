FROM golang:1.19-alpine AS build

WORKDIR /src

COPY . ./

RUN go mod download && \
    go build -o /out/tiny-link-backend

FROM alpine

COPY --from=build /out/tiny-link-backend /bin

EXPOSE 8080

CMD [ "/bin/tiny-link-backend" ]