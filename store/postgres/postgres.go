package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ory/viper"
	"github.com/pkg/errors"
)

// Config for SQL database connection
type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// NewConfig construct DB config
func NewConfig() *Config {
	viper.AutomaticEnv()

	// Binding configs with environment variables
	viper.BindEnv("dbHost", "DB_HOST")
	viper.BindEnv("dbPort", "DB_PORT")
	viper.BindEnv("dbUser", "DB_USERNAME")
	viper.BindEnv("dbPassword", "DB_PASSWORD")
	viper.BindEnv("dbName", "DB_NAME")

	return &Config{
		Host:     viper.GetString("dbHost"),
		Port:     5444,
		Username: viper.GetString("dbUser"),
		Password: viper.GetString("dbPassword"),
		Database: viper.GetString("dbName"),
	}
}

// Client a postgres database client
type Client struct {
	db     *gorm.DB
	models []interface{}
}

// NewClient create a new client connecting to a PostgreSQL database
func NewClient(c *Config, models []interface{}) (*Client, error) {
	connectionParameters := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s timezone=Europe/Paris sslmode=disable",
		c.Host,
		c.Port,
		c.Database,
		c.Username,
		c.Password,
	)
	db, err := gorm.Open("postgres", connectionParameters)
	if err != nil {
		return nil, errors.Wrap(err, "gorm Open")
	}
	db.Debug()
	db.LogMode(true)
	if viper.GetBool("sql_trace") {
		db = db.Debug()
	}
	err = db.AutoMigrate(models...).Error

	if err != nil {
		return nil, errors.Wrap(err, "gorm AutoMigrate")
	}

	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(30)

	return &Client{
		db:     db,
		models: models,
	}, nil
}

// GracefulShutdown close client connection to database
func (c *Client) GracefulShutdown() error {
	return errors.Wrap(c.db.Close(), "Close")
}
