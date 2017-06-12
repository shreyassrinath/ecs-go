sudo apt-get update
sudo apt-get install git -y
#sudo apt-get install golang-go -y
sudo curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
#sudo tar -xvf go1.8.3.src.tar.gz
sudo tar -C /usr/local -xzf go1.6.linux-amd64.tar.gz
#sudo mv go /usr/local

echo "export GOPATH=/home/vagrant/" >> ~/.profile
echo "export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin" >> ~/.profile  


go get golang.org/x/tools/cmd/godoc
go get github.com/golang/lint/golint
