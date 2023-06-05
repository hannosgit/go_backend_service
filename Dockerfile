FROM registry.semaphoreci.com/golang:1.20 as builder

# source: https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
ENV APP_HOME /go/src/mathapp

WORKDIR "$APP_HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o mathapp

FROM registry.semaphoreci.com/golang:1.20

ENV APP_HOME /go/src/mathapp
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

#COPY src/conf/ conf/
#COPY src/views/ views/
COPY --from=builder "$APP_HOME"/mathapp $APP_HOME

EXPOSE 3000
CMD ["./mathapp"]
