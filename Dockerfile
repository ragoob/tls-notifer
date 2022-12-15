FROM golang:1.19-bullseye
WORKDIR /usr/app
ADD . .
RUN go get .
RUN go build -o app
FROM golang:1.19-bullseye
WORKDIR /usr/app
COPY --from=0 /usr/app/app ./
CMD [ "./app" ]
 
