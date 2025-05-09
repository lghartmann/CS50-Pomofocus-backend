-- name: Search :many
SELECT
    u.name as userName,
    u.email as userEmail,
    p.duration,
    p.pause_duration as pauseDuration,
    p.effort,
    p.distraction,
    p.productivity,
    p.created_at as createdAt
FROM
    users u
    INNER JOIN pomodoro p ON p.user_id = u.id
    AND p.deleted_at IS NULL
    WHERE u.id = $1
    LIMIT = $2
    START = $3
    ORDER BY p.id DESC;

-- name: Create :one
INSERT INTO pomodoro(
    user_id,
    duration,
    pause_duration,
    effort,
    distraction,
    productivity,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: Update :one
UPDATE pomodoro
SET effort = $1
SET distraction = $2
SET productivity = $3
WHERE user_id = $4;


-- name: Inactivate :exec
UPDATE pomodoro
SET deleted_at = $1
WHERE id = $2
AND user_id = $3;
