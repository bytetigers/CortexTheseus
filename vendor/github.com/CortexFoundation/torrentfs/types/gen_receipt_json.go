// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/CortexFoundation/CortexTheseus/common"
	"github.com/CortexFoundation/CortexTheseus/common/hexutil"
)

var _ = (*receiptMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		ContractAddr *common.Address `json:"contractAddress"`
		GasUsed      hexutil.Uint64  `json:"gasUsed" gencodec:"required"`
		Status       hexutil.Uint64  `json:"status"`
	}
	var enc Receipt
	enc.ContractAddr = r.ContractAddr
	enc.GasUsed = hexutil.Uint64(r.GasUsed)
	enc.Status = hexutil.Uint64(r.Status)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		ContractAddr *common.Address `json:"contractAddress"`
		GasUsed      *hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		Status       *hexutil.Uint64 `json:"status"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ContractAddr != nil {
		r.ContractAddr = dec.ContractAddr
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed = uint64(*dec.GasUsed)
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	return nil
}
