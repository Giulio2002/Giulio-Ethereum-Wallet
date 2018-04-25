package EthImagePwd

import (
	"fmt"
)


func MinerHashrate() (interface{}, error) {
	resp, err := Call("miner_hashrate", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


func MinerMakeDAG(number int64) (interface{}, error) {
	resp, err := Call("miner_makeDAG", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


func MinerSetExtra(extra string) (interface{}, error) {
	resp, err := Call("miner_setExtra", []interface{}{extra})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func MinerSetGasPrice(number int64) (interface{}, error) {
	resp, err := Call("miner_setGasPrice", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func MinerStart(threads int64) (interface{}, error) {
	resp, err := Call("miner_start", []interface{}{threads})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func MinerStartAutoDAG() (interface{}, error) {
	resp, err := Call("miner_startAutoDAG", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func MinerStop() (interface{}, error) {
	resp, err := Call("miner_stop", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func MinerStopAutoDAG() (interface{}, error) {
	resp, err := Call("miner_stopAutoDAG", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
