// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package engine

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*blockMetadataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b BlockMetadata) MarshalJSON() ([]byte, error) {
	type BlockMetadata struct {
		Beneficiary    common.Address `json:"beneficiary"     gencodec:"required"`
		GasLimit       uint64         `json:"gasLimit"     gencodec:"required"`
		Timestamp      hexutil.Uint64 `json:"timestamp"     gencodec:"required"`
		MixHash        common.Hash    `json:"mixHash"     gencodec:"required"`
		ExtraData      hexutil.Bytes  `json:"extraData"     gencodec:"required"`
		TxList         hexutil.Bytes  `json:"txList"     gencodec:"required"`
		HighestBlockID *big.Int       `json:"highestBlockID"     gencodec:"required"`
		GasUsedLimit   uint64         `json:"gasUsedLimit"     gencodec:"required"`
	}
	var enc BlockMetadata
	enc.Beneficiary = b.Beneficiary
	enc.GasLimit = b.GasLimit
	enc.Timestamp = hexutil.Uint64(b.Timestamp)
	enc.MixHash = b.MixHash
	enc.ExtraData = b.ExtraData
	enc.TxList = b.TxList
	enc.HighestBlockID = b.HighestBlockID
	enc.GasUsedLimit = b.GasUsedLimit
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *BlockMetadata) UnmarshalJSON(input []byte) error {
	type BlockMetadata struct {
		Beneficiary    *common.Address `json:"beneficiary"     gencodec:"required"`
		GasLimit       *uint64         `json:"gasLimit"     gencodec:"required"`
		Timestamp      *hexutil.Uint64 `json:"timestamp"     gencodec:"required"`
		MixHash        *common.Hash    `json:"mixHash"     gencodec:"required"`
		ExtraData      *hexutil.Bytes  `json:"extraData"     gencodec:"required"`
		TxList         *hexutil.Bytes  `json:"txList"     gencodec:"required"`
		HighestBlockID *big.Int        `json:"highestBlockID"     gencodec:"required"`
		GasUsedLimit   *uint64         `json:"gasUsedLimit"     gencodec:"required"`
	}
	var dec BlockMetadata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Beneficiary == nil {
		return errors.New("missing required field 'beneficiary' for BlockMetadata")
	}
	b.Beneficiary = *dec.Beneficiary
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for BlockMetadata")
	}
	b.GasLimit = *dec.GasLimit
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for BlockMetadata")
	}
	b.Timestamp = uint64(*dec.Timestamp)
	if dec.MixHash == nil {
		return errors.New("missing required field 'mixHash' for BlockMetadata")
	}
	b.MixHash = *dec.MixHash
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extraData' for BlockMetadata")
	}
	b.ExtraData = *dec.ExtraData
	if dec.TxList == nil {
		return errors.New("missing required field 'txList' for BlockMetadata")
	}
	b.TxList = *dec.TxList
	if dec.HighestBlockID == nil {
		return errors.New("missing required field 'highestBlockID' for BlockMetadata")
	}
	b.HighestBlockID = dec.HighestBlockID
	if dec.GasUsedLimit == nil {
		return errors.New("missing required field 'gasUsedLimit' for BlockMetadata")
	}
	b.GasUsedLimit = *dec.GasUsedLimit
	return nil
}
