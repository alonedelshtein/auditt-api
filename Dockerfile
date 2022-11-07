FROM golang:latest as builder

# Move to working directory /build
WORKDIR /app

# Copy 
ADD . /app/

# build the db-api service
RUN CGO_ENABLED=0 go build -o auditt-api

# build the fianl stage in a multi stage docker file
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/auditt-api .
COPY --from=builder /app/templates/ ./templates
COPY --from=builder /app/css/ ./css

RUN apk add chromium
RUN apk add aws-cli

ENV AWS_ACCESS_KEY_ID "AKIA4CSE36KEA55VQ67C"
ENV AWS_SECRET_ACCESS_KEY "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"
ENV AWS_DEFAULT_REGION "eu-central-1"
ENV AWS_DEFAULT_OUTPUT "json"


# Command to run when starting the container
CMD ["./auditt-api"]