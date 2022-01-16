// Code generated by entc, DO NOT EDIT.

package model

import (
	"entgo.io/ent/dialect/sql"
	"github.com/efectn/go-orm-benchmarks/benchs/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Fax applies equality check predicate on the "fax" field. It's identical to FaxEQ.
func Fax(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFax), v))
	})
}

// Web applies equality check predicate on the "web" field. It's identical to WebEQ.
func Web(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeb), v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// Right applies equality check predicate on the "right" field. It's identical to RightEQ.
func Right(v bool) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRight), v))
	})
}

// Counter applies equality check predicate on the "counter" field. It's identical to CounterEQ.
func Counter(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounter), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// FaxEQ applies the EQ predicate on the "fax" field.
func FaxEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFax), v))
	})
}

// FaxNEQ applies the NEQ predicate on the "fax" field.
func FaxNEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFax), v))
	})
}

// FaxIn applies the In predicate on the "fax" field.
func FaxIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFax), v...))
	})
}

// FaxNotIn applies the NotIn predicate on the "fax" field.
func FaxNotIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFax), v...))
	})
}

// FaxGT applies the GT predicate on the "fax" field.
func FaxGT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFax), v))
	})
}

// FaxGTE applies the GTE predicate on the "fax" field.
func FaxGTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFax), v))
	})
}

// FaxLT applies the LT predicate on the "fax" field.
func FaxLT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFax), v))
	})
}

// FaxLTE applies the LTE predicate on the "fax" field.
func FaxLTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFax), v))
	})
}

// FaxContains applies the Contains predicate on the "fax" field.
func FaxContains(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFax), v))
	})
}

// FaxHasPrefix applies the HasPrefix predicate on the "fax" field.
func FaxHasPrefix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFax), v))
	})
}

// FaxHasSuffix applies the HasSuffix predicate on the "fax" field.
func FaxHasSuffix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFax), v))
	})
}

// FaxEqualFold applies the EqualFold predicate on the "fax" field.
func FaxEqualFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFax), v))
	})
}

// FaxContainsFold applies the ContainsFold predicate on the "fax" field.
func FaxContainsFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFax), v))
	})
}

// WebEQ applies the EQ predicate on the "web" field.
func WebEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeb), v))
	})
}

// WebNEQ applies the NEQ predicate on the "web" field.
func WebNEQ(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeb), v))
	})
}

// WebIn applies the In predicate on the "web" field.
func WebIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWeb), v...))
	})
}

// WebNotIn applies the NotIn predicate on the "web" field.
func WebNotIn(vs ...string) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWeb), v...))
	})
}

// WebGT applies the GT predicate on the "web" field.
func WebGT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWeb), v))
	})
}

// WebGTE applies the GTE predicate on the "web" field.
func WebGTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWeb), v))
	})
}

// WebLT applies the LT predicate on the "web" field.
func WebLT(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWeb), v))
	})
}

// WebLTE applies the LTE predicate on the "web" field.
func WebLTE(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWeb), v))
	})
}

// WebContains applies the Contains predicate on the "web" field.
func WebContains(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWeb), v))
	})
}

// WebHasPrefix applies the HasPrefix predicate on the "web" field.
func WebHasPrefix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWeb), v))
	})
}

// WebHasSuffix applies the HasSuffix predicate on the "web" field.
func WebHasSuffix(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWeb), v))
	})
}

// WebEqualFold applies the EqualFold predicate on the "web" field.
func WebEqualFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWeb), v))
	})
}

// WebContainsFold applies the ContainsFold predicate on the "web" field.
func WebContainsFold(v string) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWeb), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// RightEQ applies the EQ predicate on the "right" field.
func RightEQ(v bool) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRight), v))
	})
}

// RightNEQ applies the NEQ predicate on the "right" field.
func RightNEQ(v bool) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRight), v))
	})
}

// CounterEQ applies the EQ predicate on the "counter" field.
func CounterEQ(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounter), v))
	})
}

// CounterNEQ applies the NEQ predicate on the "counter" field.
func CounterNEQ(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCounter), v))
	})
}

// CounterIn applies the In predicate on the "counter" field.
func CounterIn(vs ...int64) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCounter), v...))
	})
}

// CounterNotIn applies the NotIn predicate on the "counter" field.
func CounterNotIn(vs ...int64) predicate.Model {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Model(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCounter), v...))
	})
}

// CounterGT applies the GT predicate on the "counter" field.
func CounterGT(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCounter), v))
	})
}

// CounterGTE applies the GTE predicate on the "counter" field.
func CounterGTE(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCounter), v))
	})
}

// CounterLT applies the LT predicate on the "counter" field.
func CounterLT(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCounter), v))
	})
}

// CounterLTE applies the LTE predicate on the "counter" field.
func CounterLTE(v int64) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCounter), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Model) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Model) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
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
func Not(p predicate.Model) predicate.Model {
	return predicate.Model(func(s *sql.Selector) {
		p(s.Not())
	})
}