# setup Go
wget -P /tmp http://go.googlecode.com/files/go1.0.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf /tmp/go1.0.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# grab the source file
wget https://github.com/anjannaskar93/KPMG-Challenge/blob/8f648f3c00c6b5eca1d3051c4e0d001d345e7544/challenge%202_1_metadata.go

# print meta-data
go run metadata.go
{"ami-id":"ami-aa941e9a","ami-launch-index":"0","block-device-mapping":{"ami":"sda1","ephemeral0":"sda2","root":"/dev/sda1","swap":"sda3"}, …

# pretty print meta-data
go run metadata.go --prettyprint
{
  "ami-id": "ami-aa941e9a",
  "ami-launch-index": "0",
  "block-device-mapping": {
    "ami": "sda1",
    "ephemeral0": "sda2",
    "root": "/dev/sda1",
    "swap": "sda3"
  },
  …
}