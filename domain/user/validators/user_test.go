package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	//given
	cases := []struct {
		email  string
		result bool
	}{
		{"email@com", true},
		{"@com", false},
		{"X@com", true},
		{"", false},
		{"emailcom", false},
	}

	//when
	for _, c := range cases {
		result := validateEmailStruct(c.email)

		//then
		assert.Equal(t, c.result, result)
	}

}
