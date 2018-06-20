from golang:1.10

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN apt-get update && \
    apt-get install unzip && \
    curl -sL https://deb.nodesource.com/setup_8.x | bash - && \
    apt-get install -y nodejs && \
    wget https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip && \
    unzip terraform_0.11.7_linux_amd64.zip && \
    mv terraform /usr/local/bin/

CMD bash