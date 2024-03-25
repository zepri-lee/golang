CREATE TABLE STOCK (
	PRODUCT_ID          BIGINT NOT NULL PRIMARY KEY,
	PRODUCT_NAME        VARCHAR(255) NOT NULL,
	STOCK_QUANTITY      INT,
	STOCK_LOCATION      VARCHAR(255)
)

CREATE TABLE SALE (
	SEQ_NO              BIGINT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
	PRODUCT_ID          BIGINT NOT NULL,
	SALE_COUNT          INT,
	SALE_DATE           DATETIME
)
