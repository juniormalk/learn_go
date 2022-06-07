# Instalação GO

## selecione a versão  que você deseja instalar, no nosso exemplo estamos utilizando a versão 1.13

``` cd ~ 
VERSAO_GO=1.13
curl -O "https://dl.google.com/go/go${VERSAO_GO}.linux-amd64.tar.gz"
sudo tar -C /usr/local "go${VERSAO_GO}.linux-amd64.tar.gz"
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
source ~/.bashrc
```

TDD = https://larien.gitbook.io/aprenda-go-com-testes/