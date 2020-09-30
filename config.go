package TaxiApi

type ServerConfig struct{
	Port uint
	Url string

}
type DatabaseConfig struct{
	User string
	Password string
	Host string
	DBname string
	Appendix string
}



func (dbconf *DatabaseConfig) GenerateUrl () string{
	return "user=" + dbconf.User + " password=" + dbconf.Password + " host=" + dbconf.Host + " dbname=" + dbconf.DBname + dbconf.Appendix
}


