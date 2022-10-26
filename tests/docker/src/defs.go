package src

const (
	ServerAddr = "go-serv.io"
)

const (
	EnvCertRootCaPemFile  = "CERT_ROOT_CA_PEM_FILE"
	EnvCertServerCertFile = "CERT_SERVER_CERT_FILE"
	EnvCertServerKeyFile  = "CERT_SERVER_KEY_FILE"
	EnvCertClientCertFile = "CERT_CLIENT_CERT_FILE"
	EnvCertClientKeyFile  = "CERT_CLIENT_KEY_FILE"
)

const (
	UnsafePort            = 3030
	TlsAnyPort            = 3031
	TlsTrustedPartiesPort = 3032
	WebProxyPort          = 3033
)

var (
	PskKey = []byte{0xd2, 0xf5, 0xc8, 0x45, 0x35, 0xf5, 0xe3, 0x12}
	ApiKey = []byte("secret")
)
