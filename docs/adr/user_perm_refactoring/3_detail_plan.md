# Детальный план реализации комбинированного решения

## Общая архитектура решения

Реализация комбинированного подхода с использованием кэширования и асинхронного обновления через Kafka:

```
┌─────────────┐    ┌──────────────┐    ┌─────────────┐
│   Client    │    │   Service    │    │   Kafka     │
│             │    │              │    │             │
│  Read Perm  │───▶│  Cache       │───▶│  Messages   │
│  (Fast)     │    │  (Redis)     │    │             │
└─────────────┘    │              │    └─────────────┘
                   │  Update      │
                   │  (Async)     │
                   │  ┌─────────┐ │
                   │  │  DB     │ │
                   │  └─────────┘ │
                   └──────────────┘
```

## Шаги реализации

### 1. Подготовка инфраструктуры

#### 1.1. Добавление Redis в систему
- Добавить Redis в docker-compose.yml
- Настроить подключение к Redis в конфигурации
- Реализовать интерфейс кэширования

#### 1.2. Подготовка Kafka
- Добавить Kafka в docker-compose.yml (если еще не добавлен)
- Настроить темы для сообщений о правах пользователей
- Реализовать producer для отправки сообщений

### 2. Реализация кэширования

#### 2.1. Создание интерфейса кэширования
```go
// cache/interface.go
type PermissionCache interface {
    GetUserPermissions(ctx context.Context, userID int) ([]*models.Permission, error)
    SetUserPermissions(ctx context.Context, userID int, permissions []*models.Permission) error
    InvalidateUserPermissions(ctx context.Context, userID int) error
    GetOrganizationPermissions(ctx context.Context, organizationID int) ([]*models.Permission, error)
    SetOrganizationPermissions(ctx context.Context, organizationID int, permissions []*models.Permission) error
    InvalidateOrganizationPermissions(ctx context.Context, organizationID int) error
}
```

#### 2.2. Реализация Redis кэша
```go
// cache/redis_cache.go
type redisCache struct {
    client *redis.Client
    ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) PermissionCache {
    return &redisCache{client: client, ttl: ttl}
}

func (c *redisCache) GetUserPermissions(ctx context.Context, userID int) ([]*models.Permission, error) {
    // Реализация получения из кэша
}

func (c *redisCache) SetUserPermissions(ctx context.Context, userID int, permissions []*models.Permission) error {
    // Реализация установки в кэш
}

func (c *redisCache) InvalidateUserPermissions(ctx context.Context, userID int) error {
    // Реализация очистки кэша
}
```

### 3. Интеграция с существующими репозиториями

#### 3.1. Модификация UserRepository
```go
// internal/repository/user_repository.go
type userRepository struct {
    db     *DB
    cache  PermissionCache // Добавляем зависимость от кэша
}

// GetUserPermissions с кэшированием
func (r *userRepository) GetUserPermissions(ctx context.Context, userID int) ([]*models.Permission, error) {
    // Сначала пытаемся получить из кэша
    cached, err := r.cache.GetUserPermissions(ctx, userID)
    if err == nil && cached != nil {
        return cached, nil
    }
    
    // Если кэш пуст или произошла ошибка, делаем запрос к БД
    permissions, err := r.getUserPermissionsFromDB(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // Сохраняем в кэш для последующего использования
    r.cache.SetUserPermissions(ctx, userID, permissions)
    
    return permissions, nil
}

// Метод для асинхронного обновления кэша
func (r *userRepository) updateUserPermissionsInCache(ctx context.Context, userID int) error {
    permissions, err := r.getUserPermissionsFromDB(ctx, userID)
    if err != nil {
        return err
    }
    
    return r.cache.SetUserPermissions(ctx, userID, permissions)
}
```

#### 3.2. Модификация OrganizationRepository
```go
// internal/repository/organization_repository.go
type organizationRepository struct {
    db     *DB
    cache  PermissionCache // Добавляем зависимость от кэша
}

// GetOrganizationPermissions с кэшированием
func (r *organizationRepository) GetOrganizationPermissions(ctx context.Context, organizationID int) ([]*models.Permission, error) {
    // Сначала пытаемся получить из кэша
    cached, err := r.cache.GetOrganizationPermissions(ctx, organizationID)
    if err == nil && cached != nil {
        return cached, nil
    }
    
    // Если кэш пуст или произошла ошибка, делаем запрос к БД
    permissions, err := r.getOrganizationPermissionsFromDB(ctx, organizationID)
    if err != nil {
        return nil, err
    }
    
    // Сохраняем в кэш для последующего использования
    r.cache.SetOrganizationPermissions(ctx, organizationID, permissions)
    
    return permissions, nil
}
```

### 4. Реализация асинхронного обновления через Kafka

