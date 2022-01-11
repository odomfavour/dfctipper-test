package app

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/ademuanthony/dfctipper/busd"
	"github.com/ademuanthony/dfctipper/util"
	"github.com/ademuanthony/dfctipper/pancake/pair"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

var (
	busdContractAddressStr = "0xe9e7cea3dedca5984780bafc599bd69add087d56"
	busdContractAddress    = common.HexToAddress(busdContractAddressStr)
)

// GenerateWallet creates a new wallet and return the private key and address
func GenerateWallet() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return privateKeyHex, address, nil
}

func (m app) CheckBusdBalance(address string) (*big.Int, error) {
	instance, err := busd.NewPancake(busdContractAddress, m.client)
	if err != nil {
		return nil, err
	}

	return instance.BalanceOf(nil, common.HexToAddress(address))
}

func (m app) transfer(ctx context.Context, privateKeyStr, to string, value *big.Int) (string, error) {
	if !util.IsValidAddress(to) {
		return "", errors.New("invalid address")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := m.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000)
	// gasPrice, err := m.client.SuggestGasPrice(ctx)
	// if err != nil {
	// 	return "", err
	// }

	feeStr := "0.00000001"
	feeFloat, err := ParseBigFloat(feeStr)
	if err != nil {
		return "", err
	}
	gasPrice := etherToWei(feeFloat)

	toAddress := common.HexToAddress(to)
	if toAddress == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		return "", errors.New("invalid address")
	}
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := m.client.NetworkID(ctx)
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = m.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

func (m app) checkBalance(ctx context.Context, address string) (*big.Int, error) {
	return m.client.BalanceAt(ctx, common.HexToAddress(address), nil)
}

func (m app) bnbPrice(ctx context.Context) (*big.Int, error) {
	address := common.HexToAddress("0x1B96B92314C44b159149f7E0303511fB2Fc4774f")

	instance, err := pair.NewPancake(address, m.client)
	if err != nil {
		return nil, err
	}

	r, err := instance.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	return r.Reserve0.Div(r.Reserve1.Mul(r.Reserve1, big.NewInt(1e18)), r.Reserve0), nil
}

func (m app) convertBnbBusd(ctx context.Context, amount *big.Int) (*big.Int, error) {
	price, err := m.bnbPrice(ctx)
	if err != nil {
		return nil, err
	}
	q := price.Mul(price, amount)
	result := q.Div(q, big.NewInt(1e18))
	return result, nil
}

func (m app) convertBusdBnb(ctx context.Context, amount *big.Int) (*big.Int, error) {
	price, err := m.bnbPrice(ctx)
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
	return etherToWei(bigFloat), nil
}

// ParseBigFloat parse string value to big.Float
func ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

func etherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func (m app) convertClubDollarToBnb(ctx context.Context, amount int64) (*big.Int, error) {
	_amountFloat := float64(amount)/float64(10000)
	bigFloat, err := ParseBigFloat(fmt.Sprintf("%.18f", _amountFloat))
	if err != nil {
		return nil, err
	}
	busd := etherToWei(bigFloat)
	return m.convertBusdBnb(ctx, busd)
}

func (m app) transferBusd(ctx context.Context, privateKeyStr, to string, value *big.Int) (string, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := m.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := m.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := value
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := m.client.EstimateGas(ctx, ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		return "", err
	}
	gasLimit *= 3 // protect against out of gas issue

	// if account balance of from address is less that gas fee, credit account
	balance, err := m.client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return "", err
	}
	gasFee := gasPrice.Mul(gasPrice, big.NewInt(int64(gasLimit)))
	if balance.Int64() < gasFee.Int64() {
		_, err := m.transfer(ctx, m.config.MasterAddressKey, fromAddress.Hex(), gasFee.Sub(gasFee, balance))
		if err != nil {
			log.Error("transfer", err)
			return "", err
		}
	}

	chainID, err := m.client.NetworkID(ctx)
	if err != nil {
		return "", err
	}

	// instance, err := busd.NewPancake(busdContractAddress, m.client)
	// if err != nil {
	// 	return "", err
	// }

	// signerFn := func(f common.Address, tx *types.Transaction) (*types.Transaction, error) {
	// 	return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	// }

	// tx, err := instance.Transfer(&bind.TransactOpts{
	// 	From:     fromAddress,
	// 	Nonce:    big.NewInt(int64(nonce)),
	// 	GasPrice: gasPrice,
	// 	GasLimit: gasLimit,
	// 	Context:  ctx,
	// 	Signer:   signerFn,
	// }, common.HexToAddress(to), value)
	// if err != nil {
	// 	log.Error("instance.Transfer", err)
	// 	return "", err
	// }
	// return tx.Hash().Hex(), nil

	tx := types.NewTransaction(nonce, busdContractAddress, common.Big0, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("signedTx", err)
		return "", err
	}

	err = m.client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Error("SendTransaction", err)
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
