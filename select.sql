WITH filtered_by_latitude AS (
    SELECT * FROM city_bikes
    WHERE start_station_latitude BETWEEN 40.72 AND 40.73
), filtered_by_tripduration AS (
    SELECT * FROM filtered_by_latitude
    WHERE tripduration < (SELECT AVG(tripduration) FROM city_bikes)
)
SELECT start_station_name, COUNT(*) AS appearances
FROM filtered_by_tripduration
GROUP BY start_station_name
ORDER BY appearances DESC
LIMIT 1;
