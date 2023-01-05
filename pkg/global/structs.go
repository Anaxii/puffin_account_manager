package global

type Account struct {
	Name            string `json:"name" bson:"name"`
	ID              string `json:"id" bson:"id"`
	Status          string `json:"status" bson:"status"`
	Email           string `json:"email" bson:"email"`
	WalletAddress   string `json:"wallet_address" bson:"wallet_address"`
	Country         string `json:"country" bson:"country"`
	PhysicalAddress string `json:"physical_address" bson:"physical_address"`
	IdentityNumber  string `json:"identity_number" bson:"identity_number"`
	DateOfBirth     string `json:"date_of_birth" bson:"date_of_birth"`
}

type ClientSettings struct {
	AdminName                   string             `json:"admin_name" bson:"admin_name"`
	ProjectName                 string             `json:"project_name" bson:"project_name"`
	AdminTelegram               string             `json:"admin_telegram" bson:"admin_telegram"`
	TeamSize                    int                `json:"team_size" bson:"team_size"`
	ProjectCommunicationChannel string             `json:"project_communication_channel" bson:"project_communication_channel"`
	ProjectWebsite              string             `json:"project_website" bson:"project_website"`
	DappName                    string             `json:"dapp_name" bson:"dapp_name"`
	RPCURL                      string             `json:"rpc_url" bson:"rpc_url"`
	GasTokenSymbol              string             `json:"gas_token_symbol" bson:"gas_token_symbol"`
	ChainID                     int64              `json:"chain_id" bson:"chain_id"`
	VM                          string             `json:"vm" bson:"vm"`
	Package                     string             `json:"package" bson:"package"`
	PackageOptions              []string           `json:"package_options" bson:"package_options"`
	MaxUsers                    int                `json:"max_users" bson:"max_users"`
	PuffinGeoAddress            string             `json:"puffin_geo_address" bson:"puffin_geo_address"`
	PuffinKYCAddress            string             `json:"puffin_kyc_address" bson:"puffin_kyc_address"`
	PuffinCoreAddress           string             `json:"puffin_core_address" bson:"puffin_core_address"`
	AdminWalletAddress          string             `json:"admin_wallet_address" bson:"admin_wallet_address"`
	UUID                        int                `json:"UUID" bson:"UUID"`
	Status                      string             `json:"status" bson:"status"`
	PaymentExpiration           int                `json:"payment_expiration" bson:"payment_expiration"`
	BlockedCountries            map[int64][]string `json:"blocked_countries" bson:"blocked_countries"`
}

type Config struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string
	Port       string `json:"port"`
	MongoDbURI string `json:"mongo_db_uri"`
}

type ClientUsers struct {
	User    string `json:"user"`
	Client  int    `json:"client"`
	Country string `json:"country"`
	Status  string `json:"status"`
}
