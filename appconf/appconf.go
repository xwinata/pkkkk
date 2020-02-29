package appconf

import (
	"fmt"
	"log"
	"os"
	"pkkkk/models"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// import postgres sql
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Configuration instance
var Configuration *Configs

// Configs main type
type Configs struct {
	Root  string      `json:"root"`
	ENV   string      `json:"env"`
	Viper viper.Viper `json:"viper"`
	DB    *gorm.DB    `json:"db"`
}

func init() {
	Configuration := &Configs{}
	Configuration.Root = os.Getenv("ROOT")
	if len(Configuration.Root) < 1 {
		panic("ROOT environtment is empty")
	}
	Configuration.ENV = os.Getenv("APPENV")
	if len(Configuration.ENV) < 1 {
		panic("ENV environtment is empty")
	}
	if err := Configuration.Viperinit(); err != nil {
		panic(fmt.Sprintf("Load viper config error : %v", err))
	}
	if err := Configuration.DBinit(); err != nil {
		panic(fmt.Sprintf("DB init error : %v", err))
	}
}

// Viperinit loads config from config file
func (c *Configs) Viperinit() error {
	var conf *viper.Viper

	conf = viper.New()
	conf.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	conf.AutomaticEnv()
	conf.SetConfigName("config")
	conf.AddConfigPath(fmt.Sprintf("$GOPATH/src/%v/appconf", os.Getenv("APPNAME")))
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		return err
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("App Config file changed %s:", e.Name)
		c.Viperinit()
	})
	c.Viper = viper.Viper(*conf)

	return nil
}

// DBinit initiates database configuration
func (c *Configs) DBinit() (err error) {
	var (
		connectionString string
	)
	type DBConfig struct {
		Adapter        string
		Host           string
		Port           string
		Username       string
		Password       string
		Table          string
		Timezone       string
		Maxlifetime    int
		IdleConnection int
		OpenConnection int
		SSL            string
		Logmode        bool
	}

	dbconf := c.Viper.GetStringMap(fmt.Sprintf("%s.database", c.ENV))
	conf := DBConfig{
		Adapter:        "postgres",
		Host:           dbconf["host"].(string),
		Port:           dbconf["port"].(string),
		Username:       dbconf["username"].(string),
		Password:       dbconf["password"].(string),
		Table:          dbconf["table"].(string),
		Timezone:       dbconf["timezone"].(string),
		Maxlifetime:    dbconf["maxlifetime"].(int),
		IdleConnection: dbconf["idle_conns"].(int),
		OpenConnection: dbconf["open_conns"].(int),
		SSL:            dbconf["sslmode"].(string),
		Logmode:        dbconf["logmode"].(bool),
	}

	switch conf.Adapter {
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.Username, conf.Password, conf.Host, conf.Port, conf.Table)
	case "postgres":
		connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Table, conf.SSL)
	}
	c.DB, err = gorm.Open(conf.Adapter, connectionString)
	if err != nil {
		return err
	}

	c.DB.LogMode(conf.Logmode)

	c.DB.Exec(fmt.Sprintf("SET TIMEZONE TO '%s'", conf.Timezone))
	c.DB.DB().SetConnMaxLifetime(time.Second * time.Duration(conf.Maxlifetime))
	c.DB.DB().SetMaxIdleConns(conf.IdleConnection)
	c.DB.DB().SetMaxOpenConns(conf.OpenConnection)

	// err = c.AutoMigrate()

	return err
}

// AutoMigrate database migration
func (c *Configs) AutoMigrate() (err error) {
	err = c.DB.AutoMigrate(&models.Provinsi{}, &models.Kota{}, &models.Kelurahan{}, &models.Kecamatan{}).Error
	if err != nil {
		return err
	}

	c.DB.Model(&models.Kota{}).AddForeignKey("provinsi_id", "provinsis(id)", "CASCADE", "RESTRICT")
	c.DB.Model(&models.Kecamatan{}).AddForeignKey("kota_id", "kota(id)", "CASCADE", "RESTRICT")
	c.DB.Model(&models.Kelurahan{}).AddForeignKey("kecamatan_id", "kecamatans(id)", "CASCADE", "RESTRICT")

	return err
}
