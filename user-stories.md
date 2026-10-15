# Todo API — User Stories

## US-01 — Create a todo
As a user, I want to create a todo so that I can record a task I need to do.

Acceptance criteria:
- The request accepts a name and an optional status.
- The name is required and cannot be blank.
- If status is not provided, it is set to `pending`.
- A created todo is returned with its ID, name, and status.
- Invalid input returns `400 Bad Request`.

## US-02 — Get a todo
As a user, I want to get a todo by its ID so that I can view its details.

Acceptance criteria:
- If the todo exists, return it with `200 OK`.
- If the todo does not exist, return `404 Not Found`.

## US-03 — List todos
As a user, I want to list todos so that I can see the tasks I have recorded.

Acceptance criteria:
- Return a list of todos.
- If there are no todos, return an empty list.
- Support a cursor to request the next page of results.
- Return `200 OK` when the request succeeds.

## US-04 — Update a todo
As a user, I want to update a todo so that I can change its name or status.

Acceptance criteria:
- The request accepts the updated name and status.
- The updated todo is returned with `200 OK`.
- If the todo does not exist, return `404 Not Found`.
- Invalid input returns `400 Bad Request`.

## US-05 — Delete a todo
As a user, I want to delete a todo so that I can remove a task I no longer need.

Acceptance criteria:
- A todo can be deleted using its ID.
- Return `204 No Content` when it is deleted successfully.
- If the todo does not exist, return `404 Not Found`.

## US-06 — Get clear error responses
As a user, I want clear error responses so that I understand when a request fails.

Acceptance criteria:
- Invalid input returns `400 Bad Request`.
- A todo that cannot be found returns `404 Not Found`.
- The error response includes a short message explaining the problem.
