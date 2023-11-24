package workplace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Notifier sends messages to a workplace group or chat.
type Notifier struct {
	Token string
	URL   string
}

// Post posts a message to a workplace group.
func (n *Notifier) Post(ctx context.Context, groupID, message string) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(map[string]string{
		"message":    message,
		"formatting": "MARKDOWN",
	}); err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s/feed", n.URL, groupID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.Token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("workplace error: status %d, message: %q", resp.StatusCode, body)
	}
	return nil
}
