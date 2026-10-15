# Todo API

A simple REST API to create, view, list, update, and delete todos. OpenAPI version **1.0.0**. Default server: `http://localhost:8080`.

Status values: `pending`, `in_progress`, `completed`. New todos default to `pending` when status is omitted.

## User stories

| ID | Story |
| --- | --- |
| US-01 | Create a todo (name required; status optional) |
| US-02 | Get a todo by ID |
| US-03 | List todos with cursor pagination |
| US-04 | Update a todo name and status |
| US-05 | Delete a todo by ID |
| US-06 | Return clear `400` / `404` errors with a message |

See [user-stories.md](user-stories.md) for acceptance criteria.

## API

JSON request and response bodies. Errors use `{ "message": "..." }`.

A todo object:

```json
{
  "id": 1,
  "name": "Finish Go assignment",
  "status": "pending",
  "created_at": "2026-10-15T09:30:00Z",
  "updated_at": "2026-10-15T09:30:00Z"
}
```

### `POST /todos` — create

Body: `{ "name": "Finish Go assignment", "status": "pending" }`  
`name` is required. `status` is optional (default `pending`).

| Status | Meaning |
| --- | --- |
| `201` | Created todo |
| `400` | Missing name or invalid status |

### `GET /todos` — list

Query: `limit` (default `10`, min `1`, max `50`), `cursor` (from a previous `next_cursor`).

Response: `{ "todos": [...], "next_cursor": "..." }` — `todos` may be empty; `next_cursor` may be `null`.

| Status | Meaning |
| --- | --- |
| `200` | Page of todos |
| `400` | Invalid limit or cursor |

### `GET /todos/{id}` — get

`id` is a positive integer.

| Status | Meaning |
| --- | --- |
| `200` | Todo found |
| `400` | Invalid ID |
| `404` | Todo not found |

### `PUT /todos/{id}` — update

Body: `{ "name": "Finish Go assignment", "status": "completed" }`  
Both `name` and `status` are required.

| Status | Meaning |
| --- | --- |
| `200` | Updated todo |
| `400` | Invalid input |
| `404` | Todo not found |

### `DELETE /todos/{id}` — delete

| Status | Meaning |
| --- | --- |
| `204` | Deleted (empty body) |
| `400` | Invalid ID |
| `404` | Todo not found |

Full contract: [swagger.yaml](swagger.yaml).

## Database

MySQL. One table, `todos`. Diagram: [schema.dbml](schema.dbml).

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigint` | Primary key, auto-increment |
| `name` | `varchar(255)` | Required |
| `status` | `varchar(20)` | Required, default `pending` |
| `created_at` | `timestamp` | Default `CURRENT_TIMESTAMP` |
| `updated_at` | `timestamp` | Default `CURRENT_TIMESTAMP` |

## Tests

[todo_test_stubs.go](todo_test_stubs.go) is a starter Go test file. Wire `newTestHandler` to the real HTTP handler before running tests.

Positive cases: create (status defaults to `pending`), get, list, update, delete (`204`).

Negative cases: missing name (`400`), unsupported status such as `started` (`400`), missing todo (`404`).
