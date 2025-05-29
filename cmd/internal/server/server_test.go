package server

import "testing"

func TestServer(t *testing.T) {
	t.Run("Overwrite a decision succeeds", func(t *testing.T) {
		// Given I am user 1 and I have an existing 'pass' on user 2
		// When I submit a 'like' on user 2
		// Then the new decision overwrites the original

		t.Errorf("not implemented")
	})

	t.Run("Get new liked you returns successfully - no pagination token", func(t *testing.T) {
		// Given I am user 1
		// When I call `ListLikedYou` with no pagination token in the request
		// Then the correct list of likes is returned
		// And a pagination token is returned if relevant

		t.Errorf("not implemented")
	})

	t.Run("Get liked you returns successfully - no pagination token", func(t *testing.T) {
		// Given I am user 1
		// When I call `ListNewLikedYou` with no pagination token in the request
		// Then the correct list of likes is returned
		// And a pagination token is returned if relevant

		t.Errorf("not implemented")
	})

	t.Run("Get liked you handles pagination", func(t *testing.T) {
		// Given I am user 1
		// When I call `ListLikedYou` with a token in the request
		// Then the correct list of likes is returned
		// And a pagination token is returned if present

		t.Errorf("not implemented")
	})

	t.Run("Liking yourself returns an error", func(t *testing.T) {
		// Given I am user 1
		// When I submit a 'like' on user 1
		// Then a sensible error is returned

		t.Errorf("not implemented")
	})

	t.Run("Put decision returns 'true' mutual like", func(t *testing.T) {
		// Given I am user 1
		// And I have an existing like from user 2
		// When I submit a 'like' on user 2
		// Then `true` is returned for `mutual_like` in response

		t.Errorf("not implemented")
	})

	t.Run("Put decision returns 'false' mutual like", func(t *testing.T) {
		// Given I am user 1
		// And I have an existing like from user 2
		// When I submit a 'like' on user 2
		// Then `false` is returned for `mutual_like` in response

		t.Errorf("not implemented")
	})
}
