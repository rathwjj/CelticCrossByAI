package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Card struct {
	Name    string
	Meaning string
}

func main() {
	deck := []Card{
		{"The Fool", "New beginnings, innocence, spontaneity."},
		{"The Magician", "Manifestation, resourcefulness, power."},
		{"The High Priestess", "Intuition, sacred knowledge, subconscious."},
		{"The Empress", "Femininity, beauty, nature, abundance."},
		{"The Emperor", "Authority, establishment, structure."},
		{"The Hierophant", "Spiritual wisdom, religious beliefs, tradition."},
		{"The Lovers", "Love, harmony, relationships, choices."},
		{"The Chariot", "Control, willpower, success, action."},
		{"Strength", "Strength, courage, persuasion, influence."},
		{"The Hermit", "Soul-searching, introspection, being alone."},
		{"Wheel of Fortune", "Good luck, karma, life cycles, destiny."},
		{"Justice", "Justice, fairness, truth, cause and effect."},
		{"The Hanged Man", "Pause, surrender, letting go, new perspectives."},
		{"Death", "Endings, change, transformation, transition."},
		{"Temperance", "Balance, moderation, patience, purpose."},
		{"The Devil", "Shadow self, attachment, addiction, restriction."},
		{"The Tower", "Sudden change, upheaval, chaos, revelation."},
		{"The Star", "Hope, faith, purpose, renewal, spirituality."},
		{"The Moon", "Illusion, fear, anxiety, subconscious, intuition."},
		{"The Sun", "Positivity, fun, warmth, success, vitality."},
		{"Judgement", "Judgement, rebirth, inner calling, absolution."},
		{"The World", "Completion, integration, accomplishment, travel."},
		// For brevity, using generic placeholders for Minor Arcana. 
		// You can expand these similarly to the Major Arcana above.
		{"Ace of Wands", "Inspiration, new opportunities, growth."},
		{"Two of Wands", "Future planning, progress, decisions."},
		{"Three of Wands", "Expansion, foresight, overseas opportunities."},
		{"Four of Wands", "Celebration, joy, relaxation, homecoming."},
		{"Five of Wands", "Conflict, disagreements, competition."},
		{"Six of Wands", "Success, public recognition, victory."},
		{"Seven of Wands", "Challenge, competition, protection."},
		{"Eight of Wands", "Speed, action, air travel, movement."},
		{"Nine of Wands", "Resilience, courage, persistence."},
		{"Ten of Wands", "Burden, extra responsibility, hard work."},
		{"Page of Wands", "Discovery, enthusiasm, exploration."},
		{"Knight of Wands", "Energy, passion, inspired action, adventure."},
		{"Queen of Wands", "Courage, confidence, independence, social butterfly."},
		{"King of Wands", "Natural-born leader, vision, entrepreneur."},
		// Add remaining Cups, Swords, and Pentacles following the same pattern.
		{"Ace of Swords", "Breakthroughs, mental clarity, success."},
		{"Two of Swords", "Difficult choices, indecision, stalemate."},
		{"Three of Swords", "Heartbreak, emotional pain, sorrow."},
		{"Four of Swords", "Rest, recovery, meditation, introspection."},
		{"Five of Swords", "Conflict, tension, winning at all costs."},
		{"Six of Swords", "Transition, moving on, necessary change."},
		{"Seven of Swords", "Deception, strategy, acting alone."},
		{"Eight of Swords", "Restriction, imprisonment, victim mentality."},
		{"Nine of Swords", "Anxiety, nightmares, fear, depression."},
		{"Ten of Swords", "Betrayal, hitting rock bottom, endings."},
		{"Page of Swords", "Curiosity, restlessness, mental energy."},
		{"Knight of Swords", "Ambition, drive, fast-paced action."},
		{"Queen of Swords", "Clarity, independence, direct communication."},
		{"King of Swords", "Mental discipline, authority, truth."}
		{"Ace of Cups", "New feelings, creativity, spirituality."},
		{"Two of Cups", "Unity, partnership, mutual attraction."},
		{"Three of Cups", "Celebration, friendship, community."},
		{"Four of Cups", "Apathy, contemplation, disconnectedness."},
		{"Five of Cups", "Regret, failure, focusing on the loss."},
		{"Six of Cups", "Nostalgia, childhood memories, innocence."},
		{"Seven of Cups", "Choices, illusions, wishful thinking."},
		{"Eight of Cups", "Disillusionment, walking away, seeking truth."},
		{"Nine of Cups", "Contentment, satisfaction, gratitude."},
		{"Ten of Cups", "Divine love, blissful relationships, harmony."},
		{"Page of Cups", "Creative opportunities, intuitive messages."},
		{"Knight of Cups", "Romance, charm, following the heart."},
		{"Queen of Cups", "Compassion, emotional security, intuition."},
		{"King of Cups", "Emotional balance, control, generosity."}
		{"Ace of Pentacles", "New financial opportunity, manifestation."},
		{"Two of Pentacles", "Balance, multitasking, adaptability."},
		{"Three of Pentacles", "Teamwork, collaboration, learning."},
		{"Four of Pentacles", "Security, conservation, frugality."},
		{"Five of Pentacles", "Financial loss, isolation, poverty."},
		{"Six of Pentacles", "Giving, receiving, sharing wealth."},
		{"Seven of Pentacles", "Investment, patience, long-term view."},
		{"Eight of Pentacles", "Apprenticeship, mastery, skill development."},
		{"Nine of Pentacles", "Abundance, luxury, self-sufficiency."},
		{"Ten of Pentacles", "Legacy, inheritance, long-term security."},
		{"Page of Pentacles", "Manifestation, financial opportunity, skill."},
		{"Knight of Pentacles", "Hard work, productivity, routine."},
		{"Queen of Pentacles", "Nurturing, practical, providing."},
		{"King of Pentacles", "Wealth, business, leadership, security."}
	}

	positions := []string{
		"1. The Present (The Heart of the Matter)",
		"2. The Challenge (Obstacles)",
		"3. The Past (Foundation)",
		"4. The Future (What is approaching)",
		"5. Above (Conscious goals/Aspirations)",
		"6. Below (Subconscious/Underlying influences)",
		"7. Advice (How to handle the situation)",
		"8. External Influences (Environment/Others)",
		"9. Hopes and Fears",
		"10. Outcome (The final result)",
	}

	spread, err := drawCards(deck, 10)
	if err != nil {
		fmt.Println("Error drawing cards:", err)
		return
	}

	fmt.Println("--- CELTIC CROSS TAROT SPREAD ---")
	for i, card := range spread {
		fmt.Printf("\n%s\n", positions[i])
		fmt.Printf("CARD: %s\n", card.Name)
		fmt.Printf("MEANING: %s\n", card.Meaning)
	}
	fmt.Println("\n--- END OF READING ---")
}

func drawCards(deck []Card, count int) ([]Card, error) {
	selected := make([]Card, 0, count)
	// Create a copy to avoid mutating a global state if this were a larger app
	tempDeck := append([]Card(nil), deck...)

	for i := 0; i < count; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(tempDeck))))
		if err != nil {
			return nil, err
		}
		
		n := idx.Int64()
		selected = append(selected, tempDeck[n])
		// Remove drawn card from temp deck to prevent duplicates
		tempDeck = append(tempDeck[:n], tempDeck[n+1:]...)
	}
	return selected, nil
}
