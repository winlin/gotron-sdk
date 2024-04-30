package keys

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	// "github.com/ethereum/go-ethereum/crypto"
)

func checkAndMakeKeyDirIfNeeded() string {
	userDir := common.GetCurrentDir()
	tronCTLDir := path.Join(userDir, ".tronctl", "keystore")
	if _, err := os.Stat(tronCTLDir); os.IsNotExist(err) {
		// Double check with Leo what is right file persmission
		os.Mkdir(tronCTLDir, 0700)
	}

	return tronCTLDir
}

func ListKeys(keystoreDir string) {
	tronCTLDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(tronCTLDir, scryptN, scryptP)
	// keystore.KeyStore
	allAccounts := ks.Accounts()
	fmt.Printf("Tron Address:%s File URL:\n", strings.Repeat(" ", address.AddressLengthBase58))
	for _, account := range allAccounts {
		fmt.Printf("%s\t\t %s\n", account.Address, account.URL)
	}
}

func AddNewKey(password string) {
	tronCTLDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(tronCTLDir, scryptN, scryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Printf("new account error: %v\n", err)
	}
	fmt.Printf("account: %s\n", account.Address)
	fmt.Printf("URL: %s\n", account.URL)
}
