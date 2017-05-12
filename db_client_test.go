package clientauth

import (
	"os"
	"testing"

	_ "github.com/lib/pq"

	log "github.com/Sirupsen/logrus"
	"github.com/irfn/goconfig"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	goconfig.BaseConfig
}

func init() {
	log.SetOutput(os.Stdout)
	conf := TestConfig{}
	conf.Load()
}

func TestCreatesAndFindsAAuthorizedApplication(t *testing.T) {
	db := initDB()
	db.Exec("INSERT INTO authorized_applications (client_id, pass_key) VALUES ('DUMMY-CLIENT-ID', 'DUMMY-PASSKEY')")

	authorizedApplication := Client{
		ClientId: "DUMMY-CLIENT-ID",
		PassKey:  "DUMMY-PASSKEY",
	}

	repository := &ClientRepository{DB: db}
	found, error := repository.GetClient(authorizedApplication.ClientId)

	assert.Equal(t, error, nil)

	assert.Equal(t, authorizedApplication.ClientId, found.ClientId)

	assert.Equal(t, authorizedApplication.PassKey, found.PassKey)

	db.Exec("DELETE FROM authorized_applications WHERE client_id = 'DUMMY-CLIENT-ID' ")
}
