package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	req "gin-vue-admin/model/postgres/request"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u postgres.SysUser
//@return: err error, userInter postgres.SysUser

func Register(u postgres.SysUser) (err error, userInter postgres.SysUser) {
	var user postgres.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *postgres.SysUser
//@return: err error, userInter *postgres.SysUser

func Login(u *postgres.SysUser) ( userInter response.User,err error) {
	var user postgres.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	//err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	err = global.GVA_DB.Model(&postgres.SysUser{}).Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error

	userInter.Username = user.Username
	userInter.ID = user.ID
	userInter.UUID = user.UUID
	userInter.CreatedAt = user.CreatedAt
	userInter.NickName = user.NickName
	userInter.HeaderImg = user.HeaderImg
	userInter.AuthorityId = user.AuthorityId

	var authorities  req.SysAuthority
	var authority postgres.SysAuthority
	err = global.GVA_DB.Table("sys_authorities").Where("parent_id = ?", user.AuthorityId).Find(&authority).Error

	authorities, err = findChildrenAuthority(authority)
	if err != nil {
		return response.User{},err
	}
	userInter.Authority = authorities
	return userInter,nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *postgres.SysUser, newPassword string
//@return: err error, userInter *postgres.SysUser

func ChangePassword(u *postgres.SysUser, newPassword string) (err error, userInter *postgres.SysUser) {
	var user postgres.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&postgres.SysUser{})
	var userList []postgres.SysUser
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&postgres.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func DeleteUser(id float64) (err error) {
	var user postgres.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser postgres.SysUser
//@return: err error, user postgres.SysUser

func SetUserInfo(reqUser postgres.SysUser) (err error, user postgres.SysUser) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *postgres.SysUser

func FindUserById(id int) (err error, user *postgres.SysUser) {
	var u postgres.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *postgres.SysUser

func FindUserByUuid(uuid string) (err error, user *postgres.SysUser) {
	var u postgres.SysUser
	if err = global.GVA_DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
