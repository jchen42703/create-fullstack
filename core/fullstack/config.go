package fullstack

type PaymentsOptions struct {
	UseStripe bool `yaml:"stripe"`
	UsePayPal bool `yaml:"paypal"`
}

type OpenIDConnectOptions struct {
	GoogleCallback   *string `yaml:"google_callback_url"`
	FacebookCallback *string `yaml:"facebook_callback_url"`
	GitHubCallback   *string `yaml:"github_callback_url"`
}

type UsernamePasswordOptions struct {
	EmailVerification bool `yaml:"email_verification"`
	UsernameIsEmail   bool `yaml:"username_is_email"`
}

type AuthOptions struct {
	OpenIDConnectOpts    *OpenIDConnectOptions    `yaml:"social_sign_in"`
	UsernamePasswordOpts *UsernamePasswordOptions `yaml:"username_password"`
}

type TemplateConfig struct {
	OutputDirectoryPath string          `yaml:"output_dir"`
	AuthOpts            AuthOptions     `yaml:"auth"`
	PaymentsOpts        PaymentsOptions `yaml:"payments"`
}
