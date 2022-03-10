package db

import (
	"query_rpc/conf"
	"query_rpc/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DBClientImpl struct {
	DBConn *gorm.DB
}

func NewDBClientImpl(driver, source string) (*DBClientImpl, error) {
	db := &DBClientImpl{}
	conn, err := gorm.Open(driver, source)
	// debug模式
	//conn.LogMode(true)
	// 全局禁用复数表名
	conn.SingularTable(true)
	conn.DB().SetMaxOpenConns(conf.Max_Conn)
	conn.DB().SetMaxIdleConns(conf.Max_Idle_Conn)
	conn.DB().SetConnMaxIdleTime(conf.Max_Idle_Time)
	if err != nil {
		return nil, err
	}
	db.DBConn = conn
	return db, nil
}

func (d *DBClientImpl) CreateUser(user *models.User) error {
	err := d.DBConn.Create(user).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] CreateUser Create err:", err)
		return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateUser err:")
	}
	return nil
}

func (d *DBClientImpl) GetUserByUserID(id int32) (*models.User, error) {
	user := &models.User{}
	// 查询用户信息
	err := d.DBConn.Where("user_id=?", id).First(user).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetUserByUserID Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetUserByUserID err:")
	}

	return user, nil
}

func (d *DBClientImpl) GetUserSpace(id int32) ([]models.Space, error) {
	// 查询用户工作空间、收藏夹
	spaces := []models.Space{}
	err := d.DBConn.Where("user_id=?", id).Find(&spaces).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetUserByUserID Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetUserByUserID err:")
	}
	return spaces, nil
}

// TODO:事务保证原子性
func (d *DBClientImpl) CreateQuery(query *models.Query, id int32) error {
	err := d.DBConn.Transaction(func(tx *gorm.DB) error {
		// 插入query
		err := tx.Create(query).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] CreateQuery Query err:", err)
			return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateQuery Query err:")
		}
		// 增加用户工作空间
		space := &models.Space{
			User_id:  id,
			Query_id: query.Query_ID,
			Type:     conf.User_WorkSpace_Code,
		}
		err = tx.Create(space).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] CreateQuery Space err:", err)
			return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateQuery Space err:")
		}
		return nil
	})
	return err
}

func (d *DBClientImpl) GetQueryByID(id int32) (*models.Query, error) {
	query := &models.Query{}
	// 查询用户信息
	err := d.DBConn.Where("query_id=?", id).First(query).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetUserByUserID Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetUserByUserID err:")
	}

	return query, nil
}
