CREATE TABLE shops (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  shop_name varchar(255) DEFAULT NULL,
  shop_description varchar(255) DEFAULT NULL,
  created_at datetime(3) DEFAULT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0;

COMMIT;