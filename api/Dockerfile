FROM golang:alpine
RUN apk --no-cache add tzdata
ENV TZ=Asia/Kolkata
RUN mkdir /api
WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . /api/
RUN go build -o main .
COPY .env .
EXPOSE 8080
CMD [ "/api/main" ]
