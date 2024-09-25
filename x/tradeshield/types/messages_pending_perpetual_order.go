package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePendingPerpetualOrder = "create_pending_perpetual_order"
	TypeMsgUpdatePendingPerpetualOrder = "update_pending_perpetual_order"
	TypeMsgDeletePendingPerpetualOrder = "delete_pending_perpetual_order"
)

var _ sdk.Msg = &MsgCreatePendingPerpetualOrder{}

func NewMsgCreatePendingPerpetualOrder(creator string) *MsgCreatePendingPerpetualOrder {
	return &MsgCreatePendingPerpetualOrder{
		OwnerAddress: creator,
	}
}

func (msg *MsgCreatePendingPerpetualOrder) Route() string {
	return RouterKey
}

func (msg *MsgCreatePendingPerpetualOrder) Type() string {
	return TypeMsgCreatePendingPerpetualOrder
}

func (msg *MsgCreatePendingPerpetualOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePendingPerpetualOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePendingPerpetualOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePendingPerpetualOrder{}

func NewMsgUpdatePendingPerpetualOrder(creator string, id uint64, orderPrice *OrderPrice) *MsgUpdatePendingPerpetualOrder {
	return &MsgUpdatePendingPerpetualOrder{
		OrderId:      id,
		OwnerAddress: creator,
		OrderPrice:   orderPrice,
	}
}

func (msg *MsgUpdatePendingPerpetualOrder) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePendingPerpetualOrder) Type() string {
	return TypeMsgUpdatePendingPerpetualOrder
}

func (msg *MsgUpdatePendingPerpetualOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePendingPerpetualOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePendingPerpetualOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePendingPerpetualOrder{}

func NewMsgDeletePendingPerpetualOrder(creator string, id uint64) *MsgDeletePendingPerpetualOrder {
	return &MsgDeletePendingPerpetualOrder{
		OrderId:      id,
		OwnerAddress: creator,
	}
}
func (msg *MsgDeletePendingPerpetualOrder) Route() string {
	return RouterKey
}

func (msg *MsgDeletePendingPerpetualOrder) Type() string {
	return TypeMsgDeletePendingPerpetualOrder
}

func (msg *MsgDeletePendingPerpetualOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePendingPerpetualOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePendingPerpetualOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}