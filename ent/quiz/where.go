// Code generated by entc, DO NOT EDIT.

package quiz

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/duiztapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// URLImg applies equality check predicate on the "url_img" field. It's identical to URLImgEQ.
func URLImg(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLImg), v))
	})
}

// URLImgEQ applies the EQ predicate on the "url_img" field.
func URLImgEQ(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLImg), v))
	})
}

// URLImgNEQ applies the NEQ predicate on the "url_img" field.
func URLImgNEQ(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURLImg), v))
	})
}

// URLImgIn applies the In predicate on the "url_img" field.
func URLImgIn(vs ...string) predicate.Quiz {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Quiz(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURLImg), v...))
	})
}

// URLImgNotIn applies the NotIn predicate on the "url_img" field.
func URLImgNotIn(vs ...string) predicate.Quiz {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Quiz(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURLImg), v...))
	})
}

// URLImgGT applies the GT predicate on the "url_img" field.
func URLImgGT(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURLImg), v))
	})
}

// URLImgGTE applies the GTE predicate on the "url_img" field.
func URLImgGTE(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURLImg), v))
	})
}

// URLImgLT applies the LT predicate on the "url_img" field.
func URLImgLT(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURLImg), v))
	})
}

// URLImgLTE applies the LTE predicate on the "url_img" field.
func URLImgLTE(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURLImg), v))
	})
}

// URLImgContains applies the Contains predicate on the "url_img" field.
func URLImgContains(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURLImg), v))
	})
}

// URLImgHasPrefix applies the HasPrefix predicate on the "url_img" field.
func URLImgHasPrefix(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURLImg), v))
	})
}

// URLImgHasSuffix applies the HasSuffix predicate on the "url_img" field.
func URLImgHasSuffix(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURLImg), v))
	})
}

// URLImgEqualFold applies the EqualFold predicate on the "url_img" field.
func URLImgEqualFold(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURLImg), v))
	})
}

// URLImgContainsFold applies the ContainsFold predicate on the "url_img" field.
func URLImgContainsFold(v string) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURLImg), v))
	})
}

// HasQuestions applies the HasEdge predicate on the "questions" edge.
func HasQuestions() predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuestionsTable, QuestionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionsWith applies the HasEdge predicate on the "questions" edge with a given conditions (other predicates).
func HasQuestionsWith(preds ...predicate.Question) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuestionsTable, QuestionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLangs applies the HasEdge predicate on the "langs" edge.
func HasLangs() predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LangsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LangsTable, LangsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLangsWith applies the HasEdge predicate on the "langs" edge with a given conditions (other predicates).
func HasLangsWith(preds ...predicate.QuizLangs) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LangsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LangsTable, LangsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Quiz) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Quiz) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
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
func Not(p predicate.Quiz) predicate.Quiz {
	return predicate.Quiz(func(s *sql.Selector) {
		p(s.Not())
	})
}