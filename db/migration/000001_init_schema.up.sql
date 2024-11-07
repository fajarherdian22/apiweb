CREATE TABLE IF NOT EXISTS `demarcation_site_link` (
    `date` varchar(20) NOT NULL,
    `siteid` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `city` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `region` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `ran_site_longitude` double null,
    `ran_site_latitude` double null,
    `site_destination` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `longitude_destination` double null,
    `latitude_destination` double null,
    `link` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `type_transport` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `interface` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `tlp` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `ioh_data_traffic_4g` double null,
    `availability` double null,
    `eut` double null,
    `cqi` double null,
    `prb` double null,
    `accesibility` double null,
    `s1_sr` double null,
    `erab_drop` double null,
    `rrc_sr` double null,
    `erab_sr` double null,
    `inter_freq` double null,
    `intra_freq` double null,
    `irat_hosr` double null,
    `ul_rssi` double null,
    `SE_DL` double null,
    `csfb_prep_sr` double null,
    `csfb_sr` double null,
    `syn_ack_ack_delay` double null,
    `syn_syn_ack_delay` double null,
    `downlink_tcp_retransmission_rate` double null,
    `max_util` double null,
    `capacity` double null,
    `traffic_origin` double null,
    `availability_origin` double null,
    `eut_origin` double null,
    `cqi_origin` double null,
    `prb_origin` double null,
    `accesibility_origin` double null,
    `s1_sr_origin` double null,
    `erab_drop_origin` double null,
    `RRC_SR_origin` double null,
    `erab_sr_origin` double null,
    `inter_freq_origin` double null,
    `intra_freq_origin` double null,
    `irat_hosr_origin` double null,
    `ul_rssi_origin` double null,
    `SE_DL_origin` double null,
    `csfb_prep_sr_origin` double null,
    `csfb_sr_origin` double null,
    `syn_ack_ack_delay_origin` double null,
    `syn_syn_ack_delay_origin` double null,
    `downlink_tcp_retransmission_rate_origin` double null,
    PRIMARY KEY (`date`, `siteid`, `site_destination`),
    KEY `date` (`date`)
);

CREATE TABLE topo_link (
    siteid VARCHAR(50) not NULL,
    interface_a VARCHAR(100) not NULL,
    TLP_a VARCHAR(100) not NULL,
    city VARCHAR(100) not NULL,
    longitude double NULL,
    latitude double NULL,
    ne_a VARCHAR(100) not NULL,
    ne_b VARCHAR(100) not NULL,
    site_destination VARCHAR(50) not NULL,
    interface_b VARCHAR(100) not NULL,
    TLP_b VARCHAR(100) not NULL,
    longitude_destination double NULL,
    latitude_destination double NULL,
    link VARCHAR(100) not NULL,
    type_transport VARCHAR(20) not NULL,
    level_ne_a VARCHAR(50) not NULL,
    level_ne_b VARCHAR(50) not NULL,
    bandwith double NULL,
    max_util double NULL,
    capacity double NULL,
    INDEX idx_city (city)
);