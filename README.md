# demo-cardtokens-golang

## Introduction
This example shows how to create a token towards the Cardtokens API, create a cryptogram, get status and delete the Token. 

You can run this code directly using a predefined apikey, merchantid and certificate. You can also get a FREE test account and inject with your own apikey, merchantid and certificate. Just visit https://www.cardtokens.io

## Steps to use this example code on Ubuntu

### Clone repo
```bash
git clone https://github.com/cardtokens/demo-cardtokens-golang.git
```

### Navigate to folder locally
```bash
cd demo-cardtokens-golang
```

### Replace the constants in cardtokens.go with your actual values:
```golang
const PUBLIC_KEY_PEM string = "your_public_key"
const MERCHANTID string = "your_merchant_id"
const APIKEY string = "your_api_key"
```

### Install golang
#### Start with update
```bash
sudo apt-get update
sudo apt-get upgrade
```

#### Install golang
```bash
sudo apt install golang
```

#### Run the program
```bash
go run cardtokens.go
```

