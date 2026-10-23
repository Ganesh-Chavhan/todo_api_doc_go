# Todo API

Checklist-based Todo API, version **1.0.0**. Base URL: `http://localhost:8080/api/v1`.

A **todo** is a main task (for example “Update package”). It can have zero or more **checklist items** (for example “Check updated package info”). Checklist items belong to one todo, can be completed independently, and do not appear as separate todos on the dashboard.

Todo **status** (`To Do`, `In Progress`, `Completed`) is separate from **checklist progress**. Progress is calculated as completed item count / total item count and is not stored as its own column.

**Archive** and **recycle bin** are separate from status:

| State | Rule |
| --- | --- |
| Active | `archived_at` and `deleted_at` are both null |
| Archived | `archived_at` is set, `deleted_at` is null |
| Recycle bin | `deleted_at` is set (soft delete; checklist is kept) |

A todo cannot be archived and in the recycle bin at the same time. Recycle-bin todos must be restored before they can be unarchived. Permanent delete physically removes the todo and cascades to its checklist items.

Authentication and user accounts are not part of this contract.

**Implement against [`todo_api_openapi.yaml`](todo_api_openapi.yaml).** [`swagger.yaml`](swagger.yaml) is an older CRUD-only spec and does not match these stories.

## Files

| File | Role |
| --- | --- |
| [`todo_epics_user_stories_final.md`](todo_epics_user_stories_final.md) | Epics US-01–US-21 and acceptance criteria |
| [`todo_api_openapi.yaml`](todo_api_openapi.yaml) | Current OpenAPI 3.0.3 contract |
| [`schema.dbml`](schema.dbml) | MySQL tables, enums, indexes, and behavior notes |
| [`schema.dbdiagram`](schema.dbdiagram) | dbdiagram.io layout metadata (not the schema itself) |
| [`swagger.yaml`](swagger.yaml) | Legacy simple API (`pending` / `in_progress` / `completed`) |
| [`todo_test_stubs.go`](todo_test_stubs.go) | Legacy Go stubs for that simple API; not wired to this contract |

JSON uses camelCase (`createdAt`, `archivedAt`). The database uses snake_case (`created_at`, `archived_at`).

## User stories

| Epic | Stories |
| --- | --- |
| 1. Dashboard | US-01 View active todos in To Do / In Progress / Completed columns |
| 2. Todo management | US-02 Create · US-03 View · US-04 Update name and/or status · US-05 Soft-delete to recycle bin |
| 3. Search, filter, sort | US-06 Search by name · US-07 Filter by status and created date · US-08 Sort |
| 4. Checklist | US-09 Add/edit/remove items · US-10 Complete/reopen · US-11 Derived progress |
| 5. Archive | US-12 Archive · US-13 Unarchive |
| 6. Recycle bin | US-14 View · US-15 Restore · US-16 Permanently delete · US-17 Empty bin |
| 7. Bulk actions | US-18 Select many · US-19 Bulk move to recycle bin · US-20 Bulk permanent delete |
| 8. Errors | US-21 Clear validation, not-found, and consistent error responses |

## API

Success bodies wrap the payload in `{ "data": ... }`. Errors use:

```json
{
  "error": {
    "code": "validation_error",
    "message": "Request validation failed.",
    "details": [{ "field": "name", "message": "Must contain at least one non-whitespace character." }]
  }
}
```

Error codes: `validation_error`, `not_found`, `invalid_state`, `conflict`, `confirmation_required`, `invalid_cursor`, `internal_error`.

Typical HTTP status: `400` validation/cursor/missing confirm, `404` missing resource, `409` wrong lifecycle state, `500` unexpected error (no internal details).

Names must be 1–255 characters and not whitespace-only. Status values: `To Do`, `In Progress`, `Completed`. Create defaults status to `To Do`.

List and dashboard support `search`, `createdFrom`, `createdTo`, `sortBy` (`name`, `status`, `createdAt`, `updatedAt`; default `createdAt`), `sortOrder` (`asc`, `desc`; default `desc`). Cursors are opaque; reuse the same filters and sort when paging.

### Dashboard

`GET /dashboard` — active todos only, three independently paged columns (`limitPerColumn` default 20, max 100; `cursorToDo`, `cursorInProgress`, `cursorCompleted`). Archived and recycled todos are excluded.

