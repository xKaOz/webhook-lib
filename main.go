package webhook_lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	URL      string  `json:"-"`
	Username string  `json:"username,omitempty"`
	Avatar   string  `json:"avatar_url,omitempty"`
	Content  string  `json:"content,omitempty"`
	Embeds   []*Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	URL         string       `json:"url,omitempty"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []*EmbedField `json:"fields,omitempty"`
}

type EmbedField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (w *Webhook) Send() error {
	jsonBody, err := json.Marshal(w)
	if err != nil {
		return err
	}

	resp, err := http.Post(w.URL, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	return nil
}
