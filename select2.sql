SELECT start_station_name, AVG(tripduration) AS appearances
FROM city_bikes
WHERE 
    (start_station_latitude BETWEEN 40.72 AND 40.73)
    AND
    (tripduration > (SELECT AVG(tripduration) FROM city_bikes))
GROUP BY start_station_name
ORDER BY appearances DESC;
