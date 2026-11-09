package tickets

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Ticket struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Closed bool   `json:"closed"`
}

// String is how one ticket looks as a line of text.
func (t Ticket) String() string {
	status := "open"
	if t.Closed {
		status = "closed"
	}
	return fmt.Sprintf("%d. %s [%s]", t.ID, t.Title, status)
}

// Close needs a pointer receiver so it changes the real ticket, not a copy.
func (t *Ticket) Close() {
	t.Closed = true
}

// Load reads the file. A missing file just means "no tickets yet".
func Load(path string) ([]Ticket, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Ticket{}, nil
		}
		return nil, err
	}

	var list []Ticket
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func Save(path string, list []Ticket) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Add returns the longer slice and the new ticket's ID.
// No delete exists, so len+1 never repeats an ID.
func Add(list []Ticket, title string) ([]Ticket, int) {
	id := len(list) + 1
	list = append(list, Ticket{ID: id, Title: title})
	return list, id
}

// CloseTicket loops by index so it reaches the real ticket, not a range copy.
func CloseTicket(list []Ticket, id int) error {
	for i := range list {
		if list[i].ID == id {
			if list[i].Closed {
				return fmt.Errorf("ticket %d is already closed", id)
			}
			list[i].Close()
			return nil
		}
	}
	return fmt.Errorf("ticket %d not found", id)
}

// Filter picks which tickets to show: open only, or everything.
func Filter(list []Ticket, all bool) []Ticket {
	shown := []Ticket{}
	for _, t := range list {
		if all || !t.Closed {
			shown = append(shown, t)
		}
	}
	return shown
}
