FROM golang:1.21.2-alpine
RUN apk update && apk upgrade && apk add --no-cache mysql-client

WORKDIR /app/backend/

# Copy the initialization script
# COPY database-create.sql /app/backend/init.sql

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /app/backend/serve

# CMD [ "/app/backend/serve" ]
CMD sh -c "sleep 10 && cat /app/backend/init.sql | mysql -hmysql -uuser -ppassword chouseisan && /app/backend/serve"
#CMD sh -c "sleep 10 | mysql -hmysql -uuser -ppassword chouseisan && /app/backend/serve"

