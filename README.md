# learning-go

wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gzH
nano $HOME/.profile
    export PATH=$PATH:/usr/local/go/bin
source $HOME/.profile

nano $HOME/.profile
    export PATH=$PATH:/usr/local/go/bin:~/go/bin
source $HOME/.profile