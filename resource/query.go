package resource

import (
	"context"
	"net/http"
	"strconv"

	"github.com/arrowsjs/pagination-server/db"
	"github.com/efritz/chevron"
	"github.com/efritz/nacelle"
	"github.com/efritz/response"
)

type QueryResource struct {
	*chevron.EmptySpec
	DB *db.DB `service:"db"`
}

func (r *QueryResource) Get(ctx context.Context, req *http.Request, logger nacelle.Logger) response.Response {
	return response.JSON(r.DB.Search(getQuery(req), getPage(req), 25))
}

func getQuery(req *http.Request) string {
	return req.URL.Query().Get("q")
}

func getPage(req *http.Request) int {
	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		return 1
	}

	return page
}
