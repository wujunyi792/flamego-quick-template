package stdao

import (
	"context"
	"gorm.io/gorm"
)

func Create[T any](m T) *Std[T] {
	return &Std[T]{model: m}
}

func (s *Std[T]) WithMigrator(mf func(db *gorm.DB) error) {
	s.migrator = mf
}

type Std[T any] struct {
	// Std is a standard struct for all dao structs
	db    *gorm.DB
	model T

	migrator func(db *gorm.DB) error
}

func (s *Std[T]) Init(db *gorm.DB) (err error) {
	s.db = db
	if s.migrator != nil {
		return s.migrator(s.db)
	}
	return s.db.AutoMigrate(s.model)
}

func (s *Std[T]) Use(plugin gorm.Plugin) (err error) {
	return s.db.Use(plugin)
}

// Session create new db session
func (s *Std[T]) Session(session *gorm.Session) *Std[T] {
	return &Std[T]{db: s.db.Session(session), model: s.model}
}

func (s *Std[T]) WithCtx(ctx context.Context) *Std[T] {
	return &Std[T]{db: s.db.WithContext(ctx), model: s.model}
}

// Debug start debug mode.
func (s *Std[T]) Debug() *Std[T] {
	return &Std[T]{db: s.db.Debug(), model: s.model}
}

// Begin a transaction with any transaction options opts
func (s *Std[T]) Begin() *Std[T] {
	return &Std[T]{db: s.db.Begin(), model: s.model}
}

// Rollback the transaction
func (s *Std[T]) Rollback() (result *gorm.DB) {
	return s.db.Rollback()
}

// Commit the transaction
func (s *Std[T]) Commit() (result *gorm.DB) {
	return s.db.Commit()
}

// Transaction start a transaction as a block, return any error will rollback the transaction
func (s *Std[T]) Transaction(fc func(tx *Std[T]) error) (err error) {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return fc(&Std[T]{db: tx, model: s.model})
	})
}

// Set value with key into current db instance's context.
func (s *Std[T]) Set(key string, value interface{}) *Std[T] {
	return &Std[T]{db: s.db.Set(key, value), model: s.model}
}

// Get value with key from current db instance's context.
func (s *Std[T]) Get(key string) (interface{}, bool) {
	return s.db.Statement.Settings.Load(key)
}

// Create the model to the database.
func (s *Std[T]) Create(model T, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].Create(model)
	}
	return s.db.Create(model)
}

// Save the model to the database.
func (s *Std[T]) Save(model T, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].Save(model)
	}
	return s.db.Save(model)
}

// Update the model to the database.
func (s *Std[T]) Update(model T, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].Updates(model)
	}
	return s.db.Updates(model)
}

func (s *Std[T]) Updates(model T, m map[string]interface{}, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].Model(model).Updates(m)
	}
	return s.db.Model(model).Updates(m)
}

// Delete the model from the database.
func (s *Std[T]) Delete(model T, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].Delete(model)
	}
	return s.db.Delete(model, model)
}

// First the model from the database.
func (s *Std[T]) First(model T, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].First(model)
	}
	return s.db.First(model, model)
}

// Find the model from the database.
func (s *Std[T]) Find(model T, tx ...*gorm.DB) (result *gorm.DB) {
	if len(tx) > 0 {
		return tx[0].Where(model).Find(model)
	}
	return s.db.Find(model, model)
}

func (s *Std[T]) Count() (count int64) {
	s.db.Model(s.model).Count(&count)
	return
}

// DB return the
func (s *Std[T]) DB() *gorm.DB {
	// protect the db in origin dao
	return s.db.Session(&gorm.Session{})
}
