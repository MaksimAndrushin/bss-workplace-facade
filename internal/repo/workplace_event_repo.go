package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"time"
)

type WorkplaceEventRepo interface {
	Add(ctx context.Context, event model.WorkplaceEvent) error
}

type workplaceEventRepo struct {
	db        *sqlx.DB
	batchSize uint
}

type workplaceEntity struct {
	ID      uint64    `db:"id"`
	Name    string    `db:"name"`
	Size    uint32    `db:"size"`
	Removed bool      `db:"removed"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}

const WORKPLACES_EVENTS_TAB = "workplaces_events_facade"
const WORKPLACES_EVENTS_ID = "id"
const WORKPLACES_EVENTS_PAYLOAD = "payload"
const WORKPLACES_EVENTS_TYPE = "type"
const WORKPLACES_EVENTS_WORKPLACE_ID = "workplace_id"
const WORKPLACES_EVENTS_STATUS = "status"
const WORKPLACES_EVENTS_UPDATED = "updated"

func NewWorkplaceEventRepo(db *sqlx.DB, batchSize uint) WorkplaceEventRepo {
	return &workplaceEventRepo{db: db, batchSize: batchSize}
}

func (r *workplaceEventRepo) Add(ctx context.Context, event model.WorkplaceEvent) error {

	query := sq.Insert(WORKPLACES_EVENTS_TAB).PlaceholderFormat(sq.Dollar).
		Columns(WORKPLACES_EVENTS_WORKPLACE_ID, WORKPLACES_EVENTS_TYPE, WORKPLACES_EVENTS_STATUS, WORKPLACES_EVENTS_UPDATED, WORKPLACES_EVENTS_PAYLOAD).
		Values(event.Entity.ID, event.Type, event.Status, "NOW()", event.Entity).
		Suffix("RETURNING id")

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := r.db.QueryContext(ctx, s, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return err
		}

		log.Debug().Msgf("Created event id - %v", id)

		return nil
	} else {
		return sql.ErrNoRows
	}
}

func (w *workplaceEntity) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("incompatible type for workplace")
	}

	res := &model.Workplace{}

	err := json.Unmarshal(source, res)

	if err != nil {
		return err
	}

	*w = workplaceEntity{
		ID:   res.ID,
		Name: res.Name,
		Size: res.Size,
	}

	return nil
}
