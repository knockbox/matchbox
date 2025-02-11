UPDATE
    ecs_task_instances
SET
    aws_arn = ?,
    pull_start = ?,
    pull_stop = ?,
    started_at = ?,
    stopped_at = ?,
    stopped_reason = ?,
    status = ?
WHERE
    ecs_task_definition_id = ?
AND
    instance_owner_id = ?