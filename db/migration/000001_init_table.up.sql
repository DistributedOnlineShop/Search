CREATE TABLE "search_index" (
  "product_id" VARCHAR(12) PRIMARY KEY NOT NULL,
  "product_name" VARCHAR(255) NOT NULL,
  "product_desc" TEXT,
  "product_attributes" JSON,
  "indexed_at" TIMESTAMP(0) NOT NULL DEFAULT NOW()
);

CREATE TABLE "search_log" (
  "search_id" UUID PRIMARY KEY NOT NULL,  
  "user_id" UUID,    
  "search_query" VARCHAR(255) NOT NULL,
  "search_filters" JSON,
  "results_count" INT NOT NULL,
  "searched_at" TIMESTAMP(0) NOT NULL DEFAULT NOW()
);