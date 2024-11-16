package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var defaultWords = &words{
	articles:   []string{"the", "a"},
	adjectives: []string{"happy", "blue", "small", "tall", "quick", "slow", "bright", "dark", "beautiful", "ugly", "strong", "weak", "smooth", "rough", "cold", "hot", "loud", "quiet", "soft", "hard", "young", "thin", "fat", "friendly", "clean", "dirty", "rich", "poor", "silly", "funny", "serious", "boring", "generous", "selfish", "peaceful", "noisy", "creative", "clumsy", "careful", "reckless", "dishonest"},
	adverbs:    []string{"quickly", "slowly", "happily", "sadly", "loudly", "quietly", "carefully", "recklessly", "bravely", "easily", "hardly", "always", "never", "sometimes", "often", "rarely", "frequently", "beautifully", "awkwardly", "smoothly", "roughly", "brightly", "darkly", "warmly", "coldly", "lazily", "seriously", "gently", "noisily", "enthusiastically", "intelligently", "clumsily", "gracefully", "mysteriously", "angrily", "calmly", "politely", "rudely", "generously"},
	nouns:      []string{"dog", "cat", "bird", "fish", "horse", "rabbit", "tiger", "lion", "giraffe", "koala", "kangaroo", "whale", "shark", "panda", "monkey", "snake", "frog", "octopus", "butterfly", "bee", "ant", "spider", "caterpillar", "squirrel", "dolphin", "zebra", "deer", "cow", "goat", "sheep", "chicken", "duck", "pig", "mouse", "lizard", "bat", "ferret", "parrot", "flamingo", "penguin"},
	verbs:      []string{"walked", "watched", "found", "worked", "cleaned", "helped", "discovered", "received", "responded", "attended", "visited", "joined", "managed", "pushed", "shoved", "hit", "smacked", "knocked", "scratched", "slapped", "tugged", "tripped", "jostled", "charged", "chased", "caught", "defended", "blocked", "overpowered", "blasted", "crashed", "spooked", "stumbled", "frightened", "hurt", "intercepted", "repelled", "scuffled", "disturbed", "meditated", "calmed", "embraced", "hugged", "smiled", "reassured", "comforted", "relaxed", "welcomed", "nurtured", "helped", "cherished", "planted", "invited", "healed", "protected", "served", "celebrated", "supported", "rested", "shared", "taught", "rescued", "advised", "forgave", "honored", "encouraged", "soothed"},
}

type words struct {
	articles   []string
	adjectives []string
	adverbs    []string
	nouns      []string
	verbs      []string
}

type sentenceBuilder struct {
	articleFirst   string
	adjectiveFirst string
	nounFirst      string
	adverb         string
	verb           string
	articleLast    string
	adjectiveLast  string
	nounLast       string
}

func (w *words) newSentenceBuilder() *sentenceBuilder {
	return &sentenceBuilder{
		articleFirst:   w.articles[rand.Intn(len(w.articles))],
		adjectiveFirst: w.adjectives[rand.Intn(len(w.adjectives))],
		nounFirst:      w.nouns[rand.Intn(len(w.nouns))],
		adverb:         w.adverbs[rand.Intn(len(w.adverbs))],
		verb:           w.verbs[rand.Intn(len(w.verbs))],
		articleLast:    w.articles[rand.Intn(len(w.articles))],
		adjectiveLast:  w.adjectives[rand.Intn(len(w.adjectives))],
		nounLast:       w.nouns[rand.Intn(len(w.nouns))],
	}
}

func (b *sentenceBuilder) getSentence(numWords int) (string, error) {
	switch numWords {
	case 1:
		return b.nounFirst, nil
	case 2:
		return fmt.Sprintf("%s %s", b.articleFirst, b.nounFirst), nil
	case 3:
		return fmt.Sprintf("%s %s %s",
			b.articleFirst,
			b.adjectiveFirst,
			b.nounFirst,
		), nil
	case 4:
		return fmt.Sprintf("%s %s %s %s",
			b.articleFirst,
			b.adjectiveFirst,
			b.nounFirst,
			b.verb,
		), nil
	case 5:
		return fmt.Sprintf("%s %s %s %s %s",
			b.articleFirst,
			b.adjectiveFirst,
			b.nounFirst,
			b.adverb,
			b.verb,
		), nil
	case 6:
		return fmt.Sprintf("%s %s %s %s %s %s",
			b.articleFirst,
			b.adjectiveFirst,
			b.nounFirst,
			b.verb,
			b.articleLast,
			b.nounLast,
		), nil
	case 7:
		return fmt.Sprintf("%s %s %s %s %s %s %s",
			b.articleFirst,
			b.adjectiveFirst,
			b.nounFirst,
			b.adverb,
			b.verb,
			b.articleLast,
			b.nounLast,
		), nil
	case 8:
		return fmt.Sprintf("%s %s %s %s %s %s %s %s",
			b.articleFirst,
			b.adjectiveFirst,
			b.nounFirst,
			b.adverb,
			b.verb,
			b.articleLast,
			b.adjectiveLast,
			b.nounLast,
		), nil
	default:
		return "", errors.New("Number of words per sentence must be [1-8]")
	}
}
