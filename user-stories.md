# Todo Application --- Epics and User Stories

A **Todo** is a main task, such as “Update package.” A todo can have multiple **checklist items**, such as “Check updated package info,” “Verify previous package,” and “Update and test the new package.” Each checklist item belongs to one todo and can be marked complete independently. Checklist items are not separate todos on the main dashboard. The todo's own status (`To Do`, `In Progress`, or `Completed`) is managed separately from its checklist progress.

## Epic 1: Todo Dashboard

### US-01: View the Todo Dashboard

**User story**\
As a user, I want to see my active todos on a dashboard so that I can
quickly understand and manage my tasks.

**Acceptance criteria** 1. The dashboard displays active todos in three columns: To Do, In Progress, and Completed. 2. Each todo shows its name and status. 3. The user can open a todo's details from the dashboard. 4. The user can move a todo directly from any status column to any other status column. 5. Moving a todo updates its status and the dashboard immediately reflects the new column. 6. Archived and deleted todos do not appear on the active dashboard. 7. The dashboard provides access to creating, searching, filtering, and sorting todos. 8. If a column has no todos, a helpful empty state is displayed. 9. If the list is large, the user can navigate through the results without missing or repeating items.

## Epic 2: Todo Management

### US-02: Create a Todo

**User story**\
As a user, I want to create a todo with a name so that I can keep track
of a task I need to complete.

**Acceptance criteria** 1. The user can create a todo with a non-empty
name. 2. Each todo receives a unique identifier. 3. If no status is
provided, the status defaults to `To Do`. 4. The supported statuses
are `To Do`, `In Progress`, and `Completed`. 5. A created todo appears
in the active todo list. 6. The todo records its creation and
last-updated dates. 7. Invalid input produces a clear validation message
and does not create a todo.

### US-03: View a Todo

**User story**\
As a user, I want to view a todo's details so that I can check its information, status, and checklist steps.

**Acceptance criteria**

1. The user can open an existing todo.
2. The details show the name, status, creation date, and last-updated date.
3. If the todo has checklist items, the details show their names and completion states.
4. If the todo has checklist items, the details show completed and total item counts.
5. If the todo does not exist or is unavailable to the user, a clear not-found message is shown.

### US-04: Update a Todo

**User story**\
As a user, I want to update a todo so that I can keep its details and
status up to date.

**Acceptance criteria** 1. The user can update a todo's name, status, or
both. 2. The name cannot be empty. 3. The status must be `To Do`,
`In Progress`, or `Completed`. 4. The last-updated date changes after a
successful update. 5. Updated details are reflected in the dashboard and
relevant lists. 6. Invalid input produces a clear validation message. 7.
If the todo does not exist or is unavailable to the user, a clear
not-found message is shown.

### US-05: Delete a Todo to the Recycle Bin

**User story**\
As a user, I want to move a todo to the recycle bin so that it is removed from my active tasks without being permanently lost immediately.

**Acceptance criteria**

1. The user can move an active todo to the recycle bin.
2. The todo no longer appears in the active dashboard or list.
3. The todo appears in the recycle bin.
4. Its details and checklist items remain available while it is in the recycle bin.
5. Moving a todo to the recycle bin is different from permanently deleting it.
6. If the todo cannot be found, a clear message is shown.

## Epic 3: Search, Filtering, and Sorting

### US-06: Search Todos by Name

**User story**\
As a user, I want to search todos by name so that I can quickly find a
specific task.

**Acceptance criteria** 1. The user can enter a search term. 2. The
results show todos whose names match the term. 3. Search works together
with the selected filters. 4. The user can clear the search. 5. If there
are no matches, a helpful message is shown.

### US-07: Filter Todos

**User story**\
As a user, I want to filter todos so that I can focus on tasks matching
my needs.

**Acceptance criteria** 1. The user can filter todos by status. 2. The
user can filter todos by creation-date range. 3. More than one filter
can be applied at the same time. 4. Results match the selected filters
and search term. 5. The user can clear individual filters or reset all
filters. 6. If no todos match, a helpful message is shown. 7. The active
list excludes archived and recycled todos unless the user is viewing
those sections.

### US-08: Sort Todos

**User story**\
As a user, I want to sort todos so that I can arrange tasks in a useful
order.

**Acceptance criteria** 1. The user can sort todos by name, creation
date, last-updated date, or status. 2. The user can choose ascending or
descending order where applicable. 3. The selected order is reflected in
the displayed results. 4. Sorting does not change a todo's stored
information. 5. Sorting works together with search and filters.

## Epic 4: Todo Checklist Management

### US-09: Add and Manage Checklist Items

**User story**\
As a user, I want to add and manage checklist items within a todo so that I can break a larger task into smaller steps.

**Acceptance criteria**

1. The user can add one or more checklist items to an existing todo.
2. Each checklist item must have a non-empty name and a unique identifier.
3. Each checklist item belongs to one parent todo and appears in that todo's details.
4. The user can edit or remove a checklist item.
5. Checklist items are steps within a todo, not separate todos in the main dashboard.
6. A todo can exist without any checklist items.
7. Invalid checklist item input produces a clear validation message.
8. Checklist items remain attached when the todo is archived or moved to the recycle bin.

### US-10: Complete and Reopen Checklist Items

**User story**\
As a user, I want to mark individual checklist items as complete or incomplete so that I can track which steps of a todo are finished.

