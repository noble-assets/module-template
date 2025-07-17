// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, NASD Inc. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package template

import (
	"github.com/spf13/cobra"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	templatev1 "template.dev/api/v1"
)

func (AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: templatev1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "PauseProtocol",
					Use:       "pause-protocol [protocol_id]",
					Short:     "Pause the outgoing bridge protocol by ID. (only authority)",
					Long: `Pause the transfers associated with the outgoing bridge protocol by ID. Finer 
					pausing logic specified targeting counterparties are not overwritten. (only authority)`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "protocol_id"},
					},
				},
				{
					RpcMethod: "PauseCounterparties",
					Use:       "pause-counterparties [protocol_id] [counterparty_ids]",
					Short:     "Pause the outgoing transfers associated with the pairs protocol ID and counterparty IDs. (only authority)",
					Long: `Pause the outgoing transfers associated with the pairs protocol ID and counterparty IDs. The
					ID of the counterparty chain is defined by the specific cross-chain protocol used. (only authority)`,
					Example: "pause-counterparties 1 channel-0",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "protocol_id"},
						{ProtoField: "counterparty_ids"},
					},
				},
				{
					RpcMethod: "UnpauseProtocol",
					Use:       "unpause-protocol [protocol_id]",
					Short:     "Unpause the outgoing bridge protocol by ID. (only authority)",
					Long: `Unpause the transfers associated with the outgoing bridge protocol by ID. Finer 
					pausing logic specified targeting counterparties are not overwritten. (only authority)`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "protocol_id"},
					},
				},
				{
					RpcMethod: "UnpauseCounterparties",
					Use:       "unpause-counterparties [protocol_id] [counterparty_ids]",
					Short:     "Unpause the outgoing transfers associated with the pairs protocol ID and counterparty IDs. (only authority)",
					Long: `Unpause the outgoing transfers associated with the pairs protocol ID and counterparty IDs. The
					ID of the counterparty chain is defined by the specific cross-chain protocol used. (only authority)`,
					Example: "unpause-counterparties 1 channel-0",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "protocol_id"},
						{ProtoField: "counterparty_ids"},
					},
				},
				{
					RpcMethod: "PauseAction",
					Use:       "pause-action [action_id]",
					Short:     "Pause the action associated with the provided action ID. (only authority)",
					Long:      `Pause the action associated with the provided action ID. (only authority)`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "action_id"},
					},
				},
				{
					RpcMethod: "UnpauseAction",
					Use:       "unpause-action [action_id]",
					Short:     "Unpause the action associated with the provided action ID. (only authority)",
					Long:      `Unpause the action associated with the provided action ID. (only authority)`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "action_id"},
					},
				},
			},
			EnhanceCustomCommand: true,
		},
		// Query: &autocliv1.ServiceCommandDescriptor{
		// 	Service: orbiterv1.Query_ServiceDesc.ServiceName,
		// 	RpcCommandOptions: []*autocliv1.RpcCommandOptions{
		// 		{
		// 			RpcMethod: "Protocol",
		// 			Use:       "protocol [protocol_id]",
		// 			Short:     "Query the pause status of a specific protocol",
		// 			Long:      "Query the pause status of a specific protocol by its ID. Returns the
		// protocol information and whether it is currently paused.",
		// 			PositionalArgs: []*autocliv1.PositionalArgDescriptor{
		// 				{ProtoField: "protocol_id"},
		// 			},
		// 		},
		// 		{
		// 			RpcMethod: "Counterparty",
		// 			Use:       "counterparty [protocol_id] [counterparty_id]",
		// 			Short:     "Query the pause status of a specific counterparty for a protocol",
		// 			Long:      "Query the pause status of a specific counterparty for a protocol. Returns
		// the counterparty information and whether it is currently paused.",
		// 			PositionalArgs: []*autocliv1.PositionalArgDescriptor{
		// 				{ProtoField: "protocol_id"},
		// 				{ProtoField: "counterparty_id"},
		// 			},
		// 		},
		// 		{
		// 			RpcMethod: "Action",
		// 			Use:       "action [action_id]",
		// 			Short:     "Query the pause status of a specific action",
		// 			Long:      "Query the pause status of a specific action by its ID. Returns the action
		// information and whether it is currently paused.",
		// 			PositionalArgs: []*autocliv1.PositionalArgDescriptor{
		// 				{ProtoField: "id"},
		// 			},
		// 		},
		// 	},
		// },
	}
}

func (AppModule) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (AppModule) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}
