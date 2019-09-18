package loom

import (
	"strconv"

	"github.com/trustwallet/blockatlas"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = InitClient(viper.GetString("loom.api"), viper.GetString("loom.rpc"))
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
		fee, err := feeStringToFloat(v.Fee)
		if err != nil {
			return nil, err
		}
		results = append(results, blockatlas.Validator{
			Status: !v.Jailed,
			ID:     v.Address,
			Reward: blockatlas.StakingReward{
				Annual:      calculateRate(TierBonusMap[TierMap[3]], fee),
				SixMonths:   calculateRate(TierBonusMap[TierMap[2]], fee),
				ThreeMonths: calculateRate(TierBonusMap[TierMap[1]], fee),
				TwoWeeks:    calculateRate(TierBonusMap[TierMap[0]], fee),
			},
		})
	}
	return results, nil
}

//Fee field from response example (100) is 1.00%
func feeStringToFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	fee, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fee, err
	}
	return fee / 100, nil
}

func calculateRate(rate, fee float64) float64 {
	return rate - (rate * fee / 100)
}
