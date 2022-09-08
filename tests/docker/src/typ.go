package src

const (
	ServerAddr = "localhost"
)

const (
	EnvCertRootCaPemFile  = "CERT_ROOT_CA_PEM_FILE"
	EnvCertServerPemFile  = "CERT_SERVER_PEM_FILE"
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
