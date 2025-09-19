package logging_test

import (
	"testing"

	"github.com/digiconvent/d9t/utils/logging"
	"github.com/google/uuid"
)

func TestTableFormatting(t *testing.T) {
	id, _ := uuid.NewV7()

	table := logging.NewTable([]string{"key", "value"})
	table.AddRow(id, id)

	first := id.String()[:4]
	last := id.String()[32:]
	expectedTable := `┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━┓
┃ key         ┃ value       ┃
┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━┫
┃ ` + first + `...` + last + ` ┃ ` + first + `...` + last + ` ┃
┗━━━━━━━━━━━━━┻━━━━━━━━━━━━━┛
`
	if table.Render() != expectedTable {
		t.Fatal("expected", expectedTable, "\ninstead got\n",
			table.Render())
	}
}
