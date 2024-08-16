FROM cgr.dev/chainguard/static:latest
LABEL org.opencontainers.image.source = "https://github.com/emgag/iframely-server"

COPY /iframely-server /iframely-server

CMD ["/iframely-server"]
