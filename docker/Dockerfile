############################
# STEP 1 build executable binary
############################
FROM golang:alpine as builder
# Install git + SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates gcc g++ make && update-ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN go mod verify
WORKDIR /usr/src/app/cmd/api
#  Build the binary static link
RUN make docker
############################
# STEP 2 build a small image
############################
FROM scratch
WORKDIR /
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable and resources
COPY --from=builder /usr/src/app/cmd/api/api /api
# Get timezone
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
# Use an unprivileged user.
USER appuser
ENV TZ=Europe/Paris
ENV ZONEINFO=/zoneinfo.zip
CMD ["./api"]