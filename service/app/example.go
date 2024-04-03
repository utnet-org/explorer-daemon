package app

// @Tags Example
// @Summary Example
// @Accept json
// @Description Example API
// @Param param body types.Example true "Request Params"
// @Success 200 {object} types.ExampleRes "Success Response"
// @Router /example [post]
//func Example(c *fiber.Ctx) error {
//	var msgReq types.
//	err := c.BodyParser(&msgReq)
//	if err != nil {
//		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
//	}
//	ex := types.ExampleRes{
//		Phone: "13666666666",
//	}
//	return c.JSON(pkg.SuccessResponse(ex))
//}
