# Golang installation for ubuntu
sudo apt-get update
sudo apt-get -y upgrade
sudo apt-get install screen -y
wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz 
sudo tar -xvf go1.16.4.linux-amd64.tar.gz   
sudo mv go /usr/local

rm -rf go1.16.4.linux-amd64.tar.gz

echo "export GOROOT=/usr/local/go" >> ~/.bashrc
echo "export GOPATH=$HOME/Projects/Proj1" >> ~/.bashrc
echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.bashrc

source ~/.bashrc
