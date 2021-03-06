package vendingMachine

type VendingMachine struct {
	insertedMoney int
	coins         map[string]int
	items         map[string]int
}

func NewVendingMachine(coins, items map[string]int) *VendingMachine {
	return &VendingMachine{coins: coins, items: items}
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
	/*
		if coin == "T" {
			m.insertedMoney = 10
		} else if coin == "F" {
			m.insertedMoney = 5
		} else if coin == "TW" {
			m.insertedMoney = 2
		} else {
			m.insertedMoney = 1
		}
	*/

}

func (m *VendingMachine) SelectSD() string {
	price := m.items["SD"]
	change := m.insertedMoney - price
	return "SD" + m.change(change)
}

func (m *VendingMachine) SelectCC() string {
	price := m.items["CC"]
	change := m.insertedMoney - price
	return "CC" + m.change(change)
}

func (m *VendingMachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0; i < len(values); i++ {
		if c >= values[i] {
			str += ", " + coins[i]
			c -= values[i]
			i--
		}
	}
	/*
		if c >= values[0] {
			str += ", " + coins[0]
			c -= values[0]
		}
		if c >= values[1] {
			str += ", " + coins[1]
			c -= values[1]
		}
		if c >= values[2] {
			str += ", " + coins[2]
			c -= values[2]
		}
		if c >= values[3] {
			str += ", " + coins[3]
			c -= values[3]
		}
	*/
	return str
}

func (vm *VendingMachine) CoinReturn() string {
	coins := vm.change(vm.insertedMoney)
	vm.insertedMoney = 0
	return coins[2:len(coins)]
}
