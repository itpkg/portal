package cms

type Dao struct {
}

//-----------------------------------------------------------------------------
func (p *Engine) Seed() error {
	return nil
}
func (p *Engine) Migrate() {
	db := p.Db
	db.AutoMigrate(&Article{}, &Tag{}, &Comment{})
	db.Model(&Article{}).AddIndex("idx_articles_title", "title")
}
