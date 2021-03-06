package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/utils"

	log "github.com/sirupsen/logrus"
)

func Request(c *gin.Context) {
	var usr types.User
	c.ShouldBind(&usr)
	arr := strings.Split(usr.Email, "@")
	if len(arr) != 2 {
		c.JSON(http.StatusBadRequest, "incorrect email address")
		return
	}
	usr.RealName = arr[0]
	usr.Username = arr[0]
	usr.Nickname = arr[0]
	usr.Platform = "basic"
	usr.IsAdmin = false
	usr.Enable = false
	usr.GlobalUnique = fmt.Sprintf("%s-%s", usr.Platform, usr.Email)
	cnt := types.GetCountByUsername(usr.Username)
	if cnt > 0 {
		usr.Username = fmt.Sprintf("%s%d", usr.Username, cnt)
	}
	_, err := config.DBEngine.InsertOne(&usr)
	if err != nil {
		log.WithError(err).WithField("email", usr.Email).Warn("failed to insert email message")
	}
	token, err := createToken(usr.ID)
	if err != nil {
		types.HandleError(c, types.BadSystemRSAService, err)
		return
	}
	url := fmt.Sprintf("%s/password/%s?token=%s", config.Configs.Homepage, usr.Username, token)
	body := strings.ReplaceAll(config.Configs.BasicRegister.Template, "URL", url)
	err = utils.SendMailWithPassword(config.Configs.BasicRegister.Hostname,
		config.Configs.BasicRegister.Port, config.Configs.BasicRegister.Sender, config.Configs.BasicRegister.Password,
		usr.Email, config.Configs.BasicRegister.Subject, body)
	if err != nil {
		types.HandleError(c, types.FailedToSendEmail, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func InitPassword(c *gin.Context) {
	var u uLogin
	err := c.ShouldBind(&u)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	usr2 := usr.(types.User)
	var pass = types.Password{
		UserID:   usr2.ID,
		Password: utils.Cipher(u.Password),
	}
	_, err = config.DBEngine.InsertOne(pass)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	usr2.Enable = true
	_, err = config.DBEngine.ID(pass.UserID).AllCols().Update(usr2)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
