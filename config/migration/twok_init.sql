# ******************************************************************************
# Create new tables
# ******************************************************************************
# CREATE DATABASE IF NOT EXISTS twok DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
# USE twok;

# ******************************************************************************
# Create tables
# ******************************************************************************
CREATE TABLE machines (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,

    name VARCHAR(50) NOT NULL,
    os VARCHAR(25) NOT NULL,
    status VARCHAR(25) NOT NULL,

    PRIMARY KEY (id)
);