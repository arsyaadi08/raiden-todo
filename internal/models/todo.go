package models

import (
	"time"

	"github.com/sev-2/raiden"
)

type Todo struct {
	raiden.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Task *string `json:"task,omitempty" column:"name:task;type:varchar;nullable"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
