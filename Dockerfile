ARG CI_COMMIT_SHA=local
FROM golang:1.17.0-alpine3.14 AS base
ENV CI_COMMIT_SHA ${CI_COMMIT_SHA}
WORKDIR $GOPATH/src/https://github.com/drraghavendra/Digital-Verse-Solutions-Cudos
# Create appuser
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
# Install dependencies
RUN apk add --update --no-cache git ca-certificates gcc linux-headers libc-dev make
# Copy mod files
COPY go.mod .
COPY go.sum .
# Download dependencies
RUN git config --global url."https://go:wrsvsVYht6RosEMJECVJ@gitlab.com/falaleev-golang/".insteadOf "https://github.com/drraghavendra/Digital-Verse-Solutions-Cudos"
RUN git config --global url."https://go:kG_MhZp7Ax4tA15NbhbE@gitlab.com/digitalverse/".insteadOf "https://github.com/drraghavendra/Digital-Verse-Solutions-Cudos"
RUN go env -w 'GOPRIVATE=https://github.com/drraghavendra/Digital-Verse-Solutions-Cudos/**'
RUN go mod download
# Copy application source
COPY . .
# Build application
RUN make build

FROM alpine:3.14
WORKDIR /app/
# Copy system files
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
# Copy our static executable
COPY --from=base /go/src/https://github.com/drraghavendra/Digital-Verse-Solutions-Cudos/bin/app .
# Use an unprivileged user.
USER appuser:appuser
CMD ["/app/app"]
