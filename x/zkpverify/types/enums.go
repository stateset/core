package types

// RuleType defines the type of symbolic rule.
type RuleType string

const (
	RuleTypeImplication   RuleType = "implication"    // if A then B
	RuleTypeConjunction   RuleType = "conjunction"    // A and B and C
	RuleTypeDisjunction   RuleType = "disjunction"    // A or B or C
	RuleTypeNegation      RuleType = "negation"       // not A
	RuleTypeUniversal     RuleType = "universal"      // for all X: P(X)
	RuleTypeExistential   RuleType = "existential"    // exists X: P(X)
	RuleTypeEquality      RuleType = "equality"       // A == B
	RuleTypeInequality    RuleType = "inequality"     // A != B
	RuleTypeComparison    RuleType = "comparison"     // A < B, A > B, etc.
	RuleTypeSetMembership RuleType = "set_membership" // A in {x, y, z}
)
