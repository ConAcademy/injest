// jester package
// injest - another ConAcademy joint
// Copyright (c) 2023 Neomantra BV

package jester

import "strings"

////////////////////////////////////////////////////////////////////////////

type Jester interface {
	MakeJoke(input string) string
}

////////////////////////////////////////////////////////////////////////////

type KnockKnockJoke struct {
}

func (j KnockKnockJoke) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Knock knock\n")
	sb.WriteString("Who's there?\n")
	sb.WriteString("Your data.\n")
	sb.WriteString("Your data, who?\n")
	sb.WriteString("Your data is not a joke:\n") // CoPilot came up with that so it stays
	sb.WriteString(input)
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type YoMammaJoke struct {
}

func (j YoMammaJoke) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Yo momma so fat that ") // HAHA Copilot: when she sat on a binary tree she flattened it to a linked list.\n")
	sb.WriteString(input)
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type RacistJoke struct {
}

func (j RacistJoke) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Enlighten yourself and read the CODE OF CONDUCT.  Users should be nice too.\n")
	sb.WriteString("https://github.com/ConAcademy/injest/blob/main/CODE_OF_CONDUCT.md\n")
	return sb.String()
}
