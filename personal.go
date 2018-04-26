package main

import (
	"fmt"
)

func PersonalImportRawKey(keyData, passphrase string) (interface{}, error) {
	resp, err := Call("personal_importRawKey", []interface{}{keyData, passphrase})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func PersonalListAccounts() (interface{}, error) {
	resp, err := Call("personal_listAccounts", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


func PersonalLockAccount(address string) (interface{}, error) {
	resp, err := Call("personal_lockAccount", []interface{}{address})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


func PersonalNewAccount(passphrase string) (string) {
	resp, err := Call("personal_newAccount", []interface{}{passphrase})
	if err != nil {
		return "0x0";
	}
	if resp.Error != nil {
		return "0x0";
	}
	return resp.Result.(string)
}


func PersonalUnlockAccount(address, passphrase string, duration int64) (interface{}, error) {
	resp, err := Call("personal_unlockAccount", []interface{}{address, passphrase, duration})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


func PersonalSendTransaction(tx *TransactionObject, passphrase string) (string, error) {
	resp, err := Call("personal_sendTransaction", []interface{}{tx, passphrase})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}
