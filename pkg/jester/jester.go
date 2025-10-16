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

type KnockKnockJester struct {
}

func (j KnockKnockJester) MakeJoke(input string) string {
	// Formulaic joke experiment: classic knock-knock structure with <A> and <B>
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

type YoMammaJester struct {
}

func (j YoMammaJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Yo momma so fat that ") // HAHA Copilot: when she sat on a binary tree she flattened it to a linked list.\n")
	sb.WriteString(input)
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type RacistJester struct {
}

func (j RacistJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Enlighten yourself and read the CODE OF CONDUCT.  Users should be nice too.\n")
	sb.WriteString("https://github.com/ConAcademy/injest/blob/main/CODE_OF_CONDUCT.md\n")
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type OfficeJester struct {
}

func (j OfficeJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString(input)
	sb.WriteString("\n  ... that's what she said!\n")
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type DadJester struct {
}

// Formulaic joke experiment: subject-driven dad joke one-liner
func (j DadJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("I would tell you a joke about ")
	sb.WriteString(input)
	sb.WriteString(", but it's too cheesy.\n")
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type RiddleJester struct {
}

// Formulaic joke experiment: simple riddle-style prompt and punchline
func (j RiddleJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Riddle me this: When is ")
	sb.WriteString(input)
	sb.WriteString(" like a door? When it's a-jar!\n")
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type CrossroadJester struct {
}

// Formulaic joke experiment: classic "Why did the <subject> cross the road?"
func (j CrossroadJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("Why did the ")
	sb.WriteString(input)
	sb.WriteString(" cross the road?\n")
	sb.WriteString("To get to the other side!\n")
	return sb.String()
}

////////////////////////////////////////////////////////////////////////////

type WhatDoYouCallJester struct {
}

// Formulaic joke experiment: "What do you call <thing>?" naming gag with a few clean punchlines
func (j WhatDoYouCallJester) MakeJoke(input string) string {
	var sb strings.Builder
	sb.WriteString("What do you call ")
	sb.WriteString(input)
	sb.WriteString("?\n")
	// Minor dynamicism: rotate among a few wholesome punchlines based on input length
	switch len(strings.TrimSpace(input)) % 4 {
	case 0:
		sb.WriteString("A classic!\n")
	case 1:
		sb.WriteString("A real gem!\n")
	case 2:
		sb.WriteString("A good one!\n")
	default:
		sb.WriteString("A keeper!\n")
	}
	return sb.String()
}
