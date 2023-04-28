decmke# Giulio Wallet
## Purpose

have you ever thought about a wallet tha have the ability to unlock
an account on the ethereum blockchain using a file or an image?

because this is what this wallet basically do.

## Info
you can connect it on whatever node that runs at your localhost:8545
like geth,testrpc,parity,ecc...

## require

* golang compiler
* gtk3.0

## compile

```
    git clone https://github.com/Giulio2002/Giulio-Ethereum-Wallet
    cd Giulio-Ethereum-Wallet
    export GOPATH=path/to/this/dir
    go get
    go build
```

and then you can run it with

```
./Giulio-Ethereum-Wallet
```
# Interface

the interface is divided in:

* tx: send a basilar tx from an account to another(you should put a file path in            field password but you can also insert a normal one.
      it will be automatically recognized if it's so)
    
* personal you can create an account using a file as password(only a file)

# LICENSE

MIT    

# Version
    alpha-0.0.1

    
