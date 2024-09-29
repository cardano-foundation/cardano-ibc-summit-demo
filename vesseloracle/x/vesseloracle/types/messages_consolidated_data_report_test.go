package types

import (
	"testing"

	"vesseloracle/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateConsolidatedDataReport_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateConsolidatedDataReport
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateConsolidatedDataReport{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateConsolidatedDataReport{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateConsolidatedDataReport_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateConsolidatedDataReport
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateConsolidatedDataReport{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateConsolidatedDataReport{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteConsolidatedDataReport_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteConsolidatedDataReport
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteConsolidatedDataReport{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteConsolidatedDataReport{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
