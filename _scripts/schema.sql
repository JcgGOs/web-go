DROP TABLE IF EXISTS t_user;
CREATE TABLE t_user
(
  id        bigint(11) NOT NULL AUTO_INCREMENT,
  create_at DATETIME      DEFAULT NULL,
  update_at DATETIME      DEFAULT NULL,

  status    varchar(255)  DEFAULT NULL,
  username  varchar(255)  DEFAULT NULL,
  password  varchar(255)  DEFAULT NULL,

  email     varchar(255)  DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY UK_user(username)
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS t_topic;
CREATE TABLE t_topic
(
  id        bigint(11)      NOT NULL AUTO_INCREMENT,
  create_at DATETIME     DEFAULT NULL,
  update_at DATETIME     DEFAULT NULL,
  status    varchar(255) DEFAULT NULL,

  title     varchar(255) DEFAULT NULL,
  content   TEXT         DEFAULT NULL,
  user_id   int(11)      DEFAULT 0,

  PRIMARY KEY (id)
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8;