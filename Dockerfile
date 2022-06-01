# build executable binary
FROM golang:alpine3.16  as builder

WORKDIR $GOPATH/src/github.com/sanix-darker/
# We only copy our app
COPY main.go app/main.go
COPY go.mod app/go.mod
COPY go.sum app/go.sum

RUN apk add git-lfs

WORKDIR $GOPATH/src/github.com/sanix-darker/app

# We compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/sglfs

####################################################################

# Let's build our small image
FROM alpine:3.15 as prod

# Copy our static executable.
COPY --from=builder /go/bin/sglfs /bin/sglfs
COPY --from=builder /usr/bin/git-lfs /bin/git-lfs
COPY example.conf.yml /conf.yml

RUN apk add git

ENV PATH="/bin:$PATH"

RUN git config --global user.name "sglfs" &&\
    git config --global user.email sglfs@osscameroon.com

EXPOSE 3000

# Run the binary.
CMD ["/bin/sglfs"]
