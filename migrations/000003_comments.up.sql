CREATE TABLE comments (
                          id uuid primary key  default gen_random_uuid(),
                          composition_id uuid,
                          user_id uuid,
                          content TEXT,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE ,
                          deleted_at TIMESTAMP WITH TIME ZONE
);