package seeddata

import (
	"context"
	"shrektionary_api/ent"
	"shrektionary_api/ent/migrate"
)

func Test(ctx context.Context, client *ent.Client) error {
	// Reset the database schema
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return err
	}

	// Delete records from the "Word" table
	_, errDelete := client.Word.
		Delete().
		Exec(ctx)
	if errDelete != nil {
		return errDelete
	}

	// Create definitions
	definitions_shrekt := []*ent.Definition{
		client.Definition.
			Create().
			SetDescription("Definition 1 for Shrek").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 2 for Shrek").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 3 for Shrek").
			SaveX(ctx),
	}

	definitions_bot := []*ent.Definition{
		client.Definition.
			Create().
			SetDescription("Definition 1 for Bot").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 2 for Bot").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 3 for Bot").
			SaveX(ctx),
	}

	// Create words and associate definitions
	words := []*ent.WordCreate{
		client.Word.
			Create().
			SetDescription("Shrek").
			AddDefinitions(
				definitions_shrekt[0],
				definitions_shrekt[1],
				definitions_shrekt[2],
			),
		client.Word.
			Create().
			SetDescription("Bot").
			AddDefinitions(
				definitions_bot[0],
				definitions_bot[1],
				definitions_bot[2],
			),
	}

	_, err := client.Word.CreateBulk(words...).Save(ctx)
	return err
}
