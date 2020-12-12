package main

import (
	"fmt"
)

// インタフェース
// 収集・研究・保存・表示などを行うなにか
type Curator interface {
	Collect() error
	Reserch() error
	Save() error
	Show() error
}

// 親クラス
type BaseCurator struct {
	target string        // 情報を集める対象
	data   []interface{} // 収集した情報
}

func (bc *BaseCurator) Collect() error {
	fmt.Println("BaseCurator.Collect")
	fmt.Println(bc.target)
	bc.data = append(bc.data, "BaseCollect")
	return nil
}

func (bc *BaseCurator) Reserch() error {
	fmt.Println("BaseCurator.Reserch")
	return nil
}

func (bc *BaseCurator) Save() error {
	fmt.Println("BaseCurator.Save")
	return nil
}

func (bc *BaseCurator) Show() error {
	fmt.Println("BaseCurator.Show")
	fmt.Println(bc.data[0])
	return nil
}

// 親クラスここまで

// 子クラス
type TwitterCurator struct {
	BaseCurator // 継承
}

// コンストラクター
func NewTwitterCurator(target string) *TwitterCurator {
	return &TwitterCurator{
		BaseCurator{
			target: target,
		},
	}
}

// オーバーライド
func (tc *TwitterCurator) Reserch() error {
	fmt.Println("TwitterCurator.Reserch")
	return nil
}

// 子クラスここまで

func Run(curator Curator) error {
	if err := curator.Collect(); err != nil {
		return err
	}

	if err := curator.Reserch(); err != nil {
		return err
	}

	if err := curator.Save(); err != nil {
		return err
	}

	if err := curator.Show(); err != nil {
		return err
	}
	return nil
}

func main() {
	tc := NewTwitterCurator("Twitter")

	if err := Run(tc); err != nil {
		panic(err)
	}
}
