package auth

import (
	"github.com/google/uuid"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/data"
	"golang.org/x/crypto/bcrypt"
)

type SessionManager struct {
	ZdkCtx *data.ZdkContext
	Config *config.Config
}

func NewSessionManager(ctx *data.ZdkContext, config *config.Config) *SessionManager {
	return &SessionManager{ZdkCtx: ctx, Config: config}
}

type UserSession struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (s *SessionManager) GenerateSession(data UserSession) (string, error) {
	sessionId := uuid.NewString()
	// jsonData, _ := json.Marshal(data)
	// s.ZdkCtx.DB.Exec()
	// err := s.Rdb.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	// if err != nil {
	// 	return "", err
	// }
	return sessionId, nil
}

func (s *SessionManager) SignIn(email, password string) (string, error) {
	// check if the user exists
	var user User
	// err := s.Conn.QueryRow("select id, first_name, last_name, email, password from users where email = ?", email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	// if err != nil {
	// 	return "", err
	// }

	// check if the password matches
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// create the session
	sessionId := uuid.NewString()
	// jsonData, _ := json.Marshal(UserSession{
	// 	Id:        user.Id,
	// 	FirstName: user.FirstName,
	// 	LastName:  user.LastName,
	// 	Email:     user.Email,
	// })

	// err = s.Rdb.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	// if err != nil {
	// 	return "", err
	// }

	return sessionId, nil
}

func (s *SessionManager) SignOut(sessionId string) error {
	// return s.Rdb.Del(context.Background(), sessionId).Err()
	return nil
}

func (s *SessionManager) GetSession(session string) (*UserSession, error) {
	// data, err := s.Rdb.Get(context.Background(), session).Result()
	// if err != nil {
	// 	return nil, err
	// }

	// unmarshal the data
	var userSession UserSession
	// err = json.Unmarshal([]byte(data), &userSession)
	// if err != nil {
	// 	return nil, err
	// }

	return &userSession, nil

}
