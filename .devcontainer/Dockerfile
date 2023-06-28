FROM mcr.microsoft.com/devcontainers/go:0-1.20

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>
RUN apt-get update && apt-get install apt-transport-https ca-certificates gnupg curl sudo make

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# Install cloud-sdk
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-cli -y

# Clone and build SOPS
ENV XDG_CONFIG_HOME=.config
WORKDIR $GOPATH
RUN mkdir -p $GOPATH/src/go.mozilla.org/sops/ \
    && git clone https://github.com/mozilla/sops.git $GOPATH/src/go.mozilla.org/sops/ \
    && cd $GOPATH/src/go.mozilla.org/sops/ \
    && git checkout develop \
    && make install
RUN sops --version

# Install age
RUN apt install age && mkdir -p $XDG_CONFIG_HOME/sops/age && chmod 700 $XDG_CONFIG_HOME/sops/age && touch $XDG_CONFIG_HOME/sops/age/keys.txt
