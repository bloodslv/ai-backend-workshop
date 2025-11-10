## Clean Architecture Structure

```
┌─────────────────────────────────────────────────────────┐
│                    HTTP Layer                            │
│  ┌─────────────────────────────────────────────────────┐ │
│  │             Fiber Routes & Middleware              │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────┬───────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────┐
│                 Handler Layer                           │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  UserHandler (HTTP Controllers)                    │ │
│  │  - GetUsers()                                       │ │
│  │  - GetUser()                                        │ │
│  │  - CreateUser()                                     │ │
│  │  - UpdateUser()                                     │ │
│  │  - DeleteUser()                                     │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────┬───────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────┐
│                Use Case Layer                           │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  UserUseCase (Business Logic)                      │ │
│  │  - GetAllUsers()                                    │ │
│  │  - GetUserByID()                                    │ │
│  │  - CreateUser()                                     │ │
│  │  - UpdateUser()                                     │ │
│  │  - DeleteUser()                                     │ │
│  │  - Business rules & validation                     │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────┬───────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────┐
│                 Domain Layer                            │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  User Entity & Interfaces                          │ │
│  │  - User struct                                      │ │
│  │  - UserRepository interface                        │ │
│  │  │  - UserUseCase interface                        │ │
│  │  - Request/Response DTOs                           │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────┬───────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────┐
│              Repository Layer                           │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  UserRepository (Data Access)                      │ │
│  │  - GetAll()                                         │ │
│  │  - GetByID()                                        │ │
│  │  - GetByEmail()                                     │ │
│  │  - Create()                                         │ │
│  │  - Update()                                         │ │
│  │  - Delete()                                         │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────┬───────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────┐
│              Infrastructure                             │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  Database (SQLite + GORM)                          │ │
│  │  Configuration                                      │ │
│  │  External Services                                  │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────┘
```

## Dependency Flow

```
Handler → UseCase → Repository → Database
   ↑         ↑          ↑
   │         │          │
Depends   Depends   Implements
  on        on       Interface
   │         │          │
   ▼         ▼          ▼
UseCase   Repository  Domain
Interface  Interface   Entity
```

## Benefits

1. **Testability** - Each layer can be tested independently
2. **Maintainability** - Changes in one layer don't affect others
3. **Flexibility** - Easy to swap implementations
4. **Business Logic Protection** - Core logic is isolated from external concerns
5. **Clear Responsibilities** - Each layer has a single responsibility
