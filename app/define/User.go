package define

import (
	"github.com/gogf/gf/util/gmeta"
	"gof/app/model"
)

// 用户详情
type UserDetail struct {
	gmeta.Meta `orm:"table:user_detail"`
	Uid        int    `json:"uid"`
	Address    string `json:"address"`
}

// 用户学分
type UserScores struct {
	gmeta.Meta `orm:"table:user_scores"`
	Id         int `json:"id"`
	Uid        int `json:"uid"`
	Score      int `json:"score"`
}

// 用户信息
type User struct {
	gmeta.Meta `orm:"table:user"`
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	UserDetail *UserDetail   `orm:"with:uid=id"`
	UserScores []*UserScores `orm:"with:uid=id"`
}

type ChannelSiteInfo struct {
	gmeta.Meta `orm:"table:orders"`
	model.Orders
	ChildInfo      *ChildInfo      `orm:"with:sn" json:"child_info"`
}

type ChildInfo struct {
	gmeta.Meta `orm:"table:order_child"`
	model.OrderChild
	ChildGoodsInfo *ChildGoodsInfo `orm:"with:order_id=id" json:"child_goods_info"`
}
type ChildGoodsInfo struct {
	gmeta.Meta `orm:"table:order_goods"`
	model.OrderGoods
}