#### 4.1. Создание сообщений Kafka
```protobuf
// api/grpc/api.proto
message UserPermissionEvent {
    int32 user_id = 1;
    repeated int32 added_permission_ids = 2;
    repeated int32 removed_permission_ids = 3;
    string event_type = 4; // "ADD", "REMOVE", "UPDATE"
    int64 timestamp = 5;
}

message OrganizationPermissionEvent {
    int32 organization_id = 1;
    repeated int32 added_permission_ids = 2;
    repeated int32 removed_permission_ids = 3;
    string event_type = 4; // "ADD", "REMOVE", "UPDATE"
    int64 timestamp = 5;
}
```

#### 4.2. Producer для отправки сообщений
```go
// kafka/producer.go
type KafkaProducer struct {
    producer sarama.SyncProducer
}

func (kp *KafkaProducer) PublishUserPermissionEvent(ctx context.Context, event *UserPermissionEvent) error {
    // Конвертация в JSON и отправка в Kafka
    return nil
}

func (kp *KafkaProducer) PublishOrganizationPermissionEvent(ctx context.Context, event *OrganizationPermissionEvent) error {
    // Конвертация в JSON и отправка в Kafka
    return nil
}
```

#### 4.3. Consumer для обновления кэша
```go
// kafka/consumer.go
type KafkaConsumer struct {
    consumer sarama.Consumer
    cache    PermissionCache
}

func (kc *KafkaConsumer) HandleUserPermissionEvent(ctx context.Context, event *UserPermissionEvent) error {
    // Обновление кэша пользователя
    return kc.cache.InvalidateUserPermissions(ctx, event.UserID)
}

func (kc *KafkaConsumer) HandleOrganizationPermissionEvent(ctx context.Context, event *OrganizationPermissionEvent) error {
    // Обновление кэша организации
    return kc.cache.InvalidateOrganizationPermissions(ctx, event.OrganizationID)
}
```

### 5. Обновление обработчиков

#### 5.1. Модификация обработчиков для отправки сообщений
```go
// internal/handlers/user_add_permissions.go
func (h *BaseHandler) AddUserPermissions(ctx context.Context, in *api_pb.UserAddPermissionsRequest) (*api_pb.Empty, error) {
    // ... существующая логика ...
    
    // Отправка события в Kafka
    event := &UserPermissionEvent{
        UserID:                 in.UserId,
        AddedPermissionIDs:     in.PermissionIds,
        EventType:              "ADD",
        Timestamp:              time.Now().Unix(),
    }
    h.kafkaProducer.PublishUserPermissionEvent(ctx, event)
    
    return &api_pb.Empty{}, nil
}
```

### 6. Добавление фоновых задач для синхронизации

#### 6.1. Периодическая синхронизация кэша
```go
// cache/sync.go
type CacheSync struct {
    cache    PermissionCache
    db       *DB
    interval time.Duration
}

func (cs *CacheSync) StartSyncLoop(ctx context.Context) {
    ticker := time.NewTicker(cs.interval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            cs.syncAllCaches(ctx)
        }
    }
}

func (cs *CacheSync) syncAllCaches(ctx context.Context) {
    // Синхронизация всех кэшей с БД
}
```

### 7. Обновление документации

#### 7.1. Добавление в ADR
Обновить документацию по архитектуре для отражения нового подхода к управлению правами.

### 8. Тестирование

#### 8.1. Unit тесты
- Тестирование кэширующих методов
- Тестирование обработчиков событий Kafka
- Тестирование синхронизации кэша

#### 8.2. Integration тесты
- Тестирование полного цикла: изменение прав → отправка события → обновление кэша → чтение из кэша

## План реализации по этапам

### Этап 1: Подготовка инфраструктуры (1-2 дня)
1. Добавление Redis в docker-compose
2. Настройка подключения к Redis
3. Создание интерфейса кэширования

### Этап 2: Реализация кэширования (2-3 дня)
1. Реализация Redis кэша
2. Модификация UserRepository с кэшированием
3. Модификация OrganizationRepository с кэшированием

### Этап 3: Асинхронное обновление (2-3 дня)
1. Создание Kafka сообщений
2. Реализация Producer
3. Реализация Consumer
4. Интеграция с обработчиками

### Этап 4: Фоновая синхронизация (1-2 дня)
1. Реализация фоновой синхронизации
2. Настройка периодических задач

### Этап 5: Тестирование и документация (1-2 дня)
1. Написание unit/integration тестов
2. Обновление документации
3. Проверка производительности

## Ожидаемые результаты

1. **Улучшение производительности**: Запросы к правам пользователей и организаций будут выполняться в 10-100 раз быстрее
2. **Снижение нагрузки на БД**: Чтение будет происходить из кэша, а не из сложных SQL-запросов
3. **Масштабируемость**: Система будет хорошо масштабироваться при увеличении нагрузки
4. **Надежность**: Резервные механизмы обеспечивают отказоустойчивость

Этот подход обеспечивает баланс между производительностью, простотой поддержки и масштабируемостью, позволяя эффективно решить проблему медленных запросов к правам пользователей и организаций.