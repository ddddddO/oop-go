package curator

import (
	"fmt"
)

// 認証情報を使って認証するなにか
type Authenticator interface {
	Auth() error
}

// APIの親クラス(Baseの子クラス)
type ApiCurator struct {
	BaseCurator // 継承

	ClientID     string
	ClientSecret string
}

func (ac *ApiCurator) Auth() error {
	fmt.Printf("ApiCurator.Auth: %s, %s\n", ac.ClientID, ac.ClientSecret)
	return nil
}
