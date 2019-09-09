package loom

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
)

const validatorSrc = `

       [
         {
      "address": "0x77A71BbD8CAFf532eDe3f82D5518339d863e668e",
      "jailed": false,
      "name": "loomX",
      "description": "test-trustwallet-endpoint",
      "website": "loomX.io",
      "delegationTotal": "739244394776671628451596",
      "fee": "100"
      },
      {
      "address": "0xa47bd25a56798cEEeDB2863F257d3bAfDCEe7c6F",
      "jailed": false,
      "name": "numero_dos",
      "delegationTotal": "1450000362284418808473621",
      "fee": "100"
      },
      {
      "address": "0x23efa4D4957735D1348A7CA1e80B55bF3181dfC6",
      "jailed": false,
      "delegationTotal": "1650000367199432869925070",
      "fee": "100"
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
		expected := blockatlas.Validator{
			Status: v.Jailed,
			ID:     v.Address,
			Reward: blockatlas.StakingReward{Annual: 20.00},
		}
		assert.Equal(t, actual[i], expected)
	}
}
