package model

type PVZWithReceptions struct {
	PVZ        PVZ
	Receptions []ReceptionWithProducts
}

type ReceptionWithProducts struct {
	Reception Reception
	Products  []Product
}
