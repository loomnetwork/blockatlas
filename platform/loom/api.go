package loom

import (
	"github.com/trustwallet/blockatlas"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = InitClient(viper.GetString("loom.rpc"))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.LOOM]
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	validators, err := p.client.GetValidators()
	if err != nil {
		return nil, err
	}
	return NormalizeValidators(validators)
}

func NormalizeValidators(validators []Validator) (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	for _, v := range validators {
		results = append(results, blockatlas.Validator{
			Status: !v.Jailed,
			ID:     v.Address,
			Reward: blockatlas.StakingReward{Annual: 20.00},
		})
	}
	return results, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// TODO
	return blockatlas.TxPage{}, nil
}
