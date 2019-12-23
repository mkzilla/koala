package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/utils"
)

func convertStateToAuth(state string) *config.OAuthIF {
	switch state {
	case "github":
		return &config.Configs.GithubLogin
	case "basic":
		return &config.Configs.OAuthV2BasicLogin
	}
	return nil
}

func LoginWithOAuth2(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	cfg := convertStateToAuth(state)
	if cfg == nil {
		types.HandleError(c, types.OAuthV2Error, nil)
		return
	}

	if code == "" {
		c.Redirect(http.StatusFound, cfg.Config.AuthCodeURL(state))
		return
	}
	token, err := cfg.Config.Exchange(c.Request.Context(), code)
	if err != nil {
		types.HandleError(c, types.OAuthV2Error, err)
		return
	}
	u, err := getUserinfo(token.AccessToken, *cfg)
	if err != nil {
		types.HandleError(c, types.OAuthV2Error, err)
		return
	}
	var usr types.User
	ext, err := config.DBEngine.Where("globalUnique = ?", u.GlobalUnique).Get(&usr)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if ext && !usr.Enable {
		types.HandleError(c, types.InBlackList, nil)
		return
	}
	if ext && usr.Enable {
		u.Username = usr.Username
		_, err := config.DBEngine.Update(u, usr)
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
			return
		}
		u.ID = usr.ID
	}
	if !ext {
		cnt := types.GetCountByUsername(u.Username)
		if cnt > 0 {
			u.Username = fmt.Sprintf("%s%d", u.Username, cnt)
		}
		_, err = config.DBEngine.InsertOne(u)
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
			return
		}
	}
	apiToken, err := createToken(u.ID)
	if err != nil {
		types.HandleError(c, types.BadSystemRSAService, err)
		return
	}
	c.SetCookie("token", apiToken, 86400, "/", "", false, false)
	c.Redirect(http.StatusFound, "/")
}

func createToken(id int64) (string, error) {
	now := time.Now()
	expSecond := 86400
	jt := jwt.NewWithClaims(jwt.SigningMethodRS256, config.Token{
		"koala",
		now.Unix(),
		now.Add(time.Duration(expSecond) * time.Second).Unix(),
		id,
	})

	return jt.SignedString(config.RsaPrivateKey)
}

func getUserinfo(token string, cfg config.OAuthIF) (u *types.User, err error) {
	req, _ := http.NewRequest("GET", cfg.API, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return ConvertUserInfo(body, cfg)
}

func ConvertUserInfo(data []byte, o config.OAuthIF) (*types.User, error) {
	var usr = types.User{
		Platform: o.Name,
		Enable:   true,
	}
	switch o.Name {
	case "github":
		var u config.GithubUser
		err := json.Unmarshal(data, &u)
		if err != nil {
			return nil, err
		}
		usr.Username = u.Login
		usr.RealName = u.Login
		usr.Nickname = u.Name
		usr.Email = u.Email
		usr.GlobalUnique = fmt.Sprintf("github-%d", u.ID)
	case "basic":
		un, err := parseValue(data, o.BasicOauthUserMap.Username)
		if err != nil {
			return nil, err
		}
		nk, err := parseValue(data, o.BasicOauthUserMap.Nickname)
		if err != nil {
			return nil, err
		}
		em, err := parseValue(data, o.BasicOauthUserMap.Email)
		if err != nil {
			return nil, err
		}
		uq, err := parseValue(data, o.BasicOauthUserMap.Unique)
		if err != nil {
			return nil, err
		}
		usr.Username = fmt.Sprintf("%s", un)
		usr.RealName = fmt.Sprintf("%s", un)
		usr.Nickname = fmt.Sprintf("%s", nk)
		usr.Email = fmt.Sprintf("%s", em)
		usr.GlobalUnique = fmt.Sprintf("%s-%v", o.Name, uq)

	}
	return &usr, nil
}

func parseValue(data []byte, m string) (interface{}, error) {
	var tmp = make(map[string]interface{})
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return "", err
	}
	arr := strings.Split(m, ".")
	for k, v := range arr {
		if k == len(arr)-1 {
			return tmp[v], nil
		} else {
			b, err := json.Marshal(tmp[v])
			if err != nil {
				return "", err
			}
			err = json.Unmarshal(b, &tmp)
			if err != nil {
				return "", err
			}
		}
	}
	return "", nil
}

//
// 用户 Login 的时候，需要传入登录方式
func LoginWithPassword(c *gin.Context) {
	var u = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := c.ShouldBind(&u)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	var usr = types.User{
		Username: u.Username,
	}
	ext, err := config.DBEngine.Get(&usr)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if ext && usr.Enable {
		var pwd = types.Password{
			UserID:   usr.ID,
			Password: utils.Cipher(u.Password),
		}
		ext, err := config.DBEngine.Get(&pwd)
		if err != nil {
			types.HandleError(c, types.FailedToGetDataFromDB, err)
			return
		}
		if !ext {
			types.HandleError(c, types.BadUserPassword, err)
			return
		}
		apiToken, err := createToken(usr.ID)
		if err != nil {
			types.HandleError(c, types.BadSystemRSAService, err)
			return
		}
		c.SetCookie("token", apiToken, 86400, "/", "", false, false)
		c.JSON(http.StatusOK, "")
		return
	}
	types.HandleError(c, types.NoSuchUser, nil)
}

func Auth(c *gin.Context) {
	tk, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	var claim config.Token
	jwt.ParseWithClaims(tk, &claim, func(token *jwt.Token) (interface{}, error) {
		return config.RsaPrivateKey.PublicKey, nil
	})
	if claim.AUD == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	var usr = types.User{
		ID: claim.AUD,
	}
	ext, err := config.DBEngine.Get(&usr)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if !ext || !usr.Enable {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Set(config.User, usr)
}

func IsAdmin(c *gin.Context) {
	usr, ext := c.Get(config.User)
	if ext && usr.(types.User).IsAdmin {
		return
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}

func IsGroupUser(c *gin.Context) {
	if c.Param("group_id") != "" {
		usr, _ := c.Get(config.User)
		cnt, err := config.DBEngine.Table("group_user").
			Where("group_id = ? and user_id = ? and enable = true", c.Param("group_id"), usr.(types.User).ID).Count()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if cnt == 0 {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
	if c.Query("username") != "" && c.Param("group_id") != "" {
		var g types.Group
		_, err := config.DBEngine.Where("id = ?", c.Param("group_id")).Get(&g)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		usr, _ := c.Get(config.User)
		if usr.(types.User).ID != g.AdminID {
			c.AbortWithStatus(http.StatusForbidden)
		}
	} else if c.Query("username") != "" {
		var u types.User
		ext, err := config.DBEngine.Where("username = ?", c.Query("username")).Get(&u)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if !ext {
			c.AbortWithStatus(http.StatusNotFound)
		}
		usr, _ := c.Get(config.User)
		if usr.(types.User).Username == c.Query("username") {
			return
		}

		cnt, err := config.DBEngine.Table("group_user").Alias("gu").
			Join("inner", "group_user", "gu.group_id = group_user.group_id and gu.user_id=? and group_user.user_id=? and group_user.enable =true and gu.enable = true", u.ID, usr.(types.User).ID).Count()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if cnt == 0 {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}

	if c.Param("task_id") != "" {
		var t types.Task
		ext, err := config.DBEngine.Where("id = ?", c.Param("task_id")).Get(&t)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if !ext {
			c.AbortWithStatus(http.StatusNotFound)
		}
		if t.AssignTo == 0 {
			return
		}

		usr, _ := c.Get(config.User)
		if usr.(types.User).ID == t.AssignTo {
			return
		}
		cnt, err := config.DBEngine.Table("group_user").Alias("v1").
			Join("inner", "group_user", "group_user.group_id = v1.group_id and group_user.user_id=? and v1.user_id=? and group_user.enable =true and v1.enable = true", t.AssignTo, usr.(types.User).ID).Count()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if cnt == 0 {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

func UserInfo(c *gin.Context) {
	usr, _ := c.Get(config.User)
	c.JSON(http.StatusOK, usr)
}
