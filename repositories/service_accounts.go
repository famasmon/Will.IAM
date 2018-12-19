package repositories

import (
	"fmt"

	"github.com/ghostec/Will.IAM/models"
)

// ServiceAccounts repository
type ServiceAccounts interface {
	Get(string) (*models.ServiceAccount, error)
	ForEmail(string) (*models.ServiceAccount, error)
	ForKeyPair(string, string) (*models.ServiceAccount, error)
	Create(*models.ServiceAccount) error
}

type serviceAccounts struct {
	storage *Storage
}

func (sas serviceAccounts) Get(id string) (*models.ServiceAccount, error) {
	return nil, nil
}

// ForEmail retrieves Service Account corresponding
func (sas serviceAccounts) ForEmail(
	email string,
) (*models.ServiceAccount, error) {
	sa := &models.ServiceAccount{}
	if _, err := sas.storage.PG.DB.Query(
		sa, `SELECT id, name, key_id, key_secret, email
		FROM service_accounts WHERE email = ? LIMIT 1`, email,
	); err != nil {
		return nil, err
	}
	if sa.ID == "" {
		return nil, fmt.Errorf("Service Account not found for email %s", email)
	}
	return sa, nil
}

// ForKeyPair retrieves Service Account corresponding
func (sas serviceAccounts) ForKeyPair(
	keyID, keySecret string,
) (*models.ServiceAccount, error) {
	sa := &models.ServiceAccount{}
	if _, err := sas.storage.PG.DB.Query(
		&sa, `SELECT id, name, key_id, key_secret, email
		FROM service_accounts WHERE key_id = ? AND key_secret`, keyID, keySecret,
	); err != nil {
		return nil, err
	}
	return sa, nil
}

func (sas serviceAccounts) Create(sa *models.ServiceAccount) error {
	_, err := sas.storage.PG.DB.Query(
		sa, "INSERT INTO service_accounts (name, email) VALUES (?name, ?email) RETURNING id", sa,
	)
	return err
}

// NewServiceAccounts serviceAccounts ctor
func NewServiceAccounts(s *Storage) ServiceAccounts {
	return &serviceAccounts{storage: s}
}
