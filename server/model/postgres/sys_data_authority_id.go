package postgres

type DataAuthorityId struct {
	SysDataAuthorityId int `json:"sys_data_authority_id" gorm:"comment:角色id"`
	DataAuthorityIdAuthorityId int `json:"data_authority_id_authority_id" gorm:"comment:资源角色id"`
}

//func (s DataAuthorityId) TableName() string {
//	return "data_authority_ids"
//}
