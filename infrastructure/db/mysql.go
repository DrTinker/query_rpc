package db

import (
	"query_rpc/conf"
	"query_rpc/models"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type DBClientImpl struct {
	DBConn *gorm.DB
}

func NewDBClientImpl(driver, source string) (*DBClientImpl, error) {
	db := &DBClientImpl{}
	conn, _ := gorm.Open(mysql.Open(source), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 指定单数表名
	}})
	// debug模式
	//conn.LogMode(true)
	// 全局禁用复数表名
	sqlDB, err := conn.DB()
	sqlDB.SetMaxOpenConns(conf.Max_Conn)
	sqlDB.SetMaxIdleConns(conf.Max_Idle_Conn)
	sqlDB.SetConnMaxIdleTime(conf.Max_Idle_Time)
	if err != nil {
		return nil, err
	}
	db.DBConn = conn
	return db, nil
}

// 创建用户
func (d *DBClientImpl) CreateUser(user *models.User) error {
	err := d.DBConn.Create(user).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] CreateUser Create err:", err)
		return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateUser err:")
	}
	return nil
}

// 根据user_id查找用户
func (d *DBClientImpl) GetUserByUserID(id int64) (*models.User, error) {
	user := &models.User{}
	// 查询用户信息
	err := d.DBConn.Where("user_id=?", id).First(user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("[DBClientImpl] GetUserByUserID Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetUserByUserID err:")
	}

	return user, nil
}

// 查询用户工作空间、收藏夹
func (d *DBClientImpl) GetUserSpace(id int64) ([]models.Space, error) {
	spaces := []models.Space{}
	err := d.DBConn.Where("user_id=?", id).Find(&spaces).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("[DBClientImpl] GetUserByUserID Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetUserByUserID err:")
	}
	return spaces, nil
}

// 创建问卷，事务保证原子性
func (d *DBClientImpl) CreateQuery(query *models.Query, id int64) error {
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

// 根据query_id查找问卷
func (d *DBClientImpl) GetQueryByID(id int64) (*models.Query, error) {
	query := &models.Query{}
	// 查询用户信息
	err := d.DBConn.Where("query_id=?", id).First(query).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("[DBClientImpl] GetUserByUserID Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetUserByUserID err:")
	}

	return query, nil
}

// 根据user_id获取全部问卷信息
func (d *DBClientImpl) GetQueryBatch(id int64, t int32) ([]*models.Query, error) {
	var querys []*models.Query
	// 查询全部问卷
	err := d.DBConn.Raw("SELECT * FROM query WHERE query_id IN (SELECT query_id FROM space WHERE user_id=? AND type=?)",
		id, t).Scan(&querys).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorln("[DBClientImpl] GetQueryBatch Select err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQueryBatch err:")
	}
	return querys, nil
}

// 增加选择题，事务保证原子性
// 一个question对应一个[]option，这里的[]option为展开后的数组
func (d *DBClientImpl) SetOptionQuestionBatch(questions []models.Question, options []models.Option) error {
	err := d.DBConn.Transaction(func(tx *gorm.DB) error {
		// 采用upsert
		// 在冲突时，更新除主键以外的所有列到新值。
		err := d.DBConn.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&questions).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] CreateOptionQuestion Insert question err: %s", err)
			return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateOptionQuestion err:")
		}

		err = d.DBConn.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&options).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] CreateOptionQuestion Insert option err: %s", err)
			return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateOptionQuestion err:")
		}

		return nil
	})
	return err
}

// 增加填空题
func (d *DBClientImpl) SetBlankQuestionBatch(questions []models.Question, blanks []models.Blank) error {
	err := d.DBConn.Transaction(func(tx *gorm.DB) error {
		err := d.DBConn.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&questions).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] CreateBlankQuestion Insert question err: %s", err)
			return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateBlankQuestion err:")
		}
		err = d.DBConn.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&blanks).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] CreateBlankQuestion Insert blank err: %s", err)
			return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateBlankQuestion err:")
		}

		return nil
	})
	return err
}

// 根据question id查找问题
func (d *DBClientImpl) GetQuestionSingle(id int64) (*models.QuestionInfo, error) {
	res := &models.QuestionInfo{}
	question := &models.Question{}
	// 获取问题信息
	err := d.DBConn.Where("question_id=?", id).First(question).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetQuestionSingle Select question err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQuestionSingle err:")
	}
	// 判断问题类型
	t := 0
	if question != nil {
		res.Info = question
		t = int(question.Question_Type)
	}
	// 选择题处理
	if t == 1 || t == 2 {
		ops := []models.Option{}
		// 获取选择题所有选项
		err := d.DBConn.Where("question_id=?", id).Find(&ops).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] GetQuestionSingle Select option err: %s", err)
			return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQuestionSingle err:")
		}
		res.Option = ops
	}
	// 填空题处理
	if t == 3 || t == 4 {
		b := &models.Blank{}
		// 获取选择题所有选项
		err := d.DBConn.Where("question_id=?", id).First(b).Error
		if err != nil {
			logrus.Errorln("[DBClientImpl] GetQuestionSingle Select blank err: %s", err)
			return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQuestionSingle err:")
		}
		res.Blank = b
	}
	return res, nil
}

// 根据query_id批量查询问题
func (d *DBClientImpl) GetQuestionBatch(id int64) (map[int64]*models.QuestionInfo, error) {
	res := make(map[int64]*models.QuestionInfo)
	questions := []models.Question{}
	// 查询时一定穿指针
	err := d.DBConn.Where("query_id=?", id).Find(&questions).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetQuestionSingle Select question err: %s", err)
		return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQuestionSingle err:")
	}
	for i := 0; i < len(questions); i++ {
		qid := questions[i].Question_ID
		qtype := questions[i].Question_Type
		// 初始化
		res[qid] = &models.QuestionInfo{}
		res[qid].Info = &questions[i]
		// 选择题处理
		if qtype == 1 || qtype == 2 {
			ops := []models.Option{}
			// 获取选择题所有选项
			err := d.DBConn.Where("question_id=?", qid).Find(&ops).Error
			if err != nil {
				logrus.Errorln("[DBClientImpl] GetQuestionSingle Select option err: %s", err)
				return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQuestionSingle err:")
			}
			res[qid].Option = ops
		}
		// 填空题处理
		if qtype == 3 || qtype == 4 {
			b := &models.Blank{}
			// 获取选择题所有选项
			err := d.DBConn.Where("question_id=?", qid).First(b).Error
			if err != nil {
				logrus.Errorln("[DBClientImpl] GetQuestionSingle Select blank err: %s", err)
				return nil, errors.Wrap(conf.DBSelectError, "[DBClientImpl] GetQuestionSingle err:")
			}
			res[qid].Blank = b
		}
	}
	return res, nil
}

// 提交问卷
// 批量存储选择题
func (d *DBClientImpl) CreateOptionResultBatch(res []models.OptionResult) error {
	err := d.DBConn.Create(res).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] CreateOptionResultBatch Insert op_res err: %s", err)
		return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateOptionResultBatch err:")
	}
	return nil
}

// 批量存储填空题
func (d *DBClientImpl) CreateBlankResultBatch(res []models.BlankResult) error {
	err := d.DBConn.Create(res).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] CreateBlankResultBatch Insert op_res err: %s", err)
		return errors.Wrap(conf.DBInsertError, "[DBClientImpl] CreateBlankResultBatch err:")
	}
	return nil
}

// 查询问卷结果
func (d *DBClientImpl) GetOptionResultByQuestion(id int64) (res map[int64]int, total int64, err error) {
	// 计算总数
	total = 0
	err = d.DBConn.Model(&models.OptionResult{}).Where("question_id=?", id).Count(&total).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetOptionResult Select op_res err: %s", err)
		return nil, 0, errors.Wrap(conf.DBInsertError, "[DBClientImpl] GetOptionResult err:")
	}
	// 分选项计算
	// select `op_id` from `optionresult` where `question_id` = id
	// 获取该问题下的全部答案
	var ops []int64
	err = d.DBConn.Select("op_id").Where("question_id=?", id).Find(ops).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetOptionResult Select op_res err: %s", err)
		return nil, 0, errors.Wrap(conf.DBInsertError, "[DBClientImpl] GetOptionResult err:")
	}
	// 分别计算每个选项各有多少人选
	counts := make(map[int64]int)
	for _, v := range ops {
		if num, ok := counts[v]; ok {
			counts[v] = num + 1
		}
		counts[v] = 1
	}
	res = counts
	return res, total, nil
}

func (d *DBClientImpl) GetBlankResultByQuestion(id int64, page int) (res []models.BlankResult, err error) {
	res = make([]models.BlankResult, 0)
	err = d.DBConn.Where("question_id=?", id).Offset(page - 1).Limit(conf.Page_Size).Find(res).Error
	if err != nil {
		logrus.Errorln("[DBClientImpl] GetOptionResult Select op_res err: %s", err)
		return nil, errors.Wrap(conf.DBInsertError, "[DBClientImpl] GetOptionResult err:")
	}
	return res, nil
}
