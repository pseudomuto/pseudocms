-- # 1 column
-- # row 1
-- ## 190
CREATE TABLE public.schema_migration (
	version VARCHAR(14) NOT NULL,
	CONSTRAINT schema_migration_pkey PRIMARY KEY (version ASC),
	UNIQUE INDEX schema_migration_version_idx (version ASC)
);
-- # row 2
-- ## 282
CREATE TABLE public.definitions (
	id UUID NOT NULL,
	name VARCHAR(200) NOT NULL,
	description VARCHAR(1000) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT definitions_pkey PRIMARY KEY (id ASC),
	UNIQUE INDEX definitions_name_idx (name ASC)
);
-- # row 3
-- ## 445
CREATE TABLE public.fields (
	id UUID NOT NULL,
	definition_id UUID NOT NULL,
	name VARCHAR(200) NOT NULL,
	description VARCHAR(1000) NOT NULL,
	kind VARCHAR(100) NOT NULL,
	constraints VARCHAR(2000) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT fields_pkey PRIMARY KEY (id ASC),
	INDEX fields_definition_id_idx (definition_id ASC),
	UNIQUE INDEX fields_definition_id_name_idx (definition_id ASC, name ASC)
);
-- # row 4
-- ## 146
ALTER TABLE public.fields ADD CONSTRAINT fields_definitions_id_fk FOREIGN KEY (definition_id) REFERENCES public.definitions(id) ON DELETE CASCADE;
-- # row 5
-- ## 115
-- Validate foreign key constraints. These can fail if there was unvalidated data during the SHOW CREATE ALL TABLES
-- # row 6
-- ## 71
ALTER TABLE public.fields VALIDATE CONSTRAINT fields_definitions_id_fk;
-- # 6 rows