**Acceptance criteria**

1. The user can mark a checklist item as complete.
2. The user can mark a completed checklist item as incomplete again.
3. Each checklist item shows whether it is complete or incomplete.
4. Updating one checklist item does not change the completion state of other items.
5. Checklist item completion does not automatically change the parent todo's status; todo status is managed separately.
6. Changes to checklist items are reflected in the todo details and checklist progress.

### US-11: View Checklist Progress

**User story**\
As a user, I want to see checklist progress for a todo so that I can understand how many steps are complete and how many remain.

**Acceptance criteria**

1. The todo details show the checklist items and each item's completion state.
2. The application shows the number of completed items and total checklist items; remaining items can be understood from these counts.
3. Checklist progress updates when an item is completed or reopened, or when an item is added or removed.
4. A todo with no checklist items displays a helpful empty state instead of misleading progress.
5. Checklist progress is calculated from the checklist items; it is not stored as a separate manually edited value.

## Epic 5: Archive Management

### US-12: Archive a Todo

**User story**\
As a user, I want to archive a todo that is not relevant right now so that I can keep my active task list focused without deleting the task.

**Acceptance criteria**

1. The user can archive an active todo.
2. The archived todo no longer appears in the active dashboard or list.
3. The todo appears in the archived-todos view.
4. Archiving preserves the todo's details, status, and checklist items.
5. Archiving does not move the todo to the recycle bin.

### US-13: Unarchive a Todo

**User story**\
As a user, I want to unarchive a todo so that I can bring it back to my active tasks when it becomes relevant again.

**Acceptance criteria**

1. The user can unarchive a todo from the archived-todos view.
2. The unarchived todo appears in the active todo list.
3. Its details, status, and checklist items are preserved.
4. A todo in the recycle bin must be restored before it can be unarchived.

## Epic 6: Recycle Bin Management

### US-14: View Deleted Todos

**User story**\
As a user, I want to view deleted todos in a recycle bin so that I can
review and recover tasks deleted by mistake.

**Acceptance criteria** 1. The user can open the recycle bin. 2. The
recycle bin displays todos moved there. 3. Each item shows enough
information to identify it. 4. Recycled todos do not appear in active or
archived lists. 5. If the recycle bin is empty, a helpful empty state is
shown.

### US-15: Restore a Todo

**User story**\
As a user, I want to restore a todo from the recycle bin so that I can recover a task I still need.

**Acceptance criteria**

1. The user can restore a todo from the recycle bin.
2. The restored todo returns to the active todo list.
3. The todo's retained details and checklist items are preserved.
4. The restored todo no longer appears in the recycle bin.
5. A permanently deleted todo and its checklist items cannot be restored.

### US-16: Permanently Delete a Todo

**User story**\
As a user, I want to permanently delete a todo from the recycle bin so that I can remove tasks I no longer need.

**Acceptance criteria**

1. The user can permanently delete a todo from the recycle bin.
2. The application asks the user to confirm permanent deletion.
3. After confirmation, the todo and its associated checklist items can no longer be viewed or restored.
4. If the todo is unavailable, a clear message is shown.

### US-17: Empty the Recycle Bin

**User story**\
As a user, I want to empty the recycle bin so that I can permanently remove all deleted todos when I no longer need them.

**Acceptance criteria**

1. The user can choose to empty the recycle bin.
2. The application warns that this permanently deletes the items in the recycle bin.
3. The user must confirm before the action is completed.
4. After confirmation, all todos in the recycle bin and their associated checklist items are permanently deleted.
5. If the user cancels, the items remain in the recycle bin.

## Epic 7: Bulk Actions

### US-18: Select Multiple Todos

**User story**\
As a user, I want to select multiple todos so that I can manage several
tasks in one action.

**Acceptance criteria** 1. The user can select multiple todos from a
list. 2. Selected todos are clearly indicated. 3. The user can clear the
selection. 4. Available bulk actions match the current list. 5. The
application indicates when no todos are selected.

### US-19: Move Multiple Todos to the Recycle Bin

**User story**\
As a user, I want to move multiple selected todos to the recycle bin at
once so that I can remove several tasks efficiently.

**Acceptance criteria** 1. The user can select multiple active todos and
move them to the recycle bin in one action. 2. The application asks for
confirmation. 3. After confirmation, selected todos are removed from the
active list and appear in the recycle bin. 4. The user receives a
summary of successful and unsuccessful items. 5. If the user cancels,
none of the selected todos are moved.

### US-20: Permanently Delete Multiple Todos

**User story**\
As a user, I want to permanently delete multiple selected todos from the
recycle bin so that I can clean up deleted tasks efficiently.

**Acceptance criteria** 1. The user can select multiple todos in the
recycle bin. 2. The application asks for confirmation before permanent
deletion. 3. After confirmation, the selected todos are permanently
deleted. 4. The user receives a summary of successful and unsuccessful
items. 5. If the user cancels, the selected todos remain in the recycle
bin.

## Epic 8: Error Handling

### US-21: Receive Clear Error Messages

**User story**\
As a user, I want clear and consistent error messages so that I understand what went wrong and what I can do next.

**Acceptance criteria**

1. Missing or invalid required input produces a clear validation message.
2. A todo or checklist item that does not exist produces a clear not-found message.
3. Error messages do not expose sensitive internal details.
4. Error responses behave consistently across the application.
5. A failed action is not presented as successful.
