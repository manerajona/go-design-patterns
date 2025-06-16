package memento

type Money float32

type BalanceMemento struct {
	balance Money
}

func newBalanceMemento(balance Money) *BalanceMemento {
	return &BalanceMemento{balance: balance}
}

func (m *BalanceMemento) GetBalance() Money {
	return m.balance
}

type BankAccount struct {
	log    []*BalanceMemento
	offset int
}

func NewBankAccount(balance Money) *BankAccount {
	return &BankAccount{
		log:    []*BalanceMemento{newBalanceMemento(balance)},
		offset: 0,
	}
}

func (b *BankAccount) GetBalance() Money {
	return b.log[b.offset].GetBalance()
}

func (b *BankAccount) snapshot(newBalance Money) {
	b.offset++
	// Discard history if new operation after undo
	if b.offset < len(b.log) {
		b.log = b.log[:b.offset]
	}
	b.log = append(b.log, newBalanceMemento(newBalance))
}

func (b *BankAccount) Push(amount float32) Money {
	newBalance := b.GetBalance() + Money(amount)
	b.snapshot(newBalance)
	return newBalance
}

func (b *BankAccount) Pull(amount float32) Money {
	newBalance := b.GetBalance() - Money(amount)
	b.snapshot(newBalance)
	return newBalance
}

func (b *BankAccount) Undo() Money {
	if b.offset > 0 {
		b.offset--
	}
	return b.GetBalance()
}

func (b *BankAccount) Redo() Money {
	if b.offset+1 < len(b.log) {
		b.offset++
	}
	return b.GetBalance()
}