### Todos

| Method | Path | Behavior |
| --- | --- | --- |
| `GET` | `/todos` | List. `view`: `active` (default), `archived`, `recycle-bin`. Repeatable `status`, plus search/filter/sort/`limit`/`cursor`. |
| `POST` | `/todos` | Create. Body: `{ "name": "...", "status": "To Do" }`. `201`. Starts active with an empty checklist. |
| `GET` | `/todos/{todoId}` | Details, checklist in display order, derived progress. Recycle-bin todos remain readable until permanent delete. |
| `PATCH` | `/todos/{todoId}` | Update name, status, or both. Archived/recycled todos cannot be edited here (`409`). |
| `DELETE` | `/todos/{todoId}` | Soft-delete an **active** todo to the recycle bin (`200` with the todo). Unarchive first if it is archived. |
| `POST` | `/todos/{todoId}/archive` | Archive an active todo. |
| `POST` | `/todos/{todoId}/unarchive` | Return an archived todo to active. Restore first if it is in the recycle bin. |
| `POST` | `/todos/{todoId}/restore` | Restore from recycle bin to active. |
| `DELETE` | `/todos/{todoId}/permanent?confirm=true` | Permanently delete from recycle bin (`204`). Client must confirm first. |

### Checklist items

Checklist completion does **not** change the parent todo’s status. Completing an item sets `completedAt`; reopening clears it.

| Method | Path | Behavior |
| --- | --- | --- |
| `GET` | `/todos/{todoId}/checklist-items` | Items ordered by `position`, then id. |
| `POST` | `/todos/{todoId}/checklist-items` | Add one or more items. `201`. Omit `position` to append. |
| `GET` | `/todos/{todoId}/checklist-items/{itemId}` | One item. |
| `PATCH` | `/todos/{todoId}/checklist-items/{itemId}` | Name, `isCompleted`, and/or `position`. At least one field. |
| `DELETE` | `/todos/{todoId}/checklist-items/{itemId}` | Remove one item (`204`). Progress is recalculated. |

### Recycle bin and bulk

Destructive bulk/empty calls require `"confirm": true` after a client-side confirmation dialog. Do not call them if the user cancels.

| Method | Path | Behavior |
| --- | --- | --- |
| `POST` | `/recycle-bin/empty` | Permanently delete every recycle-bin todo. Empty bin returns `deletedCount: 0`. |
| `POST` | `/todos/bulk/move-to-recycle-bin` | Soft-delete selected **active** todos. Per-id success/failure summary (`200`). |
| `POST` | `/todos/bulk/permanently-delete` | Permanently delete selected recycle-bin todos. Same summary shape. |

Bulk response includes `requestedCount`, `succeededCount`, `failedCount`, `succeededIds`, and `failures` (`not_found` or `invalid_state` per id).

## Database

MySQL. Diagram source: [`schema.dbml`](schema.dbml).

**`todos`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigint` | Primary key, auto-increment |
| `name` | `varchar(255)` | Required |
| `status` | enum `To Do`, `In Progress`, `Completed` | Default `To Do` |
| `created_at` | `datetime` | Default `CURRENT_TIMESTAMP` |
| `updated_at` | `datetime` | Default `CURRENT_TIMESTAMP`; set on every change (app or `ON UPDATE`) |
| `archived_at` | `datetime` | Null unless archived |
| `deleted_at` | `datetime` | Null unless in the recycle bin |

Indexes: name, status, created_at, updated_at, archived_at, deleted_at, and `(status, archived_at, deleted_at)`.

**`checklist_items`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigint` | Primary key, auto-increment |
| `todo_id` | `bigint` | FK to `todos.id`, cascade delete/update |
| `name` | `varchar(255)` | Required |
| `is_completed` | `boolean` | Default `false` |
| `position` | `int` | Default `0`; display order |
| `created_at` / `updated_at` | `datetime` | Same timestamp rules as todos |
| `completed_at` | `datetime` | Set when completed; cleared when reopened |

Indexes: `(todo_id, position)`, `(todo_id, is_completed)`.

A todo with no checklist items should show an empty state, not a fake progress percentage.

