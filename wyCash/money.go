package wyCash

//3~33行目のコメントアウトしたコードは、テスト駆動開発の過程で、削除したコードです。
/*
type Dollar struct {
	Money
}

// make method return same type of Object

func (d *Dollar) Times(i int) Dollar {
	return Dollar{d.amount * i}
}

func (d *Dollar) Equals(e Dollar) bool {
	return d.amount == e.amount
}


type Franc struct {
	Money
}

// make method return same type of Object

func (f *Franc) Times(i int) Franc {
	return Franc{f.amount * i}
}

func (f *Franc) Equals(e Franc) bool {
	return f.amount == e.amount
}
*/

type Currency string

const (
	DOLLAR Currency = "dollar"
	FRANC  Currency = "franc"
)

type Money struct {
	amount   int
	currency Currency
}

type Expression interface {
	Reduce(b Bank, c Currency) Money
}

// make method return same type of Object
func (m Money) Times(i int) Expression {
	return Money{m.amount * i, m.currency}
}

func (m Money) Equals(e Money) bool {
	return m.amount == e.amount && m.currency == e.currency
}

// 問題のメソッドです。
func (m *Money) Plus(e Expression) Sum {
	return Sum{m, e}
}

func (m Money) Reduce(b Bank, c Currency) Money {
	rate := b.rate(m.currency, c)
	return Money{m.amount / rate, c}
}

func NewDollar(i int) Money {
	return Money{i, DOLLAR}
}

func NewFranc(i int) Money {
	return Money{i, FRANC}
}

type Bank map[Pair]int

func (b Bank) Reduce(e Expression, c Currency) Money {
	return e.Reduce(b, c)
}

type Sum struct { // Sum belongs to Expression
	augend Expression
	addend Expression
}

func (s Sum) Reduce(b Bank, c Currency) Money {
	augend := s.augend.Reduce(b, c).amount
	addend := s.addend.Reduce(b, c).amount
	amount := augend + addend
	return Money{amount, c}
}

type Rate struct {
	before Currency
	after  Currency
	rate   int
}

func (b Bank) AddRate(r Rate) {
	b[Pair{r.before, r.after}] = r.rate
}

func (b Bank) rate(before Currency, after Currency) int {
	if before == after {
		return 1
	}
	p := Pair{before, after}
	return b[p]
}

type Pair struct {
	before Currency
	after  Currency
}

func (p Pair) Equals(a any) bool {
	pair := a.(Pair)
	return p.before == pair.before && p.after == pair.after
}

func (p Pair) HashCode() int {
	return 0
}
