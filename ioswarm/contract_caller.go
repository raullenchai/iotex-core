package ioswarm

import (
	"context"
	"encoding/hex"
	"math/big"
	"sync"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-core/v2/action"
	"github.com/iotexproject/iotex-core/v2/actpool"
	"go.uber.org/zap"
)

// ContractCaller signs and submits contract call transactions to the local actpool.
// It uses a dedicated reward hot wallet key, separate from the delegate's operator key.
type ContractCaller struct {
	mu      sync.Mutex
	privKey crypto.PrivateKey
	actPool actpool.ActPool
	nonce   uint64
	logger  *zap.Logger
}

// NewContractCaller creates a ContractCaller from a hex-encoded private key.
func NewContractCaller(hexKey string, ap actpool.ActPool, logger *zap.Logger) (*ContractCaller, error) {
	sk, err := crypto.HexStringToPrivateKey(hexKey)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{
		privKey: sk,
		actPool: ap,
		logger:  logger,
	}, nil
}

// Call constructs a contract execution tx, signs it with the hot wallet key,
// and submits it to the local actpool for inclusion in the next block.
//
// Parameters:
//   - contract: target contract address in ioAddr format (io1...)
//   - amount:   IOTX to send with the call (in rau), can be nil for zero-value calls
//   - data:     ABI-encoded calldata
//   - gasLimit: gas limit for the execution
func (cc *ContractCaller) Call(ctx context.Context, contract string, amount *big.Int, data []byte, gasLimit uint64) error {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	if amount == nil {
		amount = big.NewInt(0)
	}

	gasPrice := big.NewInt(1_000_000_000_000) // 1 Qev (standard IoTeX gas price)

	selp, err := action.SignedExecution(
		contract,
		cc.privKey,
		cc.nonce,
		amount,
		gasLimit,
		gasPrice,
		data,
	)
	if err != nil {
		cc.logger.Error("failed to sign reward tx", zap.Error(err))
		return err
	}

	if err := cc.actPool.Add(ctx, selp); err != nil {
		cc.logger.Error("failed to add reward tx to actpool", zap.Error(err))
		return err
	}

	h, _ := selp.Hash()
	cc.logger.Info("reward tx submitted to actpool",
		zap.String("contract", contract),
		zap.String("hash", hex.EncodeToString(h[:])),
		zap.Uint64("nonce", cc.nonce),
		zap.String("amount", amount.String()),
	)

	cc.nonce++
	return nil
}

// SetNonce sets the next nonce to use. Call this on startup with the
// current on-chain nonce of the hot wallet.
func (cc *ContractCaller) SetNonce(n uint64) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.nonce = n
}

// Address returns the ioAddr of the hot wallet.
func (cc *ContractCaller) Address() string {
	return cc.privKey.PublicKey().Address().String()
}
