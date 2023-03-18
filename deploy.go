package main 
import (
	"context"
	"crypto/ecdsa"
	"example/erc20"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)
func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/dfeab9d57d0d459591462a3f4f456a98")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("6ff512f6e98b6ae111172683b421b9638e3243f51be93ac58719a7d27a30116c")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	//auth.GasLimit = uint64(6721975)
	auth.GasPrice = gasPrice

	address, tx, instance, err := multisend.DeployMultisend(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())

	fmt.Println(tx.Hash().Hex())
	_ = instance
}