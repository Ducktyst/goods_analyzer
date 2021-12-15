package app

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Database struct {
		User   string
		Pass   string
		Host   string
		Port   int
		Dbname string
	}
	//Services   struct {
	//	Bitrix struct {
	//		URL            string
	//		AgentID        string
	//		AgentCode      string
	//		CFClientID     string
	//		CFClientSecret string
	//	}
}

//func (db *DB) connect() {
//	db.DB = pg.Connect(db.conf.DB)
//}
