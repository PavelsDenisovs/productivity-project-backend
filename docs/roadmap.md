# MVP Implementation Roadmap (6 Months)

---

### ✅ Step 1: User + Workspace system

**Database**
- `users`, `workspaces`, `workspace_members`

**Backend**
- Workspace creation (default per user)
- Middleware: load workspace via session

**Frontend**
- Workspace switcher (for >1 workspace later)
- Show workspace name in header

---

### ✅ Step 2: Templates + Types Base

**Database**
- `templates`, `template_fields`
- `types`, `type_fields`

**Backend**
- Clone template → type
- Detect type ≠ template → show "dirty" warning

**Frontend**
- View types
- View template-derived types
- (Later) Save new template from modified type

---

### ⏳ Step 3: Weekly Plans

**Database**
- `plans`, `plan_entries`

**Backend**
- `POST /plans`
- `PATCH /plans/:id/entry`

**Frontend**
- Plan editor for current period
- Inputs for each field per day
- Show goals, expected values

---

### ⏳ Step 4: Daily Reports

**Database**
- `daily_reports`, `daily_report_entries`

**Backend**
- Auto-generate report by type
- Patch specific entry fields

**Frontend**
- Daily note UI with autosave

---

### ⏳ Step 5: Statistics (MVP)

- No DB table yet (computed on read)
- Show charts of plan/report values
- Use `recharts` on frontend

---

### ⏳ Step 6: Type/Template Editing

- Edit types (field name, unit, aggregation)
- Save as new template or overwrite existing
- Track `type.original_template_id` to know if modified