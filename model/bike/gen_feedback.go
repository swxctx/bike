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

type Feedback struct {
	Id        int32     `db:"id" json:"id"`
	Uid       int32     `db:"uid" json:"uid"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Deleted   byte      `db:"deleted" json:"deleted"`

	_db *db.DBInfo      `db:"-"`
	_tx *db.Transaction `db:"-"`
}

func NewFeedback() *Feedback {
	obj := NewFeedbackWithoutDB()
	obj.SetDBInfo()
	return obj
}

func NewFeedbackWithoutDB() *Feedback {
	obj := &Feedback{}
	return obj
}

func (obj *Feedback) SetDBInfo() {
	database := GetDB()
	obj._db = &db.DBInfo{
		DB:    database.DB,
		DbMap: database.DbMap,
	}
}

func (obj *Feedback) SetTransaction(tx *db.Transaction) {
	obj._tx = tx
}

func (obj *Feedback) GetSqlExecutor() gorp.SqlExecutor {
	if obj._tx != nil {
		return obj._tx
	}
	return obj._db
}

func (obj *Feedback) PrimaryCacheKey() string {
	return fmt.Sprintf("bike:orm:feedback:id:%v", obj.Id)
}

func (obj *Feedback) SaveToCache() error {
	if RedisClient != nil {
		return rediskits.SetModelToCache(RedisClient, obj.PrimaryCacheKey(), obj, DefaultCacheTTL)
	}
	return nil
}

func (obj *Feedback) DeleteCache() error {
	if RedisClient != nil {
		return rediskits.DeleteCache(RedisClient, obj.PrimaryCacheKey(), 3)
	}
	return nil
}

func (obj *Feedback) Insert() {
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

func (obj *Feedback) Update() {
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

func (obj *Feedback) Delete() {

	obj.Deleted = 1

	obj.Update()

}

func GetFeedbackWhere(cond string, args ...interface{}) []*Feedback {
	objs := []*Feedback{}
	database := GetDB()
	_, err := database.Select(&objs, "SELECT `id`, `uid`, `content`, `created_at`, `updated_at`, `deleted` FROM `feedback` WHERE "+cond, args...)
	if err != nil {
		panic(err)
	}
	for _, obj := range objs {
		obj.SetDBInfo()
	}
	return objs
}

func GetFeedbackCount(cond string, args ...interface{}) int64 {
	database := GetDB()
	cnt, err := database.SelectInt("SELECT count(1) FROM `feedback` WHERE "+cond, args...)
	if err != nil {
		panic(err)
	}
	return cnt
}

func GetFeedbackFirst(cond string, args ...interface{}) *Feedback {
	obj := &Feedback{}
	database := GetDB()
	err := database.SelectOne(obj, "SELECT `id`, `uid`, `content`, `created_at`, `updated_at`, `deleted` FROM `feedback` WHERE "+cond+" LIMIT 1", args...)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		panic(err)
	}
	obj.SetDBInfo()
	return obj
}

func GetFeedbackByField(name string, field interface{}) *Feedback {
	obj := &Feedback{}
	database := GetDB()
	err := database.SelectOne(obj, "SELECT `id`, `uid`, `content`, `created_at`, `updated_at`, `deleted` FROM `feedback` WHERE `"+name+"`=?", field)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		panic(err)
	}
	obj.SetDBInfo()
	return obj
}

func GetFeedbackByFieldWithCondition(name, cond string, field interface{}) *Feedback {
	obj := &Feedback{}
	database := GetDB()
	err := database.SelectOne(obj, "SELECT `id`, `uid`, `content`, `created_at`, `updated_at`, `deleted` FROM `feedback` WHERE `"+name+"`=? "+cond, field)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		panic(err)
	}
	obj.SetDBInfo()
	return obj
}

func GetFeedback(key int32) *Feedback {
	obj := &Feedback{}
	var notFound = true
	var err error
	if RedisClient != nil {
		obj.Id = key
		notFound = rediskits.GetCacheToModel(RedisClient, obj.PrimaryCacheKey(), obj)
	}
	if notFound {
		database := GetDB()
		err = database.SelectOne(obj, "SELECT `id`, `uid`, `content`, `created_at`, `updated_at`, `deleted` FROM `feedback` WHERE `id`=?", key)
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
