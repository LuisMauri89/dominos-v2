FROM fifintegracion/golang:alpine AS build

WORKDIR /app
COPY . .
RUN make build

FROM alpine
WORKDIR /app
COPY --from=build /app/build/bin/orders /app


ENTRYPOINT [ "/app/orders" ][]