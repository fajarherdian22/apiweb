CREATE TABLE `plmn_traffic` (
    `DATE` varchar(20) NOT NULL,
    `moentity` varchar(15) NOT NULL,
    `ran_site_city` varchar(50) NOT NULL,
    `ran_site_region` varchar(50) NOT NULL,
    `ran_site_longitude` double NOT NULL,
    `ran_site_latitude` double NOT NULL,
    `ran_site_site_vendor` varchar(20) NOT NULL,
    `traffic` double NOT NULL,
    `traffic_type` varchar(20) NOT NULL
);