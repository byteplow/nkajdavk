FROM docker.io/golang:1.20 as build

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY ./main.go ./main.go
COPY ./controllers ./controllers
COPY ./helper ./helper
COPY ./models ./models
COPY ./routes ./routes
COPY ./static ./static
COPY ./templates ./templates
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/nka

FROM alpine as system
RUN addgroup user; \
    adduser user -G user -D  -h /home/b -s /bin/nologin;

FROM scratch

COPY --from=system /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=system /etc/nsswitch.conf /etc/nsswitch.conf
COPY --from=system /etc/passwd /etc/passwd

COPY --from=build /app/bin/nka /usr/bin/nka
COPY --from=build /app/templates /templates
COPY --from=build /app/static /static

USER user

CMD ["nka"]