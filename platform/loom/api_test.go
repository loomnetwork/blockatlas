package loom

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
)

const validatorSrc = `[
      {
        "address": "0x67CD084888F817ae1F2e4f9b783c6Ae38883Fd51",
        "name": "loomX",
        "description": "test-trustwallet-endpoint",
        "delegationTotal": "1107303053327166214720534"
      },
      {
        "address": "0xeD8cc9d288d684A51D671cC85E80F0093BcfFCA1",
        "delegationTotal": "1650000152682656494684657"
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

	actual, err := normalizeValidators(validators)
	assert.NoError(t, err)
	for i, v := range validators {
		expected := blockatlas.Validator{
			Status: !v.Jailed,
			ID:     v.Address,
			Reward: blockatlas.StakingReward{Annual: 20.00},
		}
		assert.Equal(t, actual[i], expected)
	}
}
