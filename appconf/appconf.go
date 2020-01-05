package appconf

import (
	"fmt"
	"log"
	"os"
	"pkkkk/migration"
	"pkkkk/models"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/xwinata/basemodel"

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
	dbconf := c.Viper.GetStringMap(fmt.Sprintf("%s.database", c.ENV))
	Cons := basemodel.DBConfig{
		Adapter:        basemodel.PostgresAdapter,
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
	basemodel.Start(Cons)
	c.DB = basemodel.DB

	err = c.DB.AutoMigrate(&models.Provinsi{}, &models.Kota{}, &models.Kelurahan{}, &models.Kecamatan{}).Error
	if err != nil {
		return err
	}

	c.DB.Model(&models.Kota{}).AddForeignKey("provinsi_id", "provinsis(id)", "CASCADE", "RESTRICT")
	c.DB.Model(&models.Kecamatan{}).AddForeignKey("kota_id", "kota(id)", "CASCADE", "RESTRICT")
	c.DB.Model(&models.Kelurahan{}).AddForeignKey("kecamatan_id", "kecamatans(id)", "CASCADE", "RESTRICT")

	return nil
}

// Migrate database migration
func (c *Configs) Migrate() (err error) {
	db := c.DB
	err = migration.AutoMigrate(db)

	return err
}
