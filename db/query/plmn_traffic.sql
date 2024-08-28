-- name: GetPlmnData :many
SELECT * from plmn_traffic
WHERE date = ?;

-- name: GetPlmnDataCity :many
SELECT * from plmn_traffic
WHERE date = ? and ran_site_city = ?;

-- name: GetFilterName :many
SELECT DISTINCT a.ran_site_city, a.ran_site_region from plmn_traffic a;

-- name: GetAllPlmn :many
SELECT * from plmn_traffic;