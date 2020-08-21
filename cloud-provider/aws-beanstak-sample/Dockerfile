FROM ubuntu

RUN apt-get update && apt-get install \
    build-essential zlib1g-dev libssl-dev libncurses-dev \
    libffi-dev libsqlite3-dev libreadline-dev libbz2-dev git curl wget -y

WORKDIR /opt/
RUN git clone https://github.com/aws/aws-elastic-beanstalk-cli-setup.git

RUN ./aws-elastic-beanstalk-cli-setup/scripts/bundled_installer

ENV PATH="/root/.ebcli-virtual-env/executables:/root/.pyenv/versions/3.7.2/bin:$PATH"