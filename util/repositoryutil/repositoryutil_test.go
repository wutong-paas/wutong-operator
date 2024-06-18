package repositoryutil

import "testing"

func TestLoginRepository(t *testing.T) {
	err := LoginRepository("wutong.me", "a", "b")
	if err != nil {
		t.Error(err)
	}
}
