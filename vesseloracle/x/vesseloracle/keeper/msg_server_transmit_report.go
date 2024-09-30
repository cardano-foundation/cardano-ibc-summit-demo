package keeper

import (
	"context"
	"fmt"

	"vesseloracle/x/vesseloracle/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) TransmitReport(goCtx context.Context, msg *types.MsgTransmitReport) (*types.MsgTransmitReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	report, found := k.GetConsolidatedDataReport(ctx, msg.Imo, msg.Ts)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, fmt.Sprintf("Cannot find a consolidated data report with imo %v and ts %v", msg.Imo, msg.Ts))
	}

	// TODO try to transmit via IBC to target chain aka Cardano
	_ = report
	//k.TransmitConsolidatedDataReportPacketPacket(ctx, types.ConsolidatedDataReportPacketPacketData{}, "sourcePort", "sourceChannel", types.Height(12), 12)

	return &types.MsgTransmitReportResponse{
		Imo: msg.Imo,
		Ts:  msg.Ts,
	}, nil
}
