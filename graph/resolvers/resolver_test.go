package resolvers

import (
	"context"
	"fmt"
	"groupinary/testutils"
	"net/http"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestNewSchema(t *testing.T) {
	// Open a test client and handle any errors.
	client, _, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create the executable schema using the NewSchema function.
	schema := NewSchema(client)

	// Assert that the returned schema is not nil.
	assert.NotNil(t, schema, "Expected schema to be not nil.")
	assert.NotNil(t, schema.Schema(), "Expected schema.Schema() to be not nil.")
}

func TestComplexity(t *testing.T) {
	// Open a test client and handle any errors.
	client, _, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create the executable schema using the NewSchema function.
	schema := NewSchema(client)

	// Create a mock args map.
	args := map[string]interface{}{
		"arg1": "value1",
		"arg2": 123,
		// Add more arguments as needed.
	}

	// Call the Complexity method with the mock args.
	intValue, boolValue := schema.Complexity("Definition", "createTime", 0, args)

	// Assert that the Complexity method behaves as expected.
	assert.NotNil(t, intValue, "Expected non-zero complexity.")
	assert.NotNil(t, boolValue, "Expected ok = true.")
}

func TestExec(t *testing.T) {
	// Open a test client and handle any errors.
	client, _, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create the executable schema using the NewSchema function.
	schema := NewSchema(client)

	opContext := &graphql.OperationContext{
		RawQuery:      "mock query",
		Variables:     map[string]interface{}{"key": "value"},
		OperationName: "mock operation",
		Doc:           &ast.QueryDocument{}, // You need to assign a proper QueryDocument here
		Headers:       http.Header{"Content-Type": []string{"application/json"}},

		// You may need to define a proper OperationDefinition and assign it to Operation
		Operation: &ast.OperationDefinition{},

		// This is an optional field, you can set it to false if you don't want to disable introspection
		DisableIntrospection: true,

		// You could define a custom RecoverFunc
		RecoverFunc: func(ctx context.Context, err interface{}) error {
			return fmt.Errorf("unexpected error: %v", err)
		},
		ResolverMiddleware:     nil,
		RootResolverMiddleware: nil,

		// Stats can be nil or you can initialize a new Stats object
		Stats: graphql.Stats{},
	}

	ctx := graphql.WithOperationContext(context.Background(), opContext)

	// Assert that the Exec method behaves as expected.
	assert.NotNil(t, schema.Exec(ctx), "Expected Exec result to be not nil.")
}
