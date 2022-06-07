// Code generated by entc, DO NOT EDIT.

package application

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcmvp/idm.go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// CreateBy applies equality check predicate on the "create_by" field. It's identical to CreateByEQ.
func CreateBy(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// UpdateBy applies equality check predicate on the "update_by" field. It's identical to UpdateByEQ.
func UpdateBy(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// AppName applies equality check predicate on the "app_name" field. It's identical to AppNameEQ.
func AppName(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppName), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// CreateByEQ applies the EQ predicate on the "create_by" field.
func CreateByEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// CreateByNEQ applies the NEQ predicate on the "create_by" field.
func CreateByNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateBy), v))
	})
}

// CreateByIn applies the In predicate on the "create_by" field.
func CreateByIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateBy), v...))
	})
}

// CreateByNotIn applies the NotIn predicate on the "create_by" field.
func CreateByNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateBy), v...))
	})
}

// CreateByGT applies the GT predicate on the "create_by" field.
func CreateByGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateBy), v))
	})
}

// CreateByGTE applies the GTE predicate on the "create_by" field.
func CreateByGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateBy), v))
	})
}

// CreateByLT applies the LT predicate on the "create_by" field.
func CreateByLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateBy), v))
	})
}

// CreateByLTE applies the LTE predicate on the "create_by" field.
func CreateByLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateBy), v))
	})
}

// CreateByContains applies the Contains predicate on the "create_by" field.
func CreateByContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreateBy), v))
	})
}

// CreateByHasPrefix applies the HasPrefix predicate on the "create_by" field.
func CreateByHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreateBy), v))
	})
}

// CreateByHasSuffix applies the HasSuffix predicate on the "create_by" field.
func CreateByHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreateBy), v))
	})
}

// CreateByEqualFold applies the EqualFold predicate on the "create_by" field.
func CreateByEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreateBy), v))
	})
}

// CreateByContainsFold applies the ContainsFold predicate on the "create_by" field.
func CreateByContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreateBy), v))
	})
}

// UpdateByEQ applies the EQ predicate on the "update_by" field.
func UpdateByEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByNEQ applies the NEQ predicate on the "update_by" field.
func UpdateByNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByIn applies the In predicate on the "update_by" field.
func UpdateByIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateBy), v...))
	})
}

// UpdateByNotIn applies the NotIn predicate on the "update_by" field.
func UpdateByNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateBy), v...))
	})
}

// UpdateByGT applies the GT predicate on the "update_by" field.
func UpdateByGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByGTE applies the GTE predicate on the "update_by" field.
func UpdateByGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLT applies the LT predicate on the "update_by" field.
func UpdateByLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLTE applies the LTE predicate on the "update_by" field.
func UpdateByLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByContains applies the Contains predicate on the "update_by" field.
func UpdateByContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUpdateBy), v))
	})
}

// UpdateByHasPrefix applies the HasPrefix predicate on the "update_by" field.
func UpdateByHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUpdateBy), v))
	})
}

// UpdateByHasSuffix applies the HasSuffix predicate on the "update_by" field.
func UpdateByHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUpdateBy), v))
	})
}

// UpdateByEqualFold applies the EqualFold predicate on the "update_by" field.
func UpdateByEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUpdateBy), v))
	})
}

// UpdateByContainsFold applies the ContainsFold predicate on the "update_by" field.
func UpdateByContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUpdateBy), v))
	})
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleted), v))
	})
}

// AppNameEQ applies the EQ predicate on the "app_name" field.
func AppNameEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppName), v))
	})
}

// AppNameNEQ applies the NEQ predicate on the "app_name" field.
func AppNameNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppName), v))
	})
}

// AppNameIn applies the In predicate on the "app_name" field.
func AppNameIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppName), v...))
	})
}

// AppNameNotIn applies the NotIn predicate on the "app_name" field.
func AppNameNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppName), v...))
	})
}

// AppNameGT applies the GT predicate on the "app_name" field.
func AppNameGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppName), v))
	})
}

// AppNameGTE applies the GTE predicate on the "app_name" field.
func AppNameGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppName), v))
	})
}

// AppNameLT applies the LT predicate on the "app_name" field.
func AppNameLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppName), v))
	})
}

// AppNameLTE applies the LTE predicate on the "app_name" field.
func AppNameLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppName), v))
	})
}

// AppNameContains applies the Contains predicate on the "app_name" field.
func AppNameContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppName), v))
	})
}

// AppNameHasPrefix applies the HasPrefix predicate on the "app_name" field.
func AppNameHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppName), v))
	})
}

// AppNameHasSuffix applies the HasSuffix predicate on the "app_name" field.
func AppNameHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppName), v))
	})
}

// AppNameEqualFold applies the EqualFold predicate on the "app_name" field.
func AppNameEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppName), v))
	})
}

// AppNameContainsFold applies the ContainsFold predicate on the "app_name" field.
func AppNameContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppName), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		p(s.Not())
	})
}
