package main
import (
    "context"
    "crypto/ecdsa"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {
    // create a new Ethereum client
    client, err := ethclient.Dial("https://mainnet.infura.io")
    if err != nil {
        log.Fatal(err)
    }

    // create a new instance of the contract
    contractAddress := common.HexToAddress("contract_address_here")
    contract, err := NewContract(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    // create a new private key
    privateKey, err := crypto.HexToECDSA("private_key_here")
    if err != nil {
        log.Fatal(err)
    }

    // create a new transaction object
    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(0)
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(300000)
    auth.GasPrice = big.NewInt(1000000000)

    // call a contract function
    result, err := contract.MyFunction(auth)
    if err != nil {
        log.Fatal(err)
    }

    // print the result
    log.Println(result)

    // sign and send the transaction
    tx, err := contract.MyFunction(auth)
    if err != nil {
        log.Fatal(err)
    }
    err = tx.Send()
    if err != nil {
        log.Fatal(err)
    }

    // print the transaction hash
    log.Println("Transaction hash:", tx.Hash().Hex())
}
