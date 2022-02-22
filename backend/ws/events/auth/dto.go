package auth

type LoginDTO struct {
	Token string `mapstructure:"token" validation:"required,jwt"`
}
