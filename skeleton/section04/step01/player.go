package main

import "fmt"

type player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

// プレイヤーが行えるガチャの回数
func (p *player) drawableNum() int {
	return p.tickets + p.coin/10
}

func (p *player) draw(n int) {

	if p.drawableNum() < n {
		fmt.Println("ガチャ券またはコインが不足しています")
		return
	}

	if p.tickets > n {
		p.tickets -= n
		return
	}
	n -= p.tickets

	p.tickets = 0
	p.coin -= n * 10 // 1回あたり10枚消費する
}
