package template

func GameTemplateV2(start, takeTurn func(), haveWinner func() bool, winner func() int) int {
	start()
	for !haveWinner() {
		takeTurn()
	}
	return winner()
}
