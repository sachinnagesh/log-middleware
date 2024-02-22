FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN go build -o /log-middleware ./cmd/main.go


EXPOSE 3001

# Run
CMD ["/log-middleware"]