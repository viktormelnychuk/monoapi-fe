MIGRATION_FILE = $(shell date +"migrations/%Y%m%d%H%M%S-$(name).sql")

new-migration: ## New migration (make name=add-some-table new-migration)
	touch $(MIGRATION_FILE)
	echo "-- +migrate Up\n\n-- +migrate Down" >> $(MIGRATION_FILE)
