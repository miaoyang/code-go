package dao

import (
	"code-go/core"
	"code-go/model/do"
)

// GetProfileById 获取个人信息
func GetProfileById(id int) *do.Profile {
	var profile do.Profile
	err := core.DB.Model(&profile).Where("ID = ?", id).First(&profile).Error
	//profile:=Redis.Get("profile")
	if err != nil {
		return nil
	}
	return &profile
}

// GetProfile 获取个人信息
//func GetProfile(c *gin.Context) *Profile {
//	var profile *Profile
//	sessionId, err := c.Cookie(SessionName)
//	//fmt.Println("seesionId:",sessionId)
//	if err != nil {
//		return nil
//	}
//	result, err := Redis.Get(sessionId).Result()
//	if err != nil {
//		return nil
//	}
//	err = json.Unmarshal([]byte(result), &profile)
//	if err != nil {
//		fmt.Println("unmarshall json false")
//	}
//	return profile
//}

// UpdateProfile 更新个人信息
//func UpdateProfile(c *gin.Context, id int, profile *Profile) int {
//	var _profile *Profile
//	err := global.DB.Model(&Profile{}).Where("ID = ?", id).Update(profile).Error
//	if err != nil {
//		return errmsg.ERROR
//	}
//	sessionId, err := c.Cookie(SessionName)
//	if err != nil {
//		return errmsg.ERROR
//	}
//	result, err := Redis.Get(sessionId).Result()
//	if err != nil {
//		return errmsg.ERROR
//	}
//	err = json.Unmarshal([]byte(result), &_profile)
//	if err != nil {
//		fmt.Println("unmarshall json false")
//	}
//	if _profile.ID == profile.ID {
//		//需要更新redis
//		profileJson, _ := json.Marshal(profile)
//		_, err = Redis.Set(sessionId, string(profileJson), utils.Expiration).Result()
//		if err != nil {
//			return errmsg.ERROR
//		}
//	}
//	return errmsg.SUCCESS
//}

// CreateProfile 新建个人信息
func CreateProfile(profile *do.Profile) error {
	err := core.DB.Model(&do.Profile{}).Create(&profile).Error
	if err != nil {
		return err
	}
	return nil
}
