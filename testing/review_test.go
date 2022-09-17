package summary_test

import (
	"testing"

	review "github.com/HermanSetiawan77777/JakMall/internal/service/review"
	"github.com/stretchr/testify/assert"
)

func TestGetdata(t *testing.T) {
	if assert.Nil(t, review.GetData()) {
		t.Error("Dont have data !")
	}
}
