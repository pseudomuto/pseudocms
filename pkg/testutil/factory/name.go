package factory

import (
	"fmt"
	"math/rand"

	"github.com/Pallinder/go-randomdata"
)

func name() string {
	return fmt.Sprintf("%s-%d", randomdata.Noun(), rand.Intn(1000))
}
