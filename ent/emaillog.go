// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-message-center/ent/emaillog"
)

// EmailLog is the model entity for the EmailLog schema.
type EmailLog struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The target email address | 目标邮箱地址
	Target string `json:"target,omitempty"`
	// The subject | 发送的标题
	Subject string `json:"subject,omitempty"`
	// The content | 发送的内容
	Content string `json:"content,omitempty"`
	// The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败
	SendStatus   uint8 `json:"send_status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmailLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case emaillog.FieldSendStatus:
			values[i] = new(sql.NullInt64)
		case emaillog.FieldTarget, emaillog.FieldSubject, emaillog.FieldContent:
			values[i] = new(sql.NullString)
		case emaillog.FieldCreatedAt, emaillog.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case emaillog.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmailLog fields.
func (el *EmailLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emaillog.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				el.ID = *value
			}
		case emaillog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				el.CreatedAt = value.Time
			}
		case emaillog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				el.UpdatedAt = value.Time
			}
		case emaillog.FieldTarget:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target", values[i])
			} else if value.Valid {
				el.Target = value.String
			}
		case emaillog.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				el.Subject = value.String
			}
		case emaillog.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				el.Content = value.String
			}
		case emaillog.FieldSendStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field send_status", values[i])
			} else if value.Valid {
				el.SendStatus = uint8(value.Int64)
			}
		default:
			el.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmailLog.
// This includes values selected through modifiers, order, etc.
func (el *EmailLog) Value(name string) (ent.Value, error) {
	return el.selectValues.Get(name)
}

// Update returns a builder for updating this EmailLog.
// Note that you need to call EmailLog.Unwrap() before calling this method if this EmailLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (el *EmailLog) Update() *EmailLogUpdateOne {
	return NewEmailLogClient(el.config).UpdateOne(el)
}

// Unwrap unwraps the EmailLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (el *EmailLog) Unwrap() *EmailLog {
	_tx, ok := el.config.driver.(*txDriver)
	if !ok {
		panic("ent: EmailLog is not a transactional entity")
	}
	el.config.driver = _tx.drv
	return el
}

// String implements the fmt.Stringer.
func (el *EmailLog) String() string {
	var builder strings.Builder
	builder.WriteString("EmailLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", el.ID))
	builder.WriteString("created_at=")
	builder.WriteString(el.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(el.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("target=")
	builder.WriteString(el.Target)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(el.Subject)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(el.Content)
	builder.WriteString(", ")
	builder.WriteString("send_status=")
	builder.WriteString(fmt.Sprintf("%v", el.SendStatus))
	builder.WriteByte(')')
	return builder.String()
}

// EmailLogs is a parsable slice of EmailLog.
type EmailLogs []*EmailLog
