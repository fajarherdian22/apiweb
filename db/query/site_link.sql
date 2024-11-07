-- name: GetDataByDate :many
SELECT * FROM demarcation_site_link
WHERE date = ? and city != "";

-- name: GetRegionData :many
SELECT * FROM demarcation_site_link
WHERE date = ? and region = ?;

-- name: GetCityData :many
SELECT * FROM demarcation_site_link
WHERE date = ? and city = ?;

-- name: ListCity :many
SELECT distinct(city) FROM demarcation_site_link WHERE city !="";

-- name: ListRegion :many
SELECT distinct(region) FROM demarcation_site_link;