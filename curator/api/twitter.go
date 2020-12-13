package api

import (
	"fmt"

	c "github.com/ddddddO/oop-go/curator"
)

// APIの子クラス(Baseの孫クラス)
type TwitterCurator struct {
	c.ApiCurator // 継承

	Tweets []Tweet
}

// コンストラクタ
func NewTwitterCurator(target, clientID, clientSecret string) *TwitterCurator {
	return &TwitterCurator{
		c.ApiCurator{
			BaseCurator: c.BaseCurator{
				Target: target,
			},

			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		nil,
	}
}

// オーバーライド
func (tc *TwitterCurator) Reserch() error {
	fmt.Println("TwitterCurator.Reserch")
	return nil
}

type Tweet struct {
	id   int
	text string
}

func (tw *Tweet) Text() string {
	return tw.text
}
