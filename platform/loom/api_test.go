package loom

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
)

const validatorSrc = `[
      {
        "address": "0x67CD084888F817ae1F2e4f9b783c6Ae38883Fd51",
        "name": "loomX",
        "description": "test-trustwallet-endpoint",
		"delegationTotal": "1107303053327166214720534",
		"fee": "1000"
      },
      {
        "address": "0xeD8cc9d288d684A51D671cC85E80F0093BcfFCA1",
		"delegationTotal": "1650000152682656494684657",
		"fee": "5000"
      },
      {
        "address": "0x4b7bf93583Ba525f681cD394115B8E75f91fCD03",
        "name": "numero_dos",
        "delegationTotal": "1450000173769670699152977"
      }
    ]
`

func TestGetValidatorAPI(t *testing.T) {
	validators := make([]Validator, 0)
	err := json.Unmarshal([]byte(validatorSrc), &validators)
	assert.NoError(t, err)
	actual, err := NormalizeValidators(validators)
	assert.NoError(t, err)
	for i, v := range validators {
		fee, _ := strconv.ParseFloat(v.Fee, 64)
		expected := blockatlas.Validator{
			Status: !v.Jailed,
			ID:     v.Address,
			Reward: blockatlas.StakingReward{Annual: calAnnualRate(fee / 100)},
		}
		fmt.Println(calAnnualRate(fee / 100))
		assert.Equal(t, expected, actual[i])
	}
}

func TestCalculateAnnualReward(t *testing.T) {

	tests := []struct {
		fee                string
		ExpectedAnnualRate float64
	}{
		{"100", 19.8},
		{"1000", 18},
		{"3000", 14},
		{"10000", 0},
	}

	for _, test := range tests {
		actual, err := feeStringToFloat(test.fee)
		assert.NoError(t, err)
		actual = calAnnualRate(actual)
		assert.Equal(t, test.ExpectedAnnualRate, actual)
	}
}
