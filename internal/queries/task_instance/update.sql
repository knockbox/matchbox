UPDATE
    ecs_task_instances
SET
    pull_start = ?,
    pull_stop = ?,
    started_at = ?,
    stopped_at = ?,
    stopped_reason = ?,
    status = ?
WHERE
    instance_owner_id = ?
AND
    ecs_task_definition_id = ?