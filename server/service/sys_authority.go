package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/postgres"
	req "gin-vue-admin/model/postgres/request"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gorm.io/gorm"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func CreateAuthority(auth postgres.SysAuthority) (err error, authority postgres.SysAuthority) {
	var authorityBox postgres.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.GVA_DB.Create(&auth).Error
	return err, auth

}

//@author: [piexlmax](https://github.com/kiwi520)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority postgres.SysAuthority

func CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority postgres.SysAuthority) {
	var authorityBox postgres.SysAuthority

	if !errors.Is(global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("权限菜单存在相同角色id"), authority
	}

	type CasBinRule struct {
		id int
	}
	var casbin_rules []CasBinRule

	//判断CasBinRule表中是否存在此id的数据
	err = global.GVA_DB.Table("casbin_rule").Select("id").Where("v0 = ?", copyInfo.Authority.AuthorityId).Find(&casbin_rules).Error

	if err != nil {
		return err, authority
	}

	if len(casbin_rules) > 0 {
		return errors.New("api中存在相同角色id"), authority
	}

	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		//_ = DeleteAuthority(&copyInfo.Authority)
		return err, copyInfo.Authority
	}

	newAuthorityId, _ := strconv.Atoi(copyInfo.Authority.AuthorityId)
	err, menus := GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})

	if err != nil {
		return err, copyInfo.Authority
	}
	//获取菜单权限数据
	var AuthorityMenu []postgres.SysMenu
	for _, v := range menus {
		AuthorityMenu = append(AuthorityMenu, postgres.SysMenu{
			SysAuthorityAuthorityId: newAuthorityId,
			SysBaseMenuId:           v.SysBaseMenuId,
		})
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//添加权限菜单
		if err := tx.Create(&AuthorityMenu).Error; err != nil {
			// return any error will rollback
			return err
		}

		//添加权限
		if err := tx.Create(&copyInfo.Authority).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil { //如果入库失败删除casbin里的角色api数据
		ClearCasbin(0, copyInfo.Authority.AuthorityId)
	}

	return err, copyInfo.Authority
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth postgres.SysAuthority
//@return:err error, authority postgres.SysAuthority

func UpdateAuthority(auth postgres.SysAuthority) (err error, authority postgres.SysAuthority) {
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&postgres.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/kiwi520)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *postgres.SysAuthority
//@return: err error

func DeleteAuthority(auth *postgres.SysAuthority) (err error) {
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&postgres.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}

	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).First(&postgres.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//移除权限菜单操作
		if err := tx.Where("authority_id =?", auth.AuthorityId).Delete(&postgres.SysAuthority{}).Error; err != nil {
			// return any error will rollback
			return err
		}

		//移除权限操作
		if err := tx.Where("sys_authority_authority_id =?", auth.AuthorityId).Delete(&postgres.SysMenu{}).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return err
	} else {
		//删除casbin里的角色api数据
		ClearCasbin(0, auth.AuthorityId)
	}

	//global.GVA_DB.Where("authority_id =?",auth.AuthorityId).Delete(&postgres.SysAuthority{});
	//global.GVA_DB.Where("sys_authority_authority_id =?",auth.AuthorityId).Delete(&postgres.SysMenu{});

	//ClearCasbin(0, auth.AuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Table("sys_authorities")
	var authlist []req.SysAuthority
	var authority []postgres.SysAuthority
	db.Where("parent_id = '0'")
	err = db.Where("parent_id = '0'").Count(&total).Error
	err = db.Limit(limit).Offset(offset).Where("parent_id = '0'").Find(&authority).Error

	if len(authority) > 0 {
		for k := range authority {
			if auth, err := findChildrenAuthority(authority[k]); err == nil {
				authlist = append(authlist, auth)
			}
		}
	}

	return err, authlist, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: err error, sa model.SysAuthority

func GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return:error

func SetDataAuthority(auth model.SysAuthority) error {
	var s model.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func SetMenuAuthority(auth *[]postgres.SysMenu, authorityId string) error {
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Where("sys_authority_authority_id = ?", authorityId).Delete(&postgres.SysMenu{}).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Create(&auth).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return err
}

//func SetMenuAuthority(auth *model.SysAuthority) error {
//	var s model.SysAuthority
//	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
//	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
//	return err
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func findChildrenAuthority(authority postgres.SysAuthority) (list req.SysAuthority, err error) {
	//err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	var authorityList []postgres.SysAuthority
	err = global.GVA_DB.Table("sys_authorities").Where("parent_id = ?", authority.AuthorityId).Find(&authorityList).Error
	list.AuthorityName = authority.AuthorityName
	list.AuthorityId = authority.AuthorityId
	list.ParentId = authority.ParentId
	list.DefaultRouter = authority.DefaultRouter
	list.CreatedAt = authority.CreatedAt
	list.UpdatedAt = authority.UpdatedAt
	if len(authorityList) > 0 {
		var objList []req.SysAuthority
		for k := range authorityList {
			author, _ := findChildrenAuthority(authorityList[k])
			objList = append(objList, author)
		}
		list.Children = append(list.Children, objList...)
	}
	return list, err
}
