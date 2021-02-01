// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"errors"
	"fmt"

	"github.com/dn3010/go-cennznet-rpc-client/v2/scale"
	"github.com/dn3010/go-cennznet-rpc-client/v2/signature"
)

// ExtrinsicPayloadV3 is a signing payload for an Extrinsic. For the final encoding, it is variable length based on
// the contents included. Note that `BytesBare` is absolutely critical – we don't want the method (Bytes)
// to have the length prefix included. This means that the data-as-signed is un-decodable,
// but is also doesn't need the extra information, only the pure data (and is not decoded)
// ... The same applies to V1 & V1, if we have a V4, carry move this comment to latest
type ExtrinsicPayloadV1 struct {
	Method BytesBare

	Era                ExtrinsicEra
	Nonce              UCompact
	TransactionPayment TransactionPayment

	SpecVersion        U32
	TransactionVersion U32
	GenesisHash        Hash
	BlockHash          Hash
}

type TransactionPayment struct {
	Tip         UCompact
	FeeExchange OptionFeeExchange
}

type OptionFeeExchange struct {
	HasValue    bool
	FeeExchange FeeExchange
}

type FeeExchange struct {
	IsFeeExchangeV1 bool
	AsFeeExchangeV1 FeeExchangeV1
}

type FeeExchangeV1 struct {
	AssetID    UCompact
	MaxPayment UCompact
}

// Sign the extrinsic payload with the given derivation path
func (e ExtrinsicPayloadV1) Sign(signer signature.KeyringPair) (Signature, error) {
	b, err := EncodeToBytes(e)
	if err != nil {
		return Signature{}, err
	}

	sig, err := signature.Sign(b, signer.URI)
	return NewSignature(sig), err
}

// Decode does nothing and always returns an error. ExtrinsicPayloadV1 is only used for encoding, not for decoding
func (e *ExtrinsicPayloadV1) Decode(decoder scale.Decoder) error {
	return fmt.Errorf("decoding of ExtrinsicPayloadV1 is not supported")
}

func (fe OptionFeeExchange) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(fe.HasValue, fe.FeeExchange)
}

func (fe *OptionFeeExchange) Decode(decoder scale.Decoder) error {
	flag, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	fe.HasValue = flag == 0x01
	if fe.HasValue {
		return fe.FeeExchange.Decode(decoder)
	}
	return nil
}

func (fe FeeExchange) Encode(encoder scale.Encoder) error {
	switch {
	case fe.IsFeeExchangeV1:
		if err := encoder.PushByte(0); err != nil {
			return err
		}
		return encoder.Encode(fe.AsFeeExchangeV1)
	}

	return errors.New("Only FeeExchangeV1 is supported")
}

func (fe *FeeExchange) Decode(decoder scale.Decoder) error {
	flag, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch flag {
	case 0x00:
		fe.IsFeeExchangeV1 = true
		return fe.AsFeeExchangeV1.Decode(decoder)
	}
	return nil
}

func (fe *FeeExchangeV1) Decode(decoder scale.Decoder) error {
	if err := fe.AssetID.Decode(decoder); err != nil {
		return err
	}
	return fe.MaxPayment.Decode(decoder)
}
