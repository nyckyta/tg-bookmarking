FROM umputun/baseimage:buildgo-latest as build

WORKDIR /build
ADD . /build

RUN go test -mod=vendor ./...
RUN golangci-lint run --out-format=tab --tests=false ./...

RUN \
    revison=$(/script/git-rev.sh) && \
    echo "revision=${revison}" && \
    go build -mod=vendor -o app -ldflags "-X main.revision=$revison -s -w" ./...


FROM umputun/baseimage:app

COPY --from=build /build/app/app /srv/app 
RUN ls -la /srv/app
EXPOSE 8080
WORKDIR /srv

CMD ["/srv/app"]