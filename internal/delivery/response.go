package delivery

type Response struct {
	Solana    Solana            `json:"solanaInfo"`
	Validator ValidatorResponse `json:"validatorInfo"`
	BlockData []Block           `json:"blockData"`
	History   []History         `json:"history"`
}

func (r *Response) Set(response Response) {

}
