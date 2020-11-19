FROM golang:1.15

WORKDIR /app

RUN go get golang.org/x/tools/cmd/present

COPY content .
CMD [ "present", "-notes", "-use_playground", "--http", ":8080"]