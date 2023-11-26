package seeddata

import (
	"context"
	"groupinary/ent"
	"groupinary/ent/migrate"
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

	// Delete records from the "Word" table
	_, errDelete = client.Definition.
		Delete().
		Exec(ctx)
	if errDelete != nil {
		return errDelete
	}

	// Delete records from the "Word" table
	_, errDelete = client.Group.
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

	descendants_shrekt := []*ent.Word{
		client.Word.
			Create().
			SetDescription("shrektard").
			SaveX(ctx),
		client.Word.
			Create().
			SetDescription("shrektdroid").
			SaveX(ctx),
		client.Word.
			Create().
			SetDescription("shrekt lord").
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
	// words := []*ent.WordCreate{
	// 	client.Word.
	// 		Create().
	// 		SetDescription("Shrek").
	// 		AddDefinitions(
	// 			definitions_shrekt[0],
	// 			definitions_shrekt[1],
	// 			definitions_shrekt[2],
	// 		),
	// 	client.Word.
	// 		Create().
	// 		SetDescription("Bot").
	// 		AddDefinitions(
	// 			definitions_bot[0],
	// 			definitions_bot[1],
	// 			definitions_bot[2],
	// 		),
	// }

	// Creates group
	group := []*ent.GroupCreate{
		client.Group.Create().
			SetDescription("the Boys").
			AddWords(
				client.Word.
					Create().
					SetDescription("Shrek").
					AddDefinitions(
						definitions_shrekt[0],
						definitions_shrekt[1],
						definitions_shrekt[2],
					).
					AddDescendants(
						descendants_shrekt[0],
						descendants_shrekt[1],
						descendants_shrekt[2],
					).
					SaveX(ctx),
				client.Word.Create().
					SetDescription("Bot").
					AddDefinitions(
						definitions_bot[0],
						definitions_bot[1],
						definitions_bot[2],
					).
					SaveX(ctx),
			),
	}

	_, err := client.Group.CreateBulk(group...).Save(ctx)
	return err
}
