# 🧠 Core Logic and System Behavior

This document explains the core logic behind **Daylick**: how templates, types, data, and autosaving work together.

---

## 🔁 Templates vs Types

### Templates
- Are predefined block structures (e.g., Sleep, Study, Protein).
- Can be for: `plan`, `daily`, or `stats`.
- Stored in the `templates` and `template_fields` tables.
- Shared between users or used as defaults.

### Types
- Are workspace-local "clones" of templates.
- Stored in the `types` and `type_fields` tables.
- A type can **evolve independently** from its template.
- A type has an `original_template_id` to trace its origin.
- If `type_fields` ≠ `template_fields`, it's marked as **dirty** (user changed the type).

---

## 📋 Plans

- Plans use a `type` (of kind = `plan`) to define their structure.
- Each plan has a `start_date` and `end_date` (supports flexible timelines).
- Each block in the plan (e.g., Protein: 60g) is stored in `plan_entries`.

---

## 📅 Daily Reports

- Daily reports are generated using a `type` (of kind = `daily`).
- Each report has input fields matching the type blocks.
- Users can manually update a daily report, which triggers autosave logic.

---

## 📈 Statistics

- Statistics are **computed dynamically** (not stored).
- Based on aggregation fields (`sum`, `count`, `average`) defined in type blocks.
- Aggregation is calculated from either `plan_entries` or `daily_report_entries`.

---

## 💾 Autosave Behavior

- All inputs (plans and reports) are saved via **debounced autosave**:
  - If no changes for 5 seconds → save.
  - Only changed fields are sent in a PATCH request.
- On browser close, a final save attempt is made via `navigator.sendBeacon()` (optional).
- Manual "Save" button is avoided for smoother UX.

---

## 🔧 Updating a Type

- When user changes a type (e.g., renames "Protein" to "Protein Intake"):
  - The `type_fields` are updated.
  - The type becomes "dirty" if it no longer matches its original template.
- User can choose:
  - “Save as new template”
  - “Overwrite original template”
  - Or do nothing (leave the type as-is)

---

## 🔄 Type Sync Prompt (Daily/Plans)

- If today’s daily report does not match the current `daily` type:
  - Show a prompt: “Update today’s report to match the latest structure?”
  - If accepted:
    - New fields are added
    - Old fields are kept (if still compatible)
    - Data is preserved where possible
