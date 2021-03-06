package model

type Belt struct {
	ID          uint64 `db:"id,omitempty"`
	Slug        string `db:"slug"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Fabric      string `db:"fabric"`
	Front       string `db:"front"`
	Back        string `db:"back"`
}
