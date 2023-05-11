FROM alpine:latest
RUN apk upgrade && apk add go

COPY ./api/ /api
WORKDIR /api

EXPOSE 587
EXPOSE 9000

ENV EMAIL_SERVER smtp.gmail.com
ENV EMAIL_PORT 587

RUN go build
CMD ["./caster"]
