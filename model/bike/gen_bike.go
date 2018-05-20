// generated by ginpt
package bike

import (
	"bytes"
	"database/sql"
	"fmt"
	"time"

	"github.com/domego/ginkits/db"
	"github.com/domego/ginkits/redis"
	"github.com/domego/gokits"
	"github.com/domego/gokits/log"
	"github.com/domego/gorp"
)

var _ = time.Now
var _ = fmt.Println
var _ = sql.ErrNoRows
var _ = bytes.NewBuffer
var _ = utils.Int32

type Bike struct {
	Id        int32     `db:"id" json:"id"`
	ById      int32     `db:"by_id" json:"by_id"`
	LongiTude string    `db:"longi_tude" json:"longi_tude"`
	LatiTude  string    `db:"lati_tude" json:"lati_tude"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Deleted   byte      `db:"deleted" json:"deleted"`

	_db *db.DBInfo      `db:"-"`
	_tx *db.Transaction `db:"-"`
}

func NewBike() *Bike {
	obj := NewBikeWithoutDB()
	obj.SetDBInfo()
	return obj
}

func NewBikeWithoutDB() *Bike {
	obj := &Bike{}
	return obj
}

func (obj *Bike) SetDBInfo() {
	database := GetDB()
	obj._db = &db.DBInfo{
		DB:    database.DB,
		DbMap: database.DbMap,
	}
}

func (obj *Bike) SetTransaction(tx *db.Transaction) {
	obj._tx = tx
}

func (obj *Bike) GetSqlExecutor() gorp.SqlExecutor {
	if obj._tx != nil {
		return obj._tx
	}
	return obj._db
}

func (obj *Bike) PrimaryCacheKey() string {
	return fmt.Sprintf("bike:orm:bike:id:%v", obj.Id)
}

func (obj *Bike) SaveToCache() error {
	if RedisClient != nil {
		return rediskits.SetModelToCache(RedisClient, obj.PrimaryCacheKey(), obj, DefaultCacheTTL)
	}
	return nil
}

func (obj *Bike) DeleteCache() error {
	if RedisClient != nil {
		return rediskits.DeleteCache(RedisClient, obj.PrimaryCacheKey(), 3)
	}
	return nil
}

func (obj *Bike) Insert() {
	database := obj.GetSqlExecutor()

	obj.CreatedAt = time.Now()

	obj.UpdatedAt = time.Now()

	err := database.Insert(obj)
	if err != nil {
		panic(err)
	}
	if err = obj.SaveToCache(); err != nil {
		log.Errorf("[CacheKey:%s], %s", obj.PrimaryCacheKey(), err)
	}
}

func (obj *Bike) Update() {
	var err error
	database := obj.GetSqlExecutor()

	obj.UpdatedAt = time.Now()

	_, err = database.Update(obj)
	if err != nil {
		panic(err)
	}
	if err = obj.SaveToCache(); err != nil {
		log.Errorf("[CacheKey:%s], %s", obj.PrimaryCacheKey(), err)
	}
}

func (obj *Bike) Delete() {

	obj.Deleted = 1

	obj.Update()

}

func GetBikeWhere(cond string, args ...interface{}) []*Bike {
	objs := []*Bike{}
	database := GetDB()
	_, err := database.Select(&objs, "SELECT `id`, `by_id`, `longi_tude`, `lati_tude`, `created_at`, `updated_at`, `deleted` FROM `bike` WHERE "+cond, args...)
	if err != nil {
		panic(err)
	}
	for _, obj := range objs {
		obj.SetDBInfo()
	}
	return objs
}

func GetBikeCount(cond string, args ...interface{}) int64 {
	database := GetDB()
	cnt, err := database.SelectInt("SELECT count(1) FROM `bike` WHERE "+cond, args...)
	if err != nil {
		panic(err)
	}
	return cnt
}

func GetBikeFirst(cond string, args ...interface{}) *Bike {
	obj := &Bike{}
	database := GetDB()
	err := database.SelectOne(obj, "SELECT `id`, `by_id`, `longi_tude`, `lati_tude`, `created_at`, `updated_at`, `deleted` FROM `bike` WHERE "+cond+" LIMIT 1", args...)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		panic(err)
	}
	obj.SetDBInfo()
	return obj
}

func GetBikeByField(name string, field interface{}) *Bike {
	obj := &Bike{}
	database := GetDB()
	err := database.SelectOne(obj, "SELECT `id`, `by_id`, `longi_tude`, `lati_tude`, `created_at`, `updated_at`, `deleted` FROM `bike` WHERE `"+name+"`=?", field)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		panic(err)
	}
	obj.SetDBInfo()
	return obj
}

func GetBikeByFieldWithCondition(name, cond string, field interface{}) *Bike {
	obj := &Bike{}
	database := GetDB()
	err := database.SelectOne(obj, "SELECT `id`, `by_id`, `longi_tude`, `lati_tude`, `created_at`, `updated_at`, `deleted` FROM `bike` WHERE `"+name+"`=? "+cond, field)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		panic(err)
	}
	obj.SetDBInfo()
	return obj
}

func GetBike(key int32) *Bike {
	obj := &Bike{}
	var notFound = true
	var err error
	if RedisClient != nil {
		obj.Id = key
		notFound = rediskits.GetCacheToModel(RedisClient, obj.PrimaryCacheKey(), obj)
	}
	if notFound {
		database := GetDB()
		err = database.SelectOne(obj, "SELECT `id`, `by_id`, `longi_tude`, `lati_tude`, `created_at`, `updated_at`, `deleted` FROM `bike` WHERE `id`=?", key)
		if err != nil {
			if err.Error() == sql.ErrNoRows.Error() {
				return nil
			}
			panic(err)
		}
	}
	if notFound && RedisClient != nil {
		if err = obj.SaveToCache(); err != nil {
			log.Errorf("[CacheKey:%s], %s", obj.PrimaryCacheKey(), err)
		}
	}
	obj.SetDBInfo()
	return obj
}
