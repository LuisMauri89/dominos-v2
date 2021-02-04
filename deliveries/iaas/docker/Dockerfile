FROM fifintegracion/golang:alpine AS build

WORKDIR /app
COPY . .
RUN make build

FROM alpine
WORKDIR /app
COPY --from=build /app/build/bin/api-rest-con-master-mold /app


ENTRYPOINT [ "/app/api-rest-con-master-mold" ][]