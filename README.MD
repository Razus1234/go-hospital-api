# Hospital API

## 🚀 Tech Stack
- **Language:** Go (Golang)
- **Web Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL
- **API Docs:** Swagger (swaggo)
- **Containerization:** Docker, Docker Compose
- **Reverse Proxy:** Nginx
- **Testing:** Testify
- **JWT Auth:** golang-jwt/jwt

---

## 📁 Project Structure
```
go-hospital-api/
├── Dockerfile
├── Makefile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── nginx/
│   └── default.conf
├── internal/
│   ├── db/
│   │   └── db.go
│   ├── dto/
│   │   └── dto.go
│   ├── entities/
│   │   ├── hospital.go
│   │   ├── patient.go
│   │   └── staff.go
│   ├── handlers/
│   │   ├── hospital_handler.go
│   │   ├── patient_handler.go
│   │   └── staff_handler.go
│   ├── middleware/
│   │   └── auth.go
│   ├── repository/
│   │   ├── hospital_repo.go
│   │   ├── patient_repo.go
│   │   └── staff_repo.go
│   ├── services/
│   │   ├── hospital_service.go
│   │   ├── patient_service.go
│   │   └── staff_service.go
│   ├── tests/
│   │   ├── patient_handler_test.go
│   │   └── staff_handler_test.go
│   └── utils/
│       └── jwt.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── .env
└── .vscode/
  └── launch.json
```
---

## 1.1. Go API Layering

The project follows a clean, layered architecture:

```
Client <--> Handler <--> Service <--> Repository <--> DB (Entity Model)
     ↑            ↑          ↑            ↑
   RequestDTO   Convert     Business     DB Model
   ResponseDTO  to Entity    Logic        (GORM Model)
```

- **Client**: Sends HTTP requests (e.g., Postman, frontend, curl)
- **Handler**: Receives requests, validates input, returns responses (Gin HTTP handlers)
- **Service**: Contains business logic, orchestrates data flow
- **Repository**: Handles DB access (CRUD), maps to GORM models
- **DB (Entity Model)**: GORM models representing database tables

DTOs (Data Transfer Objects) are used for request/response payloads and are converted to/from entity models as needed.



---

## 2. API Specification

### Authentication

- JWT Bearer token in `Authorization` header.
- Example: `Authorization: Bearer {token}`

### Endpoints

#### Staff

- **POST /api/staff/create**
  - Create a new staff member.
  - Request body: `StaffCreateRequest`
  - Response: `StaffResponse` or `ErrorResponse`

- **POST /api/staff/login**
  - Login as staff, returns JWT.
  - Request body: `StaffLoginRequest`
  - Response: `StaffLoginResponse` or `ErrorResponse`

#### Patients

- **POST /api/patients/search** (Requires Auth)
  - Search for patients by criteria, restricted to staff's hospital.
  - Request body: `PatientSearchCriteria`
  - Response: `[]PatientResponse` or `ErrorResponse`
---

## 3. ER-Diagram

[View ER Diagram on dbdiagram.io](https://dbdiagram.io/d/6895f0d4dd90d17865181592)




## 4. Mock Data
```sql
INSERT INTO public.hospitals (id, name, created_at, updated_at) VALUES
  ('11111111-1111-1111-1111-111111111111', 'Hospital A', NOW(), NOW()),
  ('22222222-2222-2222-2222-222222222222', 'Hospital B', NOW(), NOW());

-- Insert Staff
INSERT INTO public.staffs (id, username, password_hash, hospital_id, created_at, updated_at) VALUES
  ('aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1', 'admin1', '$2a$10$.X1SXobGCJdlX.7.QjYGkOGhVZotMiZYgbJ2BSCOyb7YE99pT5xTS', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
  ('aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2', 'admin2', '$2a$10$.X1SXobGCJdlX.7.QjYGkOGhVZotMiZYgbJ2BSCOyb7YE99pT5xTS', '22222222-2222-2222-2222-222222222222', NOW(), NOW());

-- Insert Patients for Hospital A (5 patients)
INSERT INTO public.patients (id, first_name_th, last_name_th, first_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender, hospital_id, created_at, updated_at) VALUES
  (gen_random_uuid(), 'สมชาย', 'ใจดี', 'Somchai', 'Jaidee', '1980-05-15', 'HN001', '1234567890123', NULL, '0811111111', 'somchai@example.com', 'M', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
  (gen_random_uuid(), 'สมหญิง', 'ใจงาม', 'Somying', 'Jaingam', '1985-09-20', 'HN002', '2234567890123', NULL, '0822222222', 'somying@example.com', 'F', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
  (gen_random_uuid(), 'วีระ', 'กล้าหาญ', 'Weera', 'Klaharn', '1990-01-05', 'HN003', '3234567890123', NULL, '0833333333', 'weera@example.com', 'M', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
  (gen_random_uuid(), 'กมล', 'สุขใจ', 'Kamon', 'Sookjai', '1992-11-10', 'HN004', '4234567890123', NULL, '0844444444', 'kamon@example.com', 'F', '11111111-1111-1111-1111-111111111111', NOW(), NOW()),
  (gen_random_uuid(), 'พีระ', 'พงษ์ไทย', 'Peera', 'Pongthai', '1995-08-25', 'HN005', '5234567890123', NULL, '0855555555', 'peera@example.com', 'M', '11111111-1111-1111-1111-111111111111', NOW(), NOW());

-- Insert Patients for Hospital B (7 patients)
INSERT INTO public.patients (id, first_name_th, last_name_th, first_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender, hospital_id, created_at, updated_at) VALUES
  (gen_random_uuid(), 'สมปอง', 'ใจบุญ', 'Sompong', 'Jaiboon', '1983-02-14', 'HN101', '6234567890123', NULL, '0866666666', 'sompong@example.com', 'M', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
  (gen_random_uuid(), 'วิภา', 'ใจเย็น', 'Wipha', 'Jaiyen', '1987-07-19', 'HN102', '7234567890123', NULL, '0877777777', 'wipha@example.com', 'F', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
  (gen_random_uuid(), 'สมร', 'ทองดี', 'Samorn', 'Thongdee', '1991-03-11', 'HN103', '8234567890123', NULL, '0888888888', 'samorn@example.com', 'F', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
  (gen_random_uuid(), 'บุญช่วย', 'ใจซื่อ', 'Boonchuay', 'Jaisue', '1989-06-23', 'HN104', '9234567890123', NULL, '0899999999', 'boonchuay@example.com', 'M', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
  (gen_random_uuid(), 'พรทิพย์', 'บุญมา', 'Pornthip', 'Boonma', '1994-04-17', 'HN105', '1334567890123', NULL, '0810000000', 'pornthip@example.com', 'F', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
  (gen_random_uuid(), 'วสันต์', 'จิตดี', 'Wasant', 'Jitdee', '1988-10-09', 'HN106', '2334567890123', NULL, '0820000000', 'wasant@example.com', 'M', '22222222-2222-2222-2222-222222222222', NOW(), NOW()),
  (gen_random_uuid(), 'กาญจนา', 'ทองแท้', 'Kanchana', 'Thongthae', '1996-12-01', 'HN107', '3334567890123', NULL, '0830000000', 'kanchana@example.com', 'F', '22222222-2222-2222-2222-222222222222', NOW(), NOW());
```

## ▶️ How to Run
### Docker Compose (manual)

- `docker-compose up --build -d` — Start all services in the background
- `docker-compose up -d db` — Start only the database
- `docker-compose logs -f` — View logs

### API Documentation

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Environment

- Copy `.env` and adjust as needed for DB credentials, JWT secret, etc.

### Run tests

- `go test ./internal/tests/...`
