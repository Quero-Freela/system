package database

import (
	"github.com/Quero-Freela/system/server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
)

var models = make(map[string]interface{})

func Connect() *gorm.DB {
	dsn, e := utils.GetSecret[string](utils.DatabaseDsn)

	if e != nil {
		panic(e)
	}

	db, err := gorm.Open(postgres.Open(*dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func RegisterModels(mods ...interface{}) {
	reflected := reflect.ValueOf(&utils.BaseModel{}).Elem().Type().Name()
	for _, model := range mods {
		var mod reflect.Type
		if reflect.TypeOf(model).Kind() != reflect.Ptr {
			if reflect.TypeOf(model).Kind() != reflect.Struct {
				panic("model must be a pointer to struct")
			}

			mod = reflect.TypeOf(model)
		} else {
			if reflect.TypeOf(model).Elem().Kind() != reflect.Struct {
				panic("model must be a pointer to struct")
			}

			mod = reflect.TypeOf(model).Elem()
		}

		inherits := mod.NumField() > 0 && mod.Field(0).Type.Name() == reflected
		if !inherits {
			panic("model must inherit from BaseModel")
		}

		name := mod.Name()
		models[name] = model
	}
}

func Migrate(db *gorm.DB, models ...interface{}) error {
	RegisterModels(models...)
	return db.AutoMigrate(models...)
}

func MigrateAll(db *gorm.DB) error {
	mods := make([]interface{}, 0, len(models))
	for _, model := range models {
		mods = append(mods, model)
	}

	return Migrate(db, mods...)
}
