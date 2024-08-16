package commands_user

import (
	"github.com/alianmc1988/go-restfull-gin/database"
	user_entities "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/entities"
	user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"
)

func PatchUser(userID *string, updates *map[string]interface{})  (*user_entities.UserEntity,error){
	tx := database.DB.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}
   result := tx.Model(&user_models.User{}).Where("id = ?", userID).Updates(updates)
    if result.Error != nil {
        tx.Rollback()
        return nil, result.Error
    }

	 var updatedUser user_models.User
    if err := tx.Where("id = ?", userID).First(&updatedUser).Error; err != nil {
        tx.Rollback()
        return nil, err
    }
	if err:=tx.Commit().Error; err != nil {
		return nil, err
	}
	return user_entities.NewUserEntity(updatedUser.ID, updatedUser.FirstName, updatedUser.LastName, updatedUser.Email, updatedUser.Password), nil
}