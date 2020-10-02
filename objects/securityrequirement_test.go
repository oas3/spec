package objects_test

import (
	"testing"

	"github.com/oas3/spec/objects"
)

func eqSecurityRequirement(t *testing.T, sr1, sr2 objects.SecurityRequirement) {
	eqInt(t, len(sr1), len(sr2))
	for k, r1 := range sr1 {
		r2 := sr2[k]
		eqInt(t, len(r1), len(r2))
		for i, s1 := range r1 {
			s2 := r2[i]
			eq(t, s1, s2)
		}
	}
}
