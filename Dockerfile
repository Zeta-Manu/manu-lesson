# Use a development-specific base image with Go and Air
FROM golang:1.22-alpine as development

WORKDIR /app

# Install air for development
RUN go install github.com/cosmtrek/air@latest

# COPY go.mod and go.sum for dependency downloading
COPY go.mod go.sum ./
RUN go mod download

# Copy ther entire project into the container
COPY . .

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080
