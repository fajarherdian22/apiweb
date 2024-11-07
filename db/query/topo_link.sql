-- name: GetTopo :many
SELECT
    *
FROM
    topo_link
WHERE
    city != "";

-- name: GetTopoCity :many
SELECT
    *
FROM
    topo_link
WHERE
    city = ?;

-- name: ListTopoCity :many
SELECT
    distinct(city)
FROM
    topo_link
WHERE
    city != "";