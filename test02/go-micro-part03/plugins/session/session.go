package session

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"net/http"
	"strings"
	"time"
)

var (
	sessionIdNamePrefix = "session-id-"
	store               *sessions.CookieStore
)

func init() {
	// 随机生成32位加密的key，切记，正式环境一定不要暴露，通过写到环境变量或其它安全方式
	// 我们是为了演示的步骤简单些，才直接硬编码
	store = sessions.NewCookieStore([]byte("OnNUU5RUr6Ii2HMI0d6E54bXTS52tCCL"))
}

func GetSession( r *http.Request) *sessions.Session {
	var sId string
	for _, c := range r.Cookies() {
		if strings.Index(c.Name, sessionIdNamePrefix) == 0 {
			sId = c.Name
			break
		}
	}
	if len(sId) == 0 {
		return  nil
	}
	session, _ := store.Get(r, sId)
	return session
}

func CreateSession(w http.ResponseWriter, r *http.Request, token jwt.MapClaims) *sessions.Session {

	sId := sessionIdNamePrefix + uuid.New().String()
	ses, _ := store.Get(r, sId)

	// 将sessionId设置到cookie中
	//cookie := &http.Cookie{Name: sId, Value: sId, Path: "/", Expires: time.Now().Add(30 * time.Second), MaxAge: 0}
	//http.SetCookie(w, cookie)

	ses.Values["userId"] = token["UserId"]
	ses.Values["userName"] = token["UserName"]

	// 保存新的session
	ses.ID = sId
	ses.Save(r, w)

	return ses
}


// GetSession 获取当前会话session
func GetSession2(w http.ResponseWriter, r *http.Request) *sessions.Session {
	// sessionId
	var sId string

	for _, c := range r.Cookies() {
		if strings.Index(c.Name, sessionIdNamePrefix) == 0 {
			sId = c.Name
			break
		}
	}

	// 如果cookie中没有sessionId的值，则使用随机生成
	if sId == "" {
		sId = sessionIdNamePrefix + uuid.New().String()
	}

	// 忽略错误，因为Get方法会一直都返回session
	ses, _ := store.Get(r, sId)

	// 没有id说明是新session
	if ses.ID == "" {
		// 将sessionId设置到cookie中
		cookie := &http.Cookie{Name: sId, Value: sId, Path: "/", Expires: time.Now().Add(30 * time.Second), MaxAge: 0}
		http.SetCookie(w, cookie)

		// 保存新的session
		ses.ID = sId
		ses.Save(r, w)
	}
	return ses
}
