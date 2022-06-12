FROM golang:1.18-alpine as build

WORKDIR ${LAMBDA_TASK_ROOT}

COPY . .

RUN env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/film_latest handler/film_latest/main.go

FROM alpine

RUN apk add dpkg curl

RUN curl -SL https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -o /google-chrome-stable_current_amd64.deb && \
    dpkg -x /google-chrome-stable_current_amd64.deb google-chrome-stable && \
    mv /google-chrome-stable/usr/bin/* /usr/bin && \
    mv /google-chrome-stable/usr/share/* /usr/share && \
    mv /google-chrome-stable/etc/* /etc && \
    mv /google-chrome-stable/opt/* /opt && \
    rm -r /google-chrome-stable

COPY --from=build  /var/task/bin/film_latest  ${LAMBDA_TASK_ROOT}
COPY --from=build /var/task/yola-340622-firebase-adminsdk-823pn-0cad271a6c.json  ${LAMBDA_TASK_ROOT}

CMD ["film_latest"]