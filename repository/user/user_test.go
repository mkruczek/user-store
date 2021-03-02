package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareQuery(t *testing.T) {
	//given
	testCases := []struct {
		values map[string][]string
		want   string
	}{
		{map[string][]string{"first_name": {"jacek"}}, `SELECT id,first_name, last_name, email, create_date, update_date FROM users WHERE first_name='jacek'`},
		{map[string][]string{"first_name": {"jacek", "andrzej"}}, `SELECT id,first_name, last_name, email, create_date, update_date FROM users WHERE first_name='jacek' OR first_name='andrzej'`},
		{map[string][]string{"first_name": {"jacek", "andrzej"}, "last_name": {"placek"}}, `SELECT id,first_name, last_name, email, create_date, update_date FROM users WHERE first_name='jacek' OR first_name='andrzej' AND last_name='placek'`},
	}

	for _, c := range testCases {
		//when
		result := prepareQuery(c.values)

		//then
		assert.Equal(t, c.want, result)
	}
}
