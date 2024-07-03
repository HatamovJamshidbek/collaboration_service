CREATE TYPE collaboration_role AS ENUM ('owner', 'collaborator', 'viewer');

CREATE TABLE collaborations (
                                id uuid primary key  default gen_random_uuid(),
                                composition_id uuid,
                                user_id uuid,
                                role collaboration_role DEFAULT 'collaborator',
                                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP WITH TIME ZONE ,
                                deleted_at TIMESTAMP WITH TIME ZONE
);