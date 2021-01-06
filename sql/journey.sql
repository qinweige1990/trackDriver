CREATE TABLE `journeys` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `start` text,
                          `end` text,
                          `driver` varchar(40) NOT NULL DEFAULT '',
                          `status` int DEFAULT 1,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;