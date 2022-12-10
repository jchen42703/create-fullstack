package configs

type OpenIDConnectOptions struct {
	GoogleCallback   *string
	FacebookCallback *string
	GitHubCallback   *string
}

// TODO: Make more flexibile down the line.
type PasswordValidationOptions struct {
	Validate bool
}

type UsernamePasswordOptions struct {
	ConfirmEmailWithEmail  bool
	UsernameIsEmail        bool
	PasswordValidationOpts PasswordValidationOptions
}

type AuthOptions struct {
	OpenIDConnectOpts    *OpenIDConnectOptions
	UsernamePasswordOpts UsernamePasswordOptions
}

type FullstackConfig struct {
	OutputDirectoryName string
	FrontendOpts        FrontendConfig
	BackendOpts         BackendConfig
	AuthOpts            AuthOptions
}
