CREATE TABLE IF NOT EXISTS roles (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR(200),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);