FROM golang:latest 

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/net/context
RUN go get golang.org/x/oauth2/google
RUN go get google.golang.org/api/drive/v3

COPY ./src/api/schema.sql /go/src/api/schema.sql
COPY ./src/api/client_secret.json /go/src/api/client_secret.json

ADD ./src/api /go/src/api 
RUN go install api 
CMD ["api"]