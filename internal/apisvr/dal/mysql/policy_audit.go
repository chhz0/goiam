package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type policyAudit struct {
	db *gorm.DB
}

func (p *policyAudit) ClearOutdatedAudit(ctx context.Context, maxReserveDays int) (int64, error) {
	data := time.Now().AddDate(0, 0, -maxReserveDays).Format("2006-01-02 15:04:05")

	d := p.db.Exec("delete from policy_audit where deleted_at < ?", data)

	return d.RowsAffected, d.Error
}

func newPolicyAudits(ds *dbStore) *policyAudit {
	return &policyAudit{db: ds.db}
}
