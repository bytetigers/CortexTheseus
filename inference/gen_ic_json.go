// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package inference

import (
	"encoding/json"
	"errors"

	"github.com/CortexFoundation/CortexTheseus/common/hexutil"
)

// MarshalJSON marshals as JSON.
func (i ICWork) MarshalJSON() ([]byte, error) {
	type ICWork struct {
		Type         InferType     `json:"type" gencodec:"required"`
		Model        string        `json:"model" gencodec:"required"`
		Input        hexutil.Bytes `json:"input" gencodec:"required"`
		CvmVersion   int           `json:"cvm_version"`
		CvmNetworkId int64         `json:"cvm_networkid"`
	}
	var enc ICWork
	enc.Type = i.Type
	enc.Model = i.Model
	enc.Input = i.Input
	enc.CvmVersion = i.CvmVersion
	enc.CvmNetworkId = i.CvmNetworkId
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (i *ICWork) UnmarshalJSON(input []byte) error {
	type ICWork struct {
		Type         *InferType     `json:"type" gencodec:"required"`
		Model        *string        `json:"model" gencodec:"required"`
		Input        *hexutil.Bytes `json:"input" gencodec:"required"`
		CvmVersion   *int           `json:"cvm_version"`
		CvmNetworkId *int64         `json:"cvm_networkid"`
	}
	var dec ICWork
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Type == nil {
		return errors.New("missing required field 'type' for ICWork")
	}
	i.Type = *dec.Type
	if dec.Model == nil {
		return errors.New("missing required field 'model' for ICWork")
	}
	i.Model = *dec.Model
	if dec.Input == nil {
		return errors.New("missing required field 'input' for ICWork")
	}
	i.Input = *dec.Input
	if dec.CvmVersion != nil {
		i.CvmVersion = *dec.CvmVersion
	}
	if dec.CvmNetworkId != nil {
		i.CvmNetworkId = *dec.CvmNetworkId
	}
	return nil
}
