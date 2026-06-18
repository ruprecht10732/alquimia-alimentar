package content

import (
	"encoding/json"
	"fmt"
	"os"
)

// Content holds all UI text for a given language.
type Content struct {
	Lang        string      `json:"lang"`
	Site        Site        `json:"site"`
	Header      Header      `json:"header"`
	Hero        Hero        `json:"hero"`
	Stickers    Stickers    `json:"stickers"`
	Desafios    Desafios    `json:"desafios"`
	Metodo      Metodo      `json:"metodo"`
	Lead        Lead        `json:"lead"`
	Sobre       Sobre       `json:"sobre"`
	Jornada     Jornada     `json:"jornada"`
	Depoimentos Depoimentos `json:"depoimentos"`
	CTA         CTA         `json:"cta"`
	Footer      Footer      `json:"footer"`
	WhatsApp    WhatsApp    `json:"whatsapp"`
}

// Site holds global site metadata.
type Site struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Link is a reusable label + URL pair.
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// Header holds text for the topbar and navigation.
type Header struct {
	Topbar       string     `json:"topbar"`
	LogoSubtitle string     `json:"logoSubtitle"`
	Nav          []Link     `json:"nav"`
	CTALabel     string     `json:"ctaLabel"`
	MobileCTA    string     `json:"mobileCTA"`
	Aria         HeaderAria `json:"aria"`
}

// HeaderAria holds accessibility labels for the header.
type HeaderAria struct {
	MainNav   string `json:"mainNav"`
	MobileNav string `json:"mobileNav"`
	OpenMenu  string `json:"openMenu"`
	CloseMenu string `json:"closeMenu"`
}

// Hero holds the hero section content.
type Hero struct {
	Accent       string   `json:"accent"`
	Title        []string `json:"title"`
	Description  string   `json:"description"`
	PrimaryCTA   string   `json:"primaryCTA"`
	SecondaryCTA string   `json:"secondaryCTA"`
	Badge        Badge    `json:"badge"`
}

// Badge is a small circular badge.
type Badge struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// Stickers holds the floating sticker section.
type Stickers struct {
	Items []Sticker `json:"items"`
	Quote string    `json:"quote"`
}

// Sticker is a single floating pill.
type Sticker struct {
	Emoji    string `json:"emoji"`
	Label    string `json:"label"`
	Variant  string `json:"variant"`
	Rotation string `json:"rotation"`
}

// Desafios holds the "expat challenges" section.
type Desafios struct {
	Accent      string `json:"accent"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cards       []Card `json:"cards"`
}

// Card is a reusable content card.
type Card struct {
	Emoji       string `json:"emoji"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Variant     string `json:"variant"`
}

// Metodo holds the method/pillars section.
type Metodo struct {
	Accent      string   `json:"accent"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Quote       string   `json:"quote"`
	Pillars     []Pillar `json:"pillars"`
}

// Pillar is a single method pillar.
type Pillar struct {
	Number      string `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Variant     string `json:"variant"`
}

// Lead holds the lead magnet section.
type Lead struct {
	Badge       string `json:"badge"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Placeholder string `json:"placeholder"`
	Button      string `json:"button"`
}

// Sobre holds the about section.
type Sobre struct {
	Accent     string   `json:"accent"`
	Title      []string `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Quote      string   `json:"quote"`
}

// Jornada holds the journey phases section.
type Jornada struct {
	Accent string  `json:"accent"`
	Title  string  `json:"title"`
	Phases []Phase `json:"phases"`
}

// Phase is a single journey phase.
type Phase struct {
	Phase       string `json:"phase"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Variant     string `json:"variant"`
}

// Depoimentos holds the testimonials section.
type Depoimentos struct {
	Accent       string        `json:"accent"`
	Title        string        `json:"title"`
	Testimonials []Testimonial `json:"testimonials"`
}

// Testimonial is a single customer testimonial.
type Testimonial struct {
	Text     string `json:"text"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Rotation string `json:"rotation"`
}

// CTA holds the final call-to-action section.
type CTA struct {
	Accent      string `json:"accent"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Button      string `json:"button"`
	Privacy     string `json:"privacy"`
}

// Footer holds the footer content.
type Footer struct {
	Tagline         string    `json:"tagline"`
	DigitalOffice   LinkGroup `json:"digitalOffice"`
	DailyConnection LinkGroup `json:"dailyConnection"`
	Copyright       string    `json:"copyright"`
	MadeWith        string    `json:"madeWith"`
}

// LinkGroup is a reusable footer link group.
type LinkGroup struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Link  Link   `json:"link"`
}

// WhatsApp holds the floating button content.
type WhatsApp struct {
	Phone     string `json:"phone"`
	AriaLabel string `json:"ariaLabel"`
}

// Load reads the content file for the given language.
func Load(lang string) (*Content, error) {
	path := fmt.Sprintf("web/content/%s.json", lang)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read content file %s: %w", path, err)
	}

	var content Content
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, fmt.Errorf("failed to parse content file %s: %w", path, err)
	}

	return &content, nil
}
