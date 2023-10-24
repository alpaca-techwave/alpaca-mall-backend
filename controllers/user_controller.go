package controllers

import (
	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	jtoken "github.com/golang-jwt/jwt/v4"
)


//  @Description Create user
//  @Tags       User
//  @Accept     json
//  @Produce    json
//  @Param user body models.User true "User object"
//  @Success    200  {object}  models.User
//  @Failure    500
//  @Failure    503
//  @Router     /user/create [post]
func CreateUser(c *fiber.Ctx) error {
  user := new(models.User)
  user.ID = uuid.New()
  if err := c.BodyParser(user); err != nil {
    return c.Status(503).SendString(err.Error())
  }
  config.Database.Create(&user)
  return c.Status(200).JSON(user)
}

//  @Description Get user by ID
//  @Tags       User
//  @Accept     json
//  @Produce    json
//  @Success    200  {object}  models.User
//  @Failure    500
//  @Failure    503
//  @Router     /user/info [get]
//  @Security   BearerAuth
func GetUserInfo(c *fiber.Ctx) error {
  userToken := c.Locals("user").(*jtoken.Token)
  claims := userToken.Claims.(jtoken.MapClaims) 
  id := claims["ID"].(string)
  idStr,err := uuid.Parse(id)
  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message":"Invalid ID",
    })
  }
  // idStr := c.Params("id")
  // id,err := uuid.Parse(idStr)
  // if err != nil {
  //   return c.Status(400).JSON(fiber.Map{
  //     "message":"Invalid ID",
  //   })
  // }
  var user models.User
  result := config.Database.Preload("Address").Find(&user, idStr)
  if result.RowsAffected == 0 {
    return c.Status(404).JSON(fiber.Map{
      "message":"Not found User",
    })
  }
  return c.Status(200).JSON(&user)
}
