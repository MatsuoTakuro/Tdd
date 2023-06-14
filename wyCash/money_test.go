package wyCash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*　ToDoリスト
[x] $Money{5} + Money{10}CHF = $10
[x] $Money{5} * 2  $Money{10}
[x] make Dollar.amount private
[x] avoid the side effect on amount of Dollar
[x] Equals()
[] hashCode()
[] make TestDollar as main test, and insert TestTimes() and TestEquals() into it.
[] round floatMoney{6}4 value
[x] remove duplications on Dollar and Franc
[] check the equivalent on null
[] check the equivalent on other type of object
[] make kind in Money{} unique type
[] Currency のテストをどうするか？
    [] 一度定義したCurrencyの値を変更不可にする
	[] Currencyと文字列の演算はできない
	[] Currencyの初期化には専用の関すが必要である
*/

// テスト駆動開発の過程で、削除したコードです。
/*
func TestDollar(t *testing.T) {
	t.Run("test Dollar.Times()", func(t *testing.T) {
		assert := assert.New(t)
		five := Dollar{Money{5}}

		ten := five.Times(2)
		assert.Equal(ten, Dollar{Money{10}}, "Dollar(Money{5}).Times(2) should be 10!")

		// continuously testing a same function(or method) proves that it doesnt have a side effect.
		fifteen := five.Times(3)
		assert.Equal(fifteen, Dollar{Money{5}}, "Dollar(5).Times(3) should be 15!")
	})
	t.Run("test Dollar.Equals()", func(t *testing.T) {
		assert := assert.New(t)

		firstDollar := Dollar{Money{5}}
		secondDollar := Dollar{Money{5}}
		thirdDollar := Dollar{Money{6}}

		assert.Equal(true, firstDollar.Equals(secondDollar), "fitstDollar and secondDollar must be equivalent!")
		assert.Equal(false, firstDollar.Equals(thirdDollar), "fitstDollar and secondDollar must be NOT equivalent!")
	})
}

func TestFranc(t *testing.T) {
	t.Run("test Franc.Times()", func(t *testing.T) {
		assert := assert.New(t)
		five := Franc{Money{5}}

		ten := five.Times(2)
		assert.Equal(ten, Franc{Money{10}}, "Franc(Money{5}).Times(2) should be 10!")

		// continuously testing a same function(or method) proves that it doesnt have a side effect.
		fifteen := five.Times(3)
		assert.Equal(fifteen, Franc{Money{15}}, "Franc(5).Times(3) should be 15!")
	})
	t.Run("test Franc.Equals()", func(t *testing.T) {
		assert := assert.New(t)

		firstFranc := Franc{Money{5}}
		secondFranc := Franc{Money{5}}
		thirdFranc := Franc{Money{6}}

		assert.Equal(true, firstFranc.Equals(secondFranc), "fitstFranc and secondFranc must be equivalent!")
		assert.Equal(false, firstFranc.Equals(thirdFranc), "fitstFranc and secondFranc must be NOT equivalent!")
	})
}
*/

func TestMoney(t *testing.T) {
	t.Run("test Money.Times()", func(t *testing.T) {
		assert := assert.New(t)

		//five := NewDollar(5)

		ten := (NewDollar(5)).Times(2)
		assert.Equal(ten, NewDollar(10), "Money.Dollar(5).Times(2) should be Money.Dollar(10)!")

		// continuously testing a same function(or method) proves that it doesnt have a side effect.
		//fifteen := five.Times(3)
		//assert.Equal(fifteen, NewDollar(15), "Money.Dollar(5).Times(3) should be Money.Dollar(15)!")
	})

	t.Run("test Money.Equals()", func(t *testing.T) {
		assert := assert.New(t)

		firstDollar := NewDollar(5)
		secondDollar := NewDollar(5)
		thirdDollar := NewDollar(10)
		franc := NewFranc(5)

		assert.Equal(true, firstDollar.Equals(secondDollar), "Money{5, DOLLAR} and Money{5, DOLLAR} must be equivalent!")
		assert.Equal(false, firstDollar.Equals(thirdDollar), "Money{5, DOLLAR} and Money{10, DOLLAR} must be NOT equivalent!")
		assert.Equal(false, firstDollar.Equals(franc), "Money{5, DOLLAR} and Money{5, FRANC} must be NOT equivalent!")
	})
}

/*
func TestCurrency(t *testing.T) {
	t.Run("test unable to change value in Currency type object", func(t *testing.T) {
		c := NewCurrency()
		c = "changed value!"
		assert.Equal("USD", c, "")
	})
}
*/

func TestMoneyCalculation(t *testing.T) {
	assert := assert.New(t)
	bank := Bank{}
	m := NewDollar(3)
	pm := &m                     // covert a variable of Money type to one of *Money type cuz Plus method can be called only by an variable of *Money type, not Monney type.
	sum := pm.Plus(NewDollar(4)) // this is not Money but Expression!
	reduced := bank.Reduce(sum, DOLLAR)
	assert.Equal(NewDollar(7), reduced) // <= assert.Equal(NewDollar(10), sum, "")
}

// 問題のテストコードです。
func TestPlusReturnsSum(t *testing.T) {
	assert := assert.New(t)
	m1 := NewDollar(5)
	pm1 := &m1                       // covert a variable of Money type to one of *Money type cuz Plus method can be called only by an variable of *Money type, not Monney type.
	result := pm1.Plus(NewDollar(5)) // this is not Money but Expression!
	sum := Sum(result)               // Sum Struct. (at this point, the pm1 of *Money type is assiged to sum.augend)
	m2 := NewDollar(5)
	pm2 := &m2 // covert a variable of Money type to one of *Money type cuz it will come to be able to comparable as (the same) variable of *Money type with sum.augend in the assertion of the next line.
	assert.Equal(pm2, sum.augend)
	assert.Equal(NewDollar(5), sum.addend)
}

func TestReduceMoney(t *testing.T) {
	assert := assert.New(t)
	bank := Bank{}
	ten := NewDollar(10)
	result := bank.Reduce(ten, DOLLAR)
	assert.Equal(NewDollar(10), result)
}

func TestReduceFranceToDollar(t *testing.T) {
	assert := assert.New(t)
	bank := Bank{}
	r := Rate{FRANC, DOLLAR, 2}
	bank.AddRate(r)
	result := bank.Reduce(NewFranc(10), DOLLAR)
	assert.Equal(NewDollar(5), result)
}

func TestIdentitiyRate(t *testing.T) {
	assert := assert.New(t)
	var b Bank
	assert.Equal(1, b.rate(DOLLAR, DOLLAR))
}

func TestMixedAddition(t *testing.T) {
	assert := assert.New(t)
	d := NewDollar(5)
	f := NewFranc(4)
	b := Bank{}
	r := Rate{FRANC, DOLLAR, 2}
	b.AddRate(r)
	resultDollar := b.Reduce(d.Plus(f), DOLLAR)
	assert.Equal(NewDollar(7), resultDollar)

	/*
		resultFranc := b.Reduce(f.Plus(d), FRANC)
		assert.Equal(NewFranc(14), resultFranc)

		resultTranslatedDollar := b.Reduce(f.Plus(d), DOLLAR)
		assert.Equal(NewDollar(7), resultTranslatedDollar)
	*/
}
