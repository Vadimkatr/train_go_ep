package postgresstore

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"

	"task8/pkg/model"
	"task8/pkg/repository"
)

type ContactsRepositoryInDB struct {
	db     *sql.DB
	logger *logrus.Logger
	lastID uint
}

func NewContactsRepositoryInDB(db *sql.DB) *ContactsRepositoryInDB {
	return &ContactsRepositoryInDB{
		db:     db,
		logger: logrus.New(),
		lastID: 0,
	}
}

func (r *ContactsRepositoryInDB) Save(contact model.Contact) (model.Contact, error) {
	_, err := r.db.Exec(
		"INSERT INTO contacts (id, first_name, last_name, phone, email) " +
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		r.lastID + 1,
		contact.FirstName,
		contact.LastName,
		contact.Phone,
		contact.Email,
		)
	if err != nil {
		r.logger.Error(err)
		return contact, repository.ErrNotUniqueEmailOrPhone
	}

	r.lastID ++

	return contact, nil
}

func (r *ContactsRepositoryInDB) ListAll() ([]model.Contact, error) {
	rows, err := r.db.Query(
		"SELECT id, first_name, last_name, phone, email FROM contacts",
		)

	defer rows.Close()

	if err != nil {
		r.logger.Error(err)
		return nil, repository.ErrContactListIsEmpty
	}

	var contacts = make([]model.Contact, 0)
	for rows.Next() {
		var tc model.Contact
		err := rows.Scan(&tc.ID, &tc.FirstName, &tc.LastName, &tc.Phone, &tc.Email)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}
		contacts = append(contacts, tc)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return contacts, nil
}

func (r *ContactsRepositoryInDB) GetByID(id uint) (model.Contact, error) {
	c := model.Contact{}
	if err := r.db.QueryRow(
		"SELECT id, first_name, last_name, phone, email FROM contacts WHERE id = $1",
			id,
		).Scan(
			&c.ID,
			&c.FirstName,
			&c.LastName,
			&c.Phone,
			&c.Email,
		); err != nil {
			if err == sql.ErrNoRows {
				return model.Contact{}, repository.ErrRecordNotFound
			}

			return model.Contact{}, err
		}

	return c, nil
}

func (r *ContactsRepositoryInDB) GetByPhone(phone string) (model.Contact, error) {
	c := model.Contact{}
	if err := r.db.QueryRow(
		"SELECT id, first_name, last_name, phone, email FROM contacts WHERE phone = $1",
		phone,
	).Scan(
		&c.ID,
		&c.FirstName,
		&c.LastName,
		&c.Phone,
		&c.Email,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Contact{}, repository.ErrRecordNotFound
		}

		return model.Contact{}, err
	}

	return c, nil
}

func (r *ContactsRepositoryInDB) GetByEmail(email string) (model.Contact, error) {
	c := model.Contact{}
	if err := r.db.QueryRow(
		"SELECT id, first_name, last_name, phone, email FROM contacts WHERE email = $1",
		email,
	).Scan(
		&c.ID,
		&c.FirstName,
		&c.LastName,
		&c.Phone,
		&c.Email,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Contact{}, repository.ErrRecordNotFound
		}

		return model.Contact{}, err
	}

	return c, nil
}

func (r *ContactsRepositoryInDB) SearchByName(n string) ([]model.Contact, error) {
	searchNamePattern := fmt.Sprintf("%s%%", n)

	rows, err := r.db.Query(
		"SELECT id, first_name, last_name, phone, email FROM contacts " +
			"WHERE first_name LIKE $1 OR last_name LIKE $2",
		searchNamePattern,
		searchNamePattern,
	)

	defer rows.Close()

	if err != nil {
		r.logger.Error(err)
		return nil, repository.ErrContactListIsEmpty
	}

	var contacts = make([]model.Contact, 0)
	for rows.Next() {
		var tc model.Contact
		err := rows.Scan(&tc.ID, &tc.FirstName, &tc.LastName, &tc.Phone, &tc.Email)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}
		contacts = append(contacts, tc)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return contacts, nil
}

func (r *ContactsRepositoryInDB) Delete(id uint) error {
	_, err := r.db.Exec(
		"DELETE FROM contacts WHERE id = $1",
		id,
	)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	return nil
}
