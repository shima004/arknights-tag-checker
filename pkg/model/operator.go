package model

type Operator struct {
	Name               string              `json:"name"`
	Rarity             int                 `json:"rarity"`
	Tags               []Tag               `json:"tags"`
	AcquisitionMethods []AcquisitionMethod `json:"acquisition_methods"`
}
