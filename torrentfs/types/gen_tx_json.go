// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/CortexFoundation/CortexTheseus/common"
	"github.com/CortexFoundation/CortexTheseus/common/hexutil"
)

//var _ = (*transactionMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (t Transaction) MarshalJSON() ([]byte, error) {
	type Transaction struct {
		//Price     *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		Amount   *hexutil.Big   `json:"value"    gencodec:"required"`
		GasLimit hexutil.Uint64 `json:"gas"      gencodec:"required"`
		Payload  hexutil.Bytes  `json:"input"    gencodec:"required"`
		//From      *common.Address `json:"from"     gencodec:"required"`
		Recipient *common.Address `json:"to"       rlp:"nil"`
		Hash      *common.Hash    `json:"hash"     gencodec:"required"`
		//		Receipt   *TxReceipt      `json:"receipt"  rlp:"nil"`
	}
	var enc Transaction
	//enc.Price = (*hexutil.Big)(t.Price)
	enc.Amount = (*hexutil.Big)(t.Amount)
	enc.GasLimit = hexutil.Uint64(t.GasLimit)
	enc.Payload = t.Payload
	//enc.From = t.From
	enc.Recipient = t.Recipient
	enc.Hash = t.Hash
	//	enc.Receipt = t.Receipt
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	type Transaction struct {
		//Price     *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		Amount   *hexutil.Big    `json:"value"    gencodec:"required"`
		GasLimit *hexutil.Uint64 `json:"gas"      gencodec:"required"`
		Payload  *hexutil.Bytes  `json:"input"    gencodec:"required"`
		//From      *common.Address `json:"from"     gencodec:"required"`
		Recipient *common.Address `json:"to"       rlp:"nil"`
		Hash      *common.Hash    `json:"hash"     gencodec:"required"`
		//		Receipt   *TxReceipt      `json:"receipt"  rlp:"nil"`
	}
	var dec Transaction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	//if dec.Price == nil {
	//	return errors.New("missing required field 'gasPrice' for Transaction")
	//}
	//t.Price = (*big.Int)(dec.Price)

	if dec.Amount == nil {
		return errors.New("missing required field 'value' for Transaction")
	}
	t.Amount = (*big.Int)(dec.Amount)

	if dec.GasLimit == nil {
		return errors.New("missing required field 'gas' for Transaction")
	}
	t.GasLimit = uint64(*dec.GasLimit)

	if dec.Payload == nil {
		return errors.New("missing required field 'input' for Transaction")
	}
	t.Payload = *dec.Payload

	//if dec.From == nil {
	//	return errors.New("missing required field 'from' for Transaction")
	//}
	//t.From = dec.From

	if dec.Recipient != nil {
		t.Recipient = dec.Recipient
	}

	if dec.Hash == nil {
		return errors.New("missing required field 'hash' for Transaction")
	}
	t.Hash = dec.Hash

	//	if dec.Receipt != nil {
	//		t.Receipt = dec.Receipt
	//	}
	return nil
}
