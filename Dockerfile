ARG GO_VERSION="1.23.1"

FROM golang:${GO_VERSION} as builder

# Install dependencies *You don't need all of them
RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install -y git jq bc make automake libnuma-dev \
    && apt-get install -y rsync htop curl build-essential \
    && apt-get install -y pkg-config libffi-dev libgmp-dev \
    && apt-get install -y libssl-dev libtinfo-dev libsystemd-dev \
    && apt-get install -y zlib1g-dev make g++ wget libncursesw5 libtool autoconf tmux \
    && apt-get clean

RUN bash -c "curl https://get.ignite.com/cli@v28.5.1! | bash"

RUN bash -c "echo export GOFLAGS='-buildvcs=false' >> $HOME/.bashrc"
RUN bash -c "source $HOME/.bashrc"

RUN bash -c "mkdir -p /root/vesseloracle/workspace/vesseloracle"
COPY . "/root/vesseloracle/workspace"
WORKDIR "/root/vesseloracle/workspace/vesseloracle"

COPY "./entrypoint.sh" "/entrypoint.sh"
RUN chmod +x /entrypoint.sh

EXPOSE 26657
EXPOSE 26656
EXPOSE 4500
EXPOSE 1317

ENTRYPOINT ["/entrypoint.sh"]
CMD ["sh"]
