package controllers

import (
	"app/configs"
	"app/middlewares"
	"app/models/users"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	var userCreate users.UserRegister
	c.Bind(&userCreate)
	var userDB users.User
	userDB.Name = userCreate.Name
	userDB.Email = userCreate.Email
	userDB.Password = userCreate.Password
	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "Failed insert user", nil,
		})
	}
	/*
		token, err := middlewares.GenerateToken(int(userDB.Id))
		userRensponse := users.UserToken{
			userDB.Id, userDB.Name, userDB.Email, userDB.Password, userDB.CreatedAt, userDB.UpdatedAt, token,
		}
	*/
	return c.JSON(http.StatusOK, users.UserResponseSingle{
		true, "Succsess insert user", userDB,
	})
}

func GetListUsers(c echo.Context) error {
	var dataUsers []users.User
	err := configs.DB.Find(&dataUsers).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed get user database", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponses{
		true, "success get user database", dataUsers,
	})
}

func GetUserByID(c echo.Context) error {
	var dataUsers users.User
	userID := c.Param("userid")
	err := configs.DB.First(&dataUsers, userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed get user database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponseSingle{
		true, "success get user database by id", dataUsers,
	})
}

func LoginUser(c echo.Context) error {
	var userCheck users.User
	email := c.Param("email")
	password := c.Param("password")
	err := configs.DB.Find(&userCheck, "email = ? AND password = ?", email, password).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed login user", nil,
		})
	}
	/*
		token, err := middlewares.GenerateToken(int(userDB.Id))
		userRensponse := users.UserToken{
			userDB.Id, userDB.Name, userDB.Email, userDB.Password, userDB.CreatedAt, userDB.UpdatedAt, token,
		}
	*/
	id := userCheck.Id
	getname := userCheck.Name
	getemail := userCheck.Email
	getpass := userCheck.Password

	token, err := middlewares.GenerateToken(int(id))
	userResponse := users.UserToken{
		id, getname, getemail, getpass, userCheck.CreatedAt, userCheck.UpdatedAt, token,
	}

	return c.JSON(http.StatusOK, users.UserResponseToken{
		true, "success login user", userResponse,
	})
}

func UpdateUser(c echo.Context) error {
	var userUpdate users.UserUpdater
	c.Bind(&userUpdate)
	var userDB users.User
	userDB.Name = userUpdate.Name
	userDB.Email = userUpdate.Email
	userDB.Password = userUpdate.Password
	userID := c.Param("userid")

	err := configs.DB.Where("id = ?", userID).Updates(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed update user database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponses{
		true, "success update user database by id", nil,
	})
}

func DeleteUser(c echo.Context) error {
	var dataUsers []users.User
	userID := c.Param("userid")
	err := configs.DB.Delete(&dataUsers, userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed delete user database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponses{
		true, "success delete user database by id", nil,
	})
}

func CountUsers(c echo.Context) error {
	var dataUsers []users.User
	err := configs.DB.Find(&dataUsers).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed count user", nil,
		})
	}
	countUser := configs.DB.Find(&dataUsers).RowsAffected
	return c.JSON(http.StatusOK, users.UserResponseCount{
		true, "success count user", countUser,
	})
}
