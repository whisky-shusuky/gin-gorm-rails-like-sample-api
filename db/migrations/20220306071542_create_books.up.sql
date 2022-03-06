CREATE TABLE books (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  book_name varchar(255) DEFAULT NULL,
  book_description varchar(255) DEFAULT NULL,
  sales int(10) unsigned DEFAULT NULL,
  created_at datetime(3) DEFAULT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

COMMIT;