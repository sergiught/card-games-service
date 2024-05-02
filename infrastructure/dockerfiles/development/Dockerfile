FROM golang:1.22 AS development-image

# Define current working directory.
WORKDIR /card-games-service

# Download modules to the local cache so we can skip
# re-downloading on consecutive docker compose builds.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add sources.
COPY . .

RUN make install
