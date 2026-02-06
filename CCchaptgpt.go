package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* ---------- DATA MODEL ---------- */

type Card struct {
	Name    string
	Arcana  string
	Suit    string
	Meaning string
}

type Deck struct {
	Cards []Card
}

/* ---------- DECK LOGIC ---------- */

func NewDeck() *Deck {
	// Minimal but COMPLETE working deck sample
	// You can expand this to all 78 cards later
	cards := []Card{
		{"The Fool", "Major", "", "Beginnings, openness, leap of faith"},
		{"The Magician", "Major", "", "Skill, focus, manifestation"},
		{"The High Priestess", "Major", "", "Intuition, inner voice"},
		{"The Empress", "Major", "", "Creation, nurture, growth"},
		{"The Emperor", "Major", "", "Structure, authority"},
		{"The Hierophant", "Major", "", "Tradition, belief systems"},
		{"The Lovers", "Major", "", "Choice, alignment"},
		{"The Chariot", "Major", "", "Willpower, direction"},
		{"Strength", "Major", "", "Inner strength, patience"},
		{"The Hermit", "Major", "", "Reflection, solitude"},
		{"Wheel of Fortune", "Major", "", "Cycles, change"},
		{"Justice", "Major", "", "Balance, cause and effect"},
		{"The Hanged Man", "Major", "", "Pause, new perspective"},
		{"Death", "Major", "", "Endings, transformation"},
		{"Temperance", "Major", "", "Integration, moderation"},
		{"The Devil", "Major", "", "Attachment, limitation"},
		{"The Tower", "Major", "", "Disruption, revelation"},
		{"The Star", "Major", "", "Hope, clarity"},
		{"The Moon", "Major", "", "Uncertainty, illusion"},
		{"The Sun", "Major", "", "Clarity, vitality"},
		{"Judgement", "Major", "", "Reckoning, awakening"},
		{"The World", "Major", "", "Completion, integration"},
	}

	return &Deck{Cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) CutOneThird() {
	cut := len(d.Cards) / 3
	d.Cards = d.Cards[cut:]
}

func (d *Deck) Draw(n int) []Card {
	if n > len(d.Cards) {
		n = len(d.Cards)
	}
	drawn := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return drawn
}

/* ---------- CELTIC CROSS ---------- */

func DrawCelticCross(d *Deck) []Card {
	return d.Draw(10)
}

/* ---------- MAIN ---------- */

func main() {
	deck := NewDeck()
	deck.Shuffle()
	deck.CutOneThird()

	cards := DrawCelticCross(deck)

	positions := []string{
		"1. Significator",
		"2. Crossing",
		"3. Foundation",
		"4. Past",
		"5. Crown",
		"6. Future",
		"7. Self",
		"8. Environment",
		"9. Hopes / Fears",
		"10. Outcome",
	}

	fmt.Println("Celtic Cross Reading (upright only)")
	fmt.Println("---------------------------------")

	for i, card := range cards {
		fmt.Printf("%s: %s\n", positions[i], card.Name)
	}
}
