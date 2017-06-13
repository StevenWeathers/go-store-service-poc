package main

type Response struct {
	StoreLocation []Store `json:"storeLocation"`
}

type Store struct {
	Address1          string   `json:"address1"`
	CanonicalStoreUrl string   `json:"canonicalStoreUrl"`
	City              string   `json:"city"`
	Country           string   `json:"country"`
	DailyHours        Hours    `json:"dailyHours"`
	Fax               string   `json:"fax"`
	Flags             []string `json:"flags"`
	IsOnboardedStore  string   `json:"isOnboardedStore"`
	Latitude          string   `json:"latitude"`
	Longitude         string   `json:"longitude"`
	MapURL            string   `json:"mapUrl"`
	MilesToStore      string   `json:"milesToStore"`
	Phone             string   `json:"phone"`
	State             string   `json:"state"`
	StoreName         string   `json:"storeName"`
	StoreNumber       string   `json:"storeNumber"`
	Zip               string   `json:"zip"`
}

type Hours struct {
	Monday    []int `json:"monday"`
	Tuesday   []int `json:"tuesday"`
	Wednesday []int `json:"wednesday"`
	Thursday  []int `json:"thursday"`
	Friday    []int `json:"friday"`
	Saturday  []int `json:"saturday"`
	Sunday    []int `json:"sunday"`
}
