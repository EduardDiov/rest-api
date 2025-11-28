package positions

import (
	"net/http"

	"diov.local/krasivo/internal/components/positions"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	component := positions.PositionsList()
	component.Render(r.Context(), w)
}
