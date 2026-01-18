## Цель
Подумать над упрощением запросов
- пользователи
```sql
SELECT DISTINCT p.id, p.name, p.code, p.description
FROM permissions p
         LEFT JOIN user_permissions up ON p.id = up.permission_id and up.user_id = $1
         LEFT JOIN user_roles ur ON ur.user_id = $1
         LEFT JOIN role_permissions rp ON rp.role_id = ur.role_id and p.id = rp.permission_id
WHERE ur.user_id = $1
   or up.user_id = $1
ORDER BY p.id
```
- организации
```sql
SELECT DISTINCT p.id, p.name, p.code, p.description
FROM permissions p
        LEFT JOIN organization_permissions op ON p.id = op.permission_id and op.organization_id = $1
        LEFT JOIN organization_roles ro ON ro.organization_id = $1
        LEFT JOIN role_permissions rp ON rp.role_id = ro.role_id and p.id = op.permission_id
WHERE ro.organization_id = $1
   or op.organization_id = $1
ORDER BY p.id
```

### Возможные риски
- Безиндексное чтение и/или джоины нескольких таблий работают медленно,
  при повышении читающей нагрузки будут просадки латенси

### Решение через кафка
В запросе оставить только таблицу user_permissions, а все добавления и удаления прав кидать в кафку
Читающий консюмер физически добавит или удалит право в таблице user_permissions
Это будет редким событием по сравнению с читающим трафиком, а латенси не просядет
Для сохранения метаинформации вести логи (источник возникновения/удаления права и дата когда события поступило)