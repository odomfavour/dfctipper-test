package app

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ademuanthony/dfctipper/dfc"
	"github.com/ademuanthony/dfctipper/pancake/pair"
	"github.com/ademuanthony/dfctipper/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

var (
	dfcContractAddress = common.HexToAddress("0x996c1bf72Ec220289ae0edd3a8d77080642121a2")
	tokenTransferFee   = "0.0001912"
)

func (m app) checkDfcBalance(ctx context.Context, address string) (*big.Int, error) {
	dfcToken, err := dfc.NewDfc(dfcContractAddress, m.client)
	if err != nil {
		return nil, err
	}

	dfcAddress := common.HexToAddress(address)
	return dfcToken.BalanceOf(nil, dfcAddress)
}

func (m app) dfcPrice(ctx context.Context) (*big.Int, error) {
	address := common.HexToAddress("0xBbba7668E7E36752F3eDfc0fF794FdDA090B7560")

	instance, err := pair.NewPancake(address, m.client)
	if err != nil {
		return nil, err
	}

	r, err := instance.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	return r.Reserve0.Div(r.Reserve1.Mul(r.Reserve1, big.NewInt(1e8)), r.Reserve0), nil
}

func (m app) convertDfcBnb(ctx context.Context, amount *big.Int) (*big.Int, error) {
	dfcPrice, err := m.dfcPrice(ctx)
	if err != nil {
		return nil, err
	}

	bq := dfcPrice.Mul(dfcPrice, amount)
	totalBnb := bq.Div(bq, big.NewInt(1e8))

	return totalBnb, nil
}

func (m app) convertDfcBusd(ctx context.Context, amount *big.Int) (*big.Int, error) {
	// first converting dfc to bnb and then bnb to busd
	totalBnb, err := m.convertDfcBnb(ctx, amount)
	if err != nil {
		return nil, err
	}

	return m.convertBnbBusd(ctx, totalBnb)
}

func (m app) convertBusdDfc(ctx context.Context, amount *big.Int) (*big.Int, error) {
	price, err := m.dfcPrice(ctx)
	if err != nil {
		return nil, err
	}
	amountFloat := big.NewFloat(0)
	amountFloat, ok := amountFloat.SetString(amount.String())
	if !ok {
		return nil, errors.New("too bad, not ok")
	}

	priceFloat := big.NewFloat(0)
	priceFloat, ok = priceFloat.SetString(price.String())
	if !ok {
		return nil, errors.New("too bad, not ok")
	}

	bigFloat := amountFloat.Quo(amountFloat, priceFloat)
	if err != nil {
		return nil, err
	}
	return dfcToDecimal(bigFloat), nil
}

func dfcToDecimal(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.8f", eth), ".")[1]
	fracStr += strings.Repeat("0", 8-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func (m app) sendTokenTransferFee(ctx context.Context, address string) error {
	feeFloat, err := ParseBigFloat(tokenTransferFee)
	if err != nil {
		log.Error(err)
		return errors.New("Error in processing payment. Please try again later or contact the admin for help")
	}

	amount := etherToWei(feeFloat)
	_, err = m.transfer(ctx, m.config.MasterAddressKey, string(address), amount)
	return err
}

func (m app) transferDfc(ctx context.Context, privateKeyStr, to string, value *big.Int) (string, error) {
	if !util.IsValidAddress(to) {
		return "", errors.New("invalid address")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Error(privateKeyStr)
		return "", fmt.Errorf("crypto.HexToECDSA %v", err)
	}

	dfcToken, err := dfc.NewDfc(dfcContractAddress, m.client)
	if err != nil {
		return "", fmt.Errorf("dfc.NewDfc %v", err)
	}

	toAddress := common.HexToAddress(to)

	tx, err := dfcToken.Transfer(bind.NewKeyedTransactor(privateKey), toAddress, value)
	if err != nil {
		return "", fmt.Errorf("dfcToken.Transfer %v", err)
	}

	return tx.Hash().Hex(), nil
}

func (m app) convertDollarToDfc(ctx context.Context, amount int64) (*big.Int, error) {
	_amountFloat := float64(amount)
	bigFloat, err := ParseBigFloat(fmt.Sprintf("%.18f", _amountFloat))
	if err != nil {
		return nil, err
	}
	busd := etherToWei(bigFloat)
	return m.convertBusdDfc(ctx, busd)
}
