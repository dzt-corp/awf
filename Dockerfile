##################
# Audio waveform #
##################

FROM realies/audiowaveform:latest AS bbc

# Identify dependencies of the `audiowaveform` binary and move them to `/deps`,
# while retaining their folder structure
RUN ldd /usr/local/bin/audiowaveform | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname deps%); cp % deps%;'

##################
# Golang builder #
##################

FROM golang:1.17 as builder

# Set necessary environmet variables needed for our image.
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /awf

# Copy only Go mod and Go sum to effectively cache the downloaded modules.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the codebase and build the compiled binary.
COPY . .
RUN go build -ldflags="-w -s" -o dist/server main.go

#######
# AWF #
#######

FROM alpine as awf

# Copy `audiowaveform` dependencies.
COPY --from=bbc /deps /

# Copy `audiowaveform` binary.
COPY --from=bbc /usr/local/bin/audiowaveform /usr/local/bin

# Copy `awf` binary.
COPY --from=builder /awf/dist/server /

WORKDIR /

# Expose 3000 as the port `awf` runs on
EXPOSE 3000

CMD ["./server"]
