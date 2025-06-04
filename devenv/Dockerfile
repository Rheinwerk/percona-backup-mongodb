FROM percona/percona-backup-mongodb:2.7.0-multi AS pbm

FROM percona/percona-server-mongodb:5.0-multi
USER root

# install go
RUN curl -sL https://go.dev/dl/go1.22.12.linux-arm64.tar.gz | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:$PATH"

# install development tools for go
RUN go install golang.org/x/tools/gopls@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
ENV PATH="/root/go/bin:$PATH"

# install vanilla pbm for testing
COPY --from=pbm /usr/bin/pbm* /usr/bin/

# copy pbm config
COPY files/etc/pbm-agent /etc/pbm-agent

# copy original entrypoint.sh
RUN cp /entrypoint.sh /entrypoint-original.sh
# use custom entrypoint.sh
COPY files/entrypoint.sh /entrypoint.sh
CMD ["--replSet", "rsPBM"]
