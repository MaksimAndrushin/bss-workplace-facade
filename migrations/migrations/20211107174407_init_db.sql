-- +goose Up
CREATE TABLE workplaces_events_facade
(
    id           BIGSERIAL PRIMARY KEY NOT NULL,
    workplace_id BIGSERIAL NOT NULL,
    type         INT       NOT NULL,
    status       INT       NOT NULL,
    updated      TIMESTAMP,
    payload      JSONB
);

CREATE INDEX wrkpl_event_status_idx ON workplaces_events_facade USING btree (status);

-- +goose Down
DROP TABLE workplaces_events_facade;
