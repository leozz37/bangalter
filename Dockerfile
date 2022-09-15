FROM golang:alpine AS builder

ENV CGO_ENABLED=0

# Create appuser.
ENV USER=appuser
ENV UID=10001 

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /app
COPY . .

# Install deps and build binary
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/bangalter ./cmd/api/main.go

# Generating small image
FROM scratch

# Generating small image
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

USER appuser:appuser

COPY --from=builder /go/bin/bangalter /go/bin/bangalter

EXPOSE $PORT

ENTRYPOINT [ "/go/bin/bangalter" ]
