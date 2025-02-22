package bcbp

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

const GROUP_SEPARATOR = '\x1D' // Group Separator (ASCII 0x1D)

type BCBP struct {
	Legs []*Leg `json:"legs"`
}

func (b *BCBP) String() string {

	parts := make([]string, len(b.Legs))

	for idx, l := range b.Legs {
		parts[idx] = l.String()
	}

	return strings.Join(parts, string(GROUP_SEPARATOR))
}

func Unmarshal(raw string) (*BCBP, error) {

	if string(raw[0]) != "M" {
		return nil, fmt.Errorf("BCBP string must start with M")
	}

	str_count := string(raw[1])
	count, err := strconv.Atoi(str_count)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse M (leg) count '%s', %w", str_count, err)
	}

	parts := strings.Split(raw, string(GROUP_SEPARATOR))
	legs := make([]*Leg, len(parts))

	if count != len(legs) {
		slog.Warn("BCBP length does not match M count.", "expected", count, "have", len(legs))
		return nil, fmt.Errorf("M count mismatch and liberal parsing not implemented yet")
	}

	for idx, leg_raw := range parts {

		l, err := ParseLeg(leg_raw)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse leg at offset %d (%s), %w", idx, leg_raw, err)
		}

		legs[idx] = l
	}

	b := &BCBP{
		Legs: legs,
	}

	return b, nil
}

func Marshal(b *BCBP) string {
	return b.String()
}
