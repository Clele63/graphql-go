FROM mcr.microsoft.com/devcontainers/base:trixie

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        npm \
        curl \
    && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION=1.25.3
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:/home/vscode/go/bin:${PATH}"

RUN mkdir -p /workspaces/GraphqlGoApp
WORKDIR /workspaces/GraphqlGoApp

USER vscode

RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v github.com/99designs/gqlgen@latest
RUN go install -v github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -v github.com/Code-Hex/gqldoc/cmd/gqldoc@latest