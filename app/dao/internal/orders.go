// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// OrdersDao is the manager for logic model data accessing and custom defined data operations functions management.
type OrdersDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns OrdersColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// OrdersColumns defines and stores column names for table orders.
type OrdersColumns struct {
	Id           string //                                         
    Sn           string // 订单号（父）                            
    SellerId     string // 商户id                                  
    AppId        string // 应用id                                  
    ChannelId    string //                                         
    ThirdSn      string // 三方订单号                              
    PayStatus    string // 0:未扣款1:扣款成功-1:扣款失败           
    LogisticFee  string // 物流费用【分】                          
    GoodsFee     string // 商品费用【分】                          
    Consignee    string // 收件人姓名                              
    Phone        string // 收件人联系方式                          
    Province     string // 省                                      
    City         string // 市                                      
    Area         string //                                         
    Street       string //                                         
    Address      string //                                         
    Status       string // 0未处理1:下单成功-1:下单失败-2超额订单  
    Remark       string // 备注                                    
    CheckStatus  string // 0:未检测1:检测成功-1:检测失败           
    CreatedTime  string // 创建时间                                
    Origin       string // 1:SCM 2:微盟
}

//  ordersColumns holds the columns for table orders.
var ordersColumns = OrdersColumns{
	Id:          "id",            
            Sn:          "sn",            
            SellerId:    "seller_id",     
            AppId:       "app_id",        
            ChannelId:   "channel_id",    
            ThirdSn:     "third_sn",      
            PayStatus:   "pay_status",    
            LogisticFee: "logistic_fee",  
            GoodsFee:    "goods_fee",     
            Consignee:   "consignee",     
            Phone:       "phone",         
            Province:    "province",      
            City:        "city",          
            Area:        "area",          
            Street:      "street",        
            Address:     "address",       
            Status:      "status",        
            Remark:      "remark",        
            CheckStatus: "check_status",  
            CreatedTime: "created_time",  
            Origin:      "origin",
}

// NewOrdersDao creates and returns a new DAO object for table data access.
func NewOrdersDao() *OrdersDao {
	return &OrdersDao{
		Group:   "default",
		Table:   "orders",
		Columns: ordersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrdersDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrdersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}