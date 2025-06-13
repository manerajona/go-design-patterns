package decorator

import (
	"testing"
)

func TestDragon_FlyAndCrawl(t *testing.T) {
	d := NewDragon()
	d.SetAge(9)
	d.Crawl() // Should print "Crawling!"
	d.Fly()   // Should not print "Flying!"
	d.SetAge(10)
	d.Crawl() // Should not print "Crawling!"
	d.Fly()   // Should print "Flying!"
}
