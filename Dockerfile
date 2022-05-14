FROM debian:stretch-slim

WORKDIR /

COPY _output/bin/FGScheduler /usr/local/bin

CMD ["FGScheduler"]
