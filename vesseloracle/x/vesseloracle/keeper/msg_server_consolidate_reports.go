package keeper

import (
	"context"
	"fmt"

	"vesseloracle/x/vesseloracle/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) consolidateDeparturePort(vesselData []types.Vessel) string {
	return ""
}

func (k msgServer) consolidateEta(vesselData []types.Vessel) uint64 {
	return 0
}

func (k msgServer) ConsolidateVesselData(ctx sdk.Context, imo string) (*types.ConsolidatedDataReport, error) {
	ctx.Logger().Info("Calling ConsolidateVesselData")
	vesselData := k.GetVesselsInWindow(ctx, imo, k.GetConsolidationWindowIntervalWidth(ctx), k.GetConsolidationWindowMaxItemCount(ctx))
	if vesselData != nil && len(vesselData) >= int(k.GetConsolidationWindowMinItemCount(ctx)) {
		var consolidateDataReport = types.ConsolidatedDataReport{
			Imo: imo,
		}
		consolidatedDeparturePort := k.consolidateDeparturePort(vesselData)
		consolidatedEta := k.consolidateEta(vesselData)
		ctx.Logger().Info("Consolidated departure port determined.", "depport", consolidatedDeparturePort)
		ctx.Logger().Info("Consolidated eta determined.", "eta", consolidatedEta)
		return &consolidateDataReport, nil
	}

	return nil, errorsmod.Wrap(sdkerrors.ErrLogic, fmt.Sprint("Unable to consolidate.", vesselData))
}

func (k msgServer) ConsolidateReports(goCtx context.Context, msg *types.MsgConsolidateReports) (*types.MsgConsolidateReportsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	consolidatedReport, err := k.ConsolidateVesselData(ctx, msg.Imo)
	if err != nil {
		return nil, err
	}

	k.SetConsolidatedDataReport(ctx, *consolidatedReport)

	return &types.MsgConsolidateReportsResponse{}, nil
}
