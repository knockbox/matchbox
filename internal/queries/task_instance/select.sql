SELECT
    *
FROM
    ecs_task_instances
WHERE
    ecs_task_definition_id = ?
AND
    instance_owner_id = ?