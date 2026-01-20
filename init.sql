CREATE DATABASE url_shortener;
USE url_shortener;
CREATE TABLE links(
    code       VARCHAR(16) PRIMARY KEY,
    url        TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);