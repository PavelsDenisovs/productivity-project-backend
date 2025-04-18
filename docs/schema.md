# 🗃️ Database Schema (MVP Scope)

---

## ✅ users

| Column        | Type      | Description                  |
|---------------|-----------|------------------------------|
| id            | UUID      | Primary key                  |
| email         | TEXT      | User email (unique)          |
| password_hash | TEXT      | Hashed password              |
| is_verified   | BOOLEAN   | Email verification flag      |
| created_at    | TIMESTAMPTZ | Account creation timestamp |
| updated_at    | TIMESTAMPTZ | Last update timestamp      |

---

## ✅ verification_codes

| Column     | Type        | Description                                |
|------------|-------------|--------------------------------------------|
| id         | UUID        | Primary key                                |
| user_id    | UUID        | FK → users.id (unique per user)            |
| code       | TEXT        | Verification code                          |
| created_at | TIMESTAMPTZ | When code was generated                    |
| expires_at | TIMESTAMPTZ | When code expires                          |
| used       | BOOLEAN     | If code was already used                   |

---

## ✅ workspaces

| Column         | Type        | Description                        |
|----------------|-------------|------------------------------------|
| id             | UUID        | Primary key                        |
| name           | TEXT        | Workspace name                     |
| owner_user_id  | UUID        | FK → users.id                      |
| type_id        | UUID        | FK → types.id (default plan type)  |
| created_at     | TIMESTAMPTZ | Creation timestamp                 |

---

## ✅ workspace_members

| Column       | Type    | Description                         |
|--------------|---------|-------------------------------------|
| user_id      | UUID    | FK → users.id                       |
| workspace_id | UUID    | FK → workspaces.id                  |
| role         | ENUM    | 'owner', 'editor', or 'viewer'      |

---

## ✅ templates

| Column            | Type    | Description                        |
|-------------------|---------|------------------------------------|
| id                | UUID    | Primary key                        |
| name              | TEXT    | Template name                      |
| kind              | ENUM    | 'plan', 'daily', or 'stats'        |
| author_user_id    | UUID    | FK → users.id                      |
| created_at        | TIMESTAMPTZ | When template was created     |

---

## ✅ template_fields

| Column        | Type    | Description                             |
|---------------|---------|-----------------------------------------|
| id            | UUID    | Primary key                             |
| template_id   | UUID    | FK → templates.id                       |
| name          | TEXT    | Block name (e.g. Sleep)                 |
| unit          | TEXT    | Measurement unit (e.g. hours, g)        |
| aggregation   | ENUM    | 'sum', 'count', 'average'               |
| order         | INT     | Display order                           |

---

## ✅ types

| Column                | Type    | Description                              |
|-----------------------|---------|------------------------------------------|
| id                    | UUID    | Primary key                              |
| kind                  | ENUM    | 'plan', 'daily', or 'stats'              |
| workspace_id          | UUID    | FK → workspaces.id                       |
| original_template_id  | UUID    | FK → templates.id (optional reference)   |
| name                  | TEXT    | Name of the type                         |
| created_at            | TIMESTAMPTZ | Timestamp                            |

---

## ✅ type_fields

| Column      | Type    | Description                          |
|-------------|---------|--------------------------------------|
| id          | UUID    | Primary key                          |
| type_id     | UUID    | FK → types.id                        |
| name        | TEXT    | Field name (e.g., "Protein")         |
| unit        | TEXT    | Measurement unit (g, hours, etc.)    |
| aggregation | ENUM    | 'sum', 'average', 'count'            |
| order       | INT     | Position in display                  |

---

## ✅ plans

| Column       | Type        | Description                          |
|--------------|-------------|--------------------------------------|
| id           | UUID        | Primary key                          |
| workspace_id | UUID        | FK → workspaces.id                   |
| type_id      | UUID        | FK → types.id                        |
| start_date   | DATE        | Start of plan period                 |
| end_date     | DATE        | End of plan period                   |
| created_at   | TIMESTAMPTZ | Timestamp                            |

---

## ✅ plan_entries

| Column        | Type    | Description                          |
|---------------|---------|--------------------------------------|
| id            | UUID    | Primary key                          |
| plan_id       | UUID    | FK → plans.id                        |
| type_field_id | UUID    | FK → type_fields.id                  |
| value         | FLOAT   | Target value for the day             |
| day           | DATE    | The specific date                    |

---

## ✅ daily_reports

| Column       | Type        | Description                     |
|--------------|-------------|---------------------------------|
| id           | UUID        | Primary key                     |
| workspace_id | UUID        | FK → workspaces.id              |
| type_id      | UUID        | FK → types.id                   |
| date         | DATE        | Date of the report              |
| notes        | TEXT        | Optional freeform notes         |
| created_at   | TIMESTAMPTZ | Timestamp                       |

---

## ✅ daily_report_entries

| Column        | Type    | Description                          |
|---------------|---------|--------------------------------------|
| id            | UUID    | Primary key                          |
| daily_report_id | UUID  | FK → daily_reports.id                |
| type_field_id | UUID    | FK → type_fields.id                  |
| value         | FLOAT   | User-entered value                   |
