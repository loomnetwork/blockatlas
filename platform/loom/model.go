package loom

const loomAnnualRate = 20.00

// Fee - also references the "amount" struct
type Fee struct {
	FeeAmount []Amount `json:"amount"`
}

// Amount - the asset & quantity. Always seems to be enclosed in an array/list for some reason.
// Perhaps used for multiple tokens transferred in a single sender/reciever transfer?
type Amount struct {
	Denom    string `json:"denom"`
	Quantity string `json:"amount"`
}

// # Staking

type StakingPool struct {
	NotBondedTokens string `json:"not_bonded_tokens"`
	BondedTokens    string `json:"bonded_tokens"`
}

type ValidatorPage []Validator

type Validator struct {
	Address         string `json:"address,omitempty"`
	Jailed          bool   `json:"jailed,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	Image           string `json:"image,omitempty"`
	Website         string `json:"website,omitempty"`
	DelegationTotal string `json:"delegationTotal,omitempty"`
	Fee             string `json:"fee,omitempty"`
}

type Tx struct {
	Block string `json:"height"`
	Date  string `json:"timestamp"`
	ID    string `json:"txhash"`
	Data  Data   `json:"tx"`
}

type Data struct {
	Contents Contents `json:"value"`
}

type Contents struct {
	Message []Message `json:"msg"`
	Fee     Fee       `json:"fee"`
	Memo    string    `json:"memo"`
}

type Message struct {
	Type  string
	Value interface{}
}

// Block - top object of get las block request
type Block struct {
	Meta BlockMeta `json:"block_meta"`
}

//BlockMeta - "Block" sub object
type BlockMeta struct {
	Header BlockHeader `json:"header"`
}

//BlockHeader - "BlockMeta" sub object, height
type BlockHeader struct {
	Height string `json:"height"`
}
