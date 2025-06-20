// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"context"

	"github.com/3thix/solana-go"
)

// GetTokenLargestAccounts returns the 20 largest accounts of a particular SPL Token type.
func (cl *Client) GetTokenLargestAccounts(
	ctx context.Context,
	tokenMint solana.PublicKey, // Pubkey of token Mint to query
	commitment CommitmentType, // optional
) (out *GetTokenLargestAccountsResult, err error) {
	params := []interface{}{tokenMint}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getTokenLargestAccounts", params)
	return
}

type GetTokenLargestAccountsResult struct {
	RPCContext
	Value []*TokenLargestAccountsResult `json:"value"`
}
type TokenLargestAccountsResult struct {
	Address solana.PublicKey `json:"address"` // the address of the token account
	UiTokenAmount
}
