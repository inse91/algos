package _929_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		emails []string
		res    int
	}{
		{
			name: "1",
			emails: []string{
				"test.email+alex@leetcode.com",
				"test.e.mail+bob.cathy@leetcode.com",
				"testemail+david@lee.tcode.com",
			},
			res: 2,
		},
		{
			name: "2",
			emails: []string{
				"a@leetcode.com",
				"b@leetcode.com",
				"c@leetcode.com",
			},
			res: 3,
		},
		{
			name:   "184",
			emails: []string{"test.email+alex@leetcode.com", "test.email@leetcode.com"},
			res:    1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, numUniqueEmails(c.emails))
		})
	}
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{}, len(emails)/2)

	for _, email := range emails {
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			panic("invalid email")
		}

		domainName := parts[1]
		name := parts[0]

		var sb strings.Builder

		for _, r := range name {
			if r == '+' {
				break
			}
			if r == '.' {
				continue
			}
			sb.WriteRune(r)
		}
		name = sb.String()

		m[name+"@"+domainName] = struct{}{}
	}

	return len(m)
}
