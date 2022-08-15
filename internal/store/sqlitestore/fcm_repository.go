package sqlitestore

import (
	"database/sql"
	"genxoft.dev/internal/model"
	"time"
)

type FcmRepository struct {
	db *sql.DB
}

func (r *FcmRepository) Find(token string) (*model.FcmSettings, error) {
	f := &model.FcmSettings{}
	var createdAt, updatedAt int64
	if err := r.db.QueryRow(
		"SELECT token, timezone, ip, created_at, updated_at FROM fcm_settings WHERE token = $1",
		token,
	).Scan(
		&f.Token,
		&f.Timezone,
		&f.IP,
		&createdAt,
		&updatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	f.CreatedAt = time.Unix(createdAt, 0)
	f.UpdatedAt = time.Unix(updatedAt, 0)

	return f, nil
}

func (r *FcmRepository) Create(f *model.FcmSettings) error {
	f.BeforeCreate()
	_, err := r.db.Exec(
		"INSERT INTO fcm_settings (token, timezone, ip, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		f.Token,
		f.Timezone,
		f.IP,
		f.CreatedAt.Unix(),
		f.UpdatedAt.Unix(),
	)

	return err
}

func (r *FcmRepository) Update(f *model.FcmSettings) error {
	f.BeforeUpdate()
	_, err := r.db.Exec(
		`UPDATE fcm_settings 
				SET timezone=$1, 
				    ip=$2,
				    created_at=$3, 
				    updated_at=$4
				WHERE
					token=$5`,
		f.Timezone,
		f.IP,
		f.CreatedAt.Unix(),
		f.UpdatedAt.Unix(),
		f.Token,
	)

	return err
}
