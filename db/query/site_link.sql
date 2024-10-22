-- name: GetDataByDate :many
SELECT * FROM demarcation_site_link
WHERE date = ?;

-- name: GetRegionData :many
SELECT * FROM demarcation_site_link
WHERE date = ? and region = ?;

-- name: GetCityData :many
SELECT * FROM demarcation_site_link
WHERE date = ? and city = ? LIMIT 10;

-- name: ListCity :many
SELECT distinct(city) FROM demarcation_site_link;

-- name: ListRegion :many
SELECT distinct(region) FROM demarcation_site_link;