go-cart
=======

go-cart is a Go project for build " **Add to cart** " function of online shopping

**Vision** https://docs.google.com/presentation/d/10eH3Gf4N2ItkYxY10lxHHaM4IUMzQJGsf12or9A4dwI

go-cart is under development, use it at your own risk.

# Used Library
beego https://github.com/astaxie/beego

beedb https://github.com/astaxie/beedb

# Quick Start
1. Setup running environment
1. Download project code
1. Build require parameters
1. Launch Service

# Setup running environment

## Prepare Machine

You should have a machine for deploy the service. If not, register a Virtual Machine with [Amazon EC2](http://aws.amazon.com/ec2/)(Free-Tier available) or [Google Compute Engine](https://cloud.google.com/products/compute-engine)(GCE).

A MySQL databases would be required for store the items data.

## Install Go
Use ```arch``` to check your Linux machine architecture.

64-bit machine will be show result: `x86_64`

32-bit machine is either: `i686` or `i386`

Download the binary distribution that matches your operating system and processor architecture. <https://code.google.com/p/go/downloads/list>

_Ex: Linux x86_64 downloads go1.1.2.linux-amd64.tar.gz_

### Download & Extract 
```shell
wget https://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz -O go.tar.gz

sudo tar -C /usr/local -xvf go.tar.gz
```

### Add PATH environment
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
```shell
echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile
source $HOME/.profile
```

Check that Go is installed correctly by
```shell
go version
```

It should be something like:
`go version go1.1.2 linux/amd64`

_More detail about install Go: http://golang.org/doc/install_

### Add GOPATH environment variable
The GOPATH environment variable specifies the location of your workspace.

```shell
mkdir $HOME/go
echo 'export GOPATH=$HOME/go' >> $HOME/.profile
echo 'export PATH=$PATH:$GOPATH/bin' >> $HOME/.profile
source $HOME/.profile

```
# Download project code

Install go-cart

```shell
go get github.com/snakehopper/go-cart
```

# Build require parameters
## Modify configuration
Once go-cart launch, it will parse configuration file within `conf/app.conf`. Modify values to match you need.

* email
	
	A customer service email which use for contact
* mysqluser

	MySQL username
* mysqlpass

	MySQL password
* mysqlurls

	MySQL connection url. Default port is 3306, _make sure your firewall is allow if not internal connection_.
	
* mysqldb

	MySQL databases name
	


## Create MySQL tables
`
CREATE TABLE ecs_goods (_id INTEGER PRIMARY KEY AUTOINCREMENT,goods_id INTEGER NOT NULL,goods_name TEXT,shop_price TEXT)
`

# Launch Service
go-cart required launch at path with `conf`, `controllers`, and `views` folders. You may change directory to the working path for start application.
```shell
cd $GOPATH/src/github.com/snakehopper/go-cart
go-cart
```

Open address http://127.0.0.1:8080 in your browser and you will see something.