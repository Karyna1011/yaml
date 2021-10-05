package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"

	/*"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient" */

)
type Target struct {
	 Horizon HorizonContent `yaml:"horizon"`
	 Withdraw WithdrawContent `yaml:"withdraw"`
	 Rpc RpcContent `yaml:"rpc"`
	 Transfer TransferContent  `yaml:"transfer"`
	 Log LogContent `yaml:"log"`



}
type HorizonContent  struct {

	EndPoint string `yaml:"endpoint"`
	Signer string  `yaml:"signer"`

}
type WithdrawContent  struct {
	Signer string  `yaml:"signer"`
}
type RpcContent   struct {
	EndPoint string  `yaml:"endpoint"`
}
type TransferContent struct {
	Seed string `yaml:"seed"`
	Address string `yaml:"address"`
	Confirmations int64 `yaml:"confirmations"`
	Gas_limit int64 `yaml:"Gas_limit"`
	Gas_price int64 `yaml:"Gas_price"`
}
type LogContent struct {
	Label  string `yaml:"label"`
	Disable_sentry bool `yaml:"disable_sentry"`
}

func (c *Target) getConf() *Target {

	yamlFile, err := ioutil.ReadFile("compose.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
func MakeTransaction(){
	var MyTarget Target;
	MyTarget.getConf()

	fmt.Println(MyTarget)
	/*MyTarget:=Target{}
	MyLink:=MyTarget.Address

	client, err := ethclient.Dial("MyLink") //connect to a client

	if err != nil {
		log.Fatal(err)
	}


	privateKey:= Target.key //getting private key

	publicKey := privateKey.Public() //public address of of the account we're sending from

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // установить тип нашей переменной publicKey и присвоить ее publicKeyECDSA
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	from:= crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), from) //getting nonce
	if err != nil {
		log.Fatal(err)
	}


	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	to:=MyTarget.EndPoint
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil) //generate unsigned transaction

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}*/


	fmt.Println("Transaction was made successfully")

}

func main() {
	d := time.NewTicker(3 * time.Second)
	MyChannel := make(chan bool)

	go func() {
		time.Sleep(9 * time.Second)
		MyChannel <- true
	}()

	for {
		select {
		case <-MyChannel:
			fmt.Println("Completed!")
			return
		case tm := <-d.C:
			MakeTransaction()
			fmt.Println("The Current time is: ", tm)
		}
	}
}