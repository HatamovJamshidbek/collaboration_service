CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined');
CREATE TABLE invitations (
                             id uuid primary key  default gen_random_uuid(),
                             composition_id uuid,
                             inviter_id uuid,
                             invitee_id uuid,
                             status invitation_status DEFAULT 'pending',
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE ,
                             deleted_at TIMESTAMP WITH TIME ZONE
);