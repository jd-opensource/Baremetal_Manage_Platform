package dao

import (
	"fmt"
	"strings"

	clog "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

var (
	IronicRdb *gorm.DB
	IronicWdb *gorm.DB
)

type gLogger struct {
	traceid string
}

func (g *gLogger) Print(values ...interface{}) {
	var (
		level = values[0]
		// source = values[1]
	)
	logger := clog.New("./log/ironic-console.log")
	defer logger.Flush()
	logger.SetStableFields([]string{"sql", "args", "cost"})
	logger.Point("logid", g.traceid)

	if level == "sql" {
		logger.Point("sql", values[3])
		logger.Point("args", values[4])
		logger.Point("cost", fmt.Sprint(values[2]))
	} else {
		logger.Point("trace", values)
	}
}

func InitGormDb(conn string) error {
	var err error
	IronicRdb, err = gorm.Open("mysql", conn)
	if err != nil {
		return err
	}
	IronicRdb.LogMode(true)
	IronicRdb.DB().SetMaxOpenConns(100)
	IronicRdb.DB().SetMaxIdleConns(20)

	IronicWdb, err = gorm.Open("mysql", conn)
	if err != nil {
		return err
	}
	IronicWdb.LogMode(true)
	IronicWdb.DB().SetMaxOpenConns(100)
	IronicWdb.DB().SetMaxIdleConns(20)

	return nil
}

func GetGormTx(logger *clog.Logger) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := IronicWdb.New()
	db.SetLogger(gl)
	return db
}

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// sql build where
func WhereBuild(g *gorm.DB, query map[string]interface{}) (*gorm.DB, error) {

	for k, v := range query {
		ks := strings.Split(k, ".")
		if len(ks) > 2 {
			fmt.Println(fmt.Errorf("Error in query condition: %s. ", k))
			return g, fmt.Errorf("Error in query condition: %s. ", k)
		}
		switch len(ks) {
		case 1:
			g = g.Where(fmt.Sprintf("%s = ?", k), v)
			break
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				g = g.Where(fmt.Sprintf("%s = ?", k), v)
				break
			case ">", "gt":
				g = g.Where(fmt.Sprintf("%s > ?", k), v)
			case ">=", "gte":
				g = g.Where(fmt.Sprintf("%s >= ?", k), v)
				break
			case "<", "lt":
				g = g.Where(fmt.Sprintf("%s < ?", k), v)
				break
			case "<=", "lte":
				g = g.Where(fmt.Sprintf("%s <= ?", k), v)
				break
			case "!=", "<>":
				g = g.Where(fmt.Sprintf("%s <> ?", k), v)
				break
			case "in":
				g = g.Where(fmt.Sprintf("%s IN (?)", k), v)
				break
			case "like":
				g = g.Where(fmt.Sprintf("%s LIKE ?", k), v)
				break
			}
			break
		}
	}
	return g, nil
}
