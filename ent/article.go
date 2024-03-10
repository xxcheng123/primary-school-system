// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xxcheng123/primary-school-system/ent/article"
	"github.com/xxcheng123/primary-school-system/ent/category"
	"github.com/xxcheng123/primary-school-system/ent/user"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	// 雪花ID
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Status holds the value of the "status" field.
	Status int64 `json:"status,omitempty"`
	// Img holds the value of the "img" field.
	Img string `json:"img,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleQuery when eager-loading is set.
	Edges            ArticleEdges `json:"edges"`
	category_article *int64
	user_article     *int64
	selectValues     sql.SelectValues
}

// ArticleEdges holds the relations/edges for other nodes in the graph.
type ArticleEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleEdges) CategoryOrErr() (*Category, error) {
	if e.Category != nil {
		return e.Category, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: category.Label}
	}
	return nil, &NotLoadedError{edge: "category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID, article.FieldStatus:
			values[i] = new(sql.NullInt64)
		case article.FieldTitle, article.FieldContent, article.FieldImg:
			values[i] = new(sql.NullString)
		case article.FieldCreateTime, article.FieldUpdateTime, article.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case article.ForeignKeys[0]: // category_article
			values[i] = new(sql.NullInt64)
		case article.ForeignKeys[1]: // user_article
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int64(value.Int64)
		case article.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case article.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				a.UpdateTime = value.Time
			}
		case article.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case article.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = value.Int64
			}
		case article.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				a.Img = value.String
			}
		case article.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_article", value)
			} else if value.Valid {
				a.category_article = new(int64)
				*a.category_article = int64(value.Int64)
			}
		case article.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_article", value)
			} else if value.Valid {
				a.user_article = new(int64)
				*a.user_article = int64(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Article.
// This includes values selected through modifiers, order, etc.
func (a *Article) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Article entity.
func (a *Article) QueryUser() *UserQuery {
	return NewArticleClient(a.config).QueryUser(a)
}

// QueryCategory queries the "category" edge of the Article entity.
func (a *Article) QueryCategory() *CategoryQuery {
	return NewArticleClient(a.config).QueryCategory(a)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return NewArticleClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("create_time=")
	builder.WriteString(a.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(a.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(a.Content)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", ")
	builder.WriteString("img=")
	builder.WriteString(a.Img)
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article
