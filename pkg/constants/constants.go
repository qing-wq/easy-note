package constants

const (
	MySQLDefaultDSN         = "root:root@tcp(localhost:3306)/easy_note?charset=utf8&parseTime=True&loc=Local"
	NoteTableName           = "note"
	UserTableName           = "user"
	EtcdAddress             = "127.0.0.1:2379"
	NoteServiceAddr         = "127.0.0.1:8888"
	UserServiceAddr         = "127.0.0.1:8889"
	ApiServiceAddr          = "127.0.0.1:8080"
	UserServiceName         = "demouser"
	NoteServiceName         = "demonote"
	ApiServiceName          = "demoapi"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	SecretKey               = "secrete key"
)
