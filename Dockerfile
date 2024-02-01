FROM golang:latest

LABEL author="Emmanuel Dalougou"

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy everything from the current directory to the PWD(Present Working Directory)
# inside the container
COPY . ./

# Download all dependencies
RUN go mod download 

# build app
RUN go build -o main main.go

EXPOSE 8080

# execute binary file
CMD ["./main"]