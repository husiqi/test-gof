package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"gof/app/dao"
	"gof/app/define"
	"gof/app/model"
)

var Test = testService{}

type testService struct{}

func (s *testService) Test() (res *define.ChannelSiteInfo, err error) {
	dao.Orders.Ctx(context.Background()).WithAll().Limit(1).Scan(&res)
	return
}

func Insert() error {
	ctx := context.Background()
	for i := 1; i <= 5; i++ {
		// User.
		user := model.User{
			Name: fmt.Sprintf(`name_%d`, i),
		}
		lastInsertId, err := dao.User.Ctx(ctx).Data(user).OmitEmpty().InsertAndGetId()
		if err != nil {
			return err
		}
		// Detail.
		userDetail := model.UserDetail{
			Uid:     uint(lastInsertId),
			Address: fmt.Sprintf(`address_%d`, lastInsertId),
		}
		_, err = dao.UserDetail.Ctx(ctx).Data(userDetail).OmitEmpty().Insert()
		if err != nil {
			return err
		}
		// Scores.
		for j := 1; j <= 5; j++ {
			userScore := model.UserScores{
				Uid:   uint(lastInsertId),
				Score: uint(j),
			}
			_, err = dao.UserScores.Ctx(ctx).Data(userScore).OmitEmpty().Insert()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Find() {

	res, _ := dao.Orders.Ctx(context.Background()).As("o").
		InnerJoin(dao.OrderChild.Table, "o1", "o.sn = o1.sn").
		Where("o."+dao.Orders.Columns.Id, 1).All()
	g.Dump(res)
}

func Redis() {

	g.Redis()
}

func Addr() {

}
