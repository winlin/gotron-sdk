package account

import (
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

// FrozenResource by account
type FrozenResource struct {
	Type       core.ResourceCode
	Amount     int64
	DelegateTo string
	Expire     int64
}

// UnfrozenResource by account
type UnfrozenResource struct {
	Type   core.ResourceCode
	Amount int64
	Expire int64
}

// Account detailed view
type Account struct {
	Address                 string             `json:"address"`
	Type                    string             `json:"type"`
	Name                    string             `json:"name"`
	ID                      string             `json:"id"`
	Balance                 int64              `json:"balance"`
	Allowance               int64              `json:"allowance"`
	LastWithdraw            int64              `json:"lastWithdraw"`
	IsWitness               bool               `json:"isWitness"`
	IsElected               bool               `json:"isElected"`
	Assets                  map[string]int64   `json:"assetList"`
	TronPower               int64              `json:"tronPower"`
	TronPowerUsed           int64              `json:"tronPowerUsed"`
	FrozenBalance           int64              `json:"frozenBalance"`
	FrozenResources         []FrozenResource   `json:"frozenList"`
	FrozenBalanceV2         int64              `json:"frozenBalanceV2"`
	FrozenResourcesV2       []FrozenResource   `json:"frozenListV2"`
	UnfrozenResource        []UnfrozenResource `json:"unfrozenList"`
	Votes                   map[string]int64   `json:"voteList"`
	BWTotal                 int64              `json:"bandwidthTotal"`
	BWUsed                  int64              `json:"bandwidthUsed"`
	EnergyTotal             int64              `json:"energyTotal"`
	EnergyUsed              int64              `json:"energyUsed"`
	Rewards                 int64              `json:"rewards"`
	WithdrawableBalance     int64              `json:"withdrawableBalance"`
	UnfreezeLeft            int64              `json:"countUnfreezeLeft"`
	MaxCanDelegateBandwidth int64              `json:"maxCanDelegateBandwidth"`
	MaxCanDelegateEnergy    int64              `json:"maxCanDelegateEnergy"`
	TotalEnergyWeight       int64              `json:"totalEnergyWeight"`
	TotalEnergyLimit        int64              `json:"totalEnergyLimit"`
	TotalNetWeight          int64              `json:"totalNetWeight"`
	TotalNetLimit           int64              `json:"totalNetLimit"`
	OwnerPermission         *core.Permission   `protobuf:"bytes,31,opt,name=owner_permission,json=ownerPermission,proto3" json:"owner_permission,omitempty"`
	WitnessPermission       *core.Permission   `protobuf:"bytes,32,opt,name=witness_permission,json=witnessPermission,proto3" json:"witness_permission,omitempty"`
	ActivePermission        []*core.Permission `protobuf:"bytes,33,rep,name=active_permission,json=activePermission,proto3" json:"active_permission,omitempty"`
}
