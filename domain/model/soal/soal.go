package soal

type GetFractionsMoney struct {
	InputMoney int `json:"money"`
}

type ResponseMoney struct {
	Result []interface{} `json:"result"`
}

type GetCompareTwoString struct {
	Kalimat1 string `json:"kalimat_1"`
	Kalimat2 string `json:"kalimat_2"`
}

type ResponseCompareString struct {
	Result bool `json:"result"`
}
