package config

import (
	"strings"
	"time"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type (
	Config struct {
		Server   *ServerCfg   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2Cfg   `mapstructure:"oauth2" validate:"required"`
		Database *DatabaseCfg `mapstructure:"database" validate:"required"`
	}

	ServerCfg struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
	}
	OAuth2Cfg struct {
		PlayerRedirectUrl string `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientID          string `mapstructure:"clientID" validate:"required"`
		ClientSecret      string `mapstructure:"clientSecret" validate:"required"`
		Endpoints         struct {
			AuthUrl       string `mapstructure:"authUrl" validate:"required"`
			TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
			DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
		} `mapstructure:"endpoints" validate:"required"`
		Scopes      []string `mapstructure:"scopes" validate:"required"`
		UserInfoUrl string   `mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl   string   `mapstructure:"revokeUrl" validate:"required"`
	}
	DatabaseCfg struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // server.port -> SERVER_PORT

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validate := validator.New()
		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}

		setGoogleOAuth2Config(configInstance.OAuth2)
	})
	return configInstance
}

func setGoogleOAuth2Config(oauth2Conf *OAuth2Cfg) {
	PlayerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	AdminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}