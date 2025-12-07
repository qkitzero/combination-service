CREATE TABLE element_category (
  element_id VARCHAR(36) NOT NULL,
  category_id VARCHAR(36) NOT NULL,
  PRIMARY KEY (element_id, category_id),
  FOREIGN KEY (element_id) REFERENCES elements(id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
