package db

import (
	"errors"
	"mini-wallet-julo/models"
)

// key = token

type Wallet struct {
	walletData map[string]models.Wallet
}

func (w *Wallet) winitWallet() {
	w.walletData = make(map[string]models.Wallet)
}

func (w *Wallet) InitWalletKey(key string) {
	w.walletData[key] = models.Wallet{}
}

func (w *Wallet) DisableWallet(key string) error {
	if e, ok := w.walletData[key]; !ok {
		return errors.New("Token Is Unregistered")
	} else {
		e.IsActive = false
		w.walletData[key] = e
	}
	return nil
}

func (w *Wallet) EnableWallet(key string) error {
	if e, ok := w.walletData[key]; !ok {
		return errors.New("Token Is Unregistered")
	} else {
		e.IsActive = true
		w.walletData[key] = e
	}
	return nil
}

// func (w *Wallet) ()  {

// }

func Init() {
	initWallet()
}
