CREATE TABLE "public".b_collection (
  id              uuid NOT NULL, 
  user_profile_id uuid NOT NULL, 
  name            varchar(255) NOT NULL, 
  day             date NOT NULL, 
  CONSTRAINT collection_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".b_collection_comment (
  id              uuid NOT NULL, 
  user_profile_id uuid NOT NULL, 
  b_collection_id uuid NOT NULL, 
  date_created    timestamp with time zone NOT NULL, 
  content         varchar(255) NOT NULL, 
  CONSTRAINT collection_comment_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".b_collection_reaction (
  id              uuid NOT NULL, 
  user_profile_id uuid NOT NULL, 
  b_collection_id uuid NOT NULL, 
  date_created    timestamp with time zone NOT NULL, 
  reaction_id     int4 NOT NULL, 
  CONSTRAINT collection_reaction_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".collection_post (
  b_collection_id uuid NOT NULL, 
  post_id         uuid NOT NULL, 
  date_added      timestamp with time zone NOT NULL, 
  CONSTRAINT collection_post_id_pkey 
    PRIMARY KEY (b_collection_id, 
  post_id));

CREATE TABLE "public".friendship (
  request_user_profile_id  uuid NOT NULL, 
  response_user_profile_id uuid NOT NULL, 
  status                   varchar(10) NOT NULL, 
  date_updated             timestamp with time zone NOT NULL);

CREATE TABLE "public".post (
  id              uuid NOT NULL, 
  user_profile_id uuid NOT NULL, 
  caption         varchar(255) NOT NULL, 
  date_created    timestamp with time zone NOT NULL, 
  image           varchar(255) NOT NULL, 
  image_2         varchar(255), 
  reaction_count  int4 DEFAULT 0 NOT NULL, 
  in_collection   bool DEFAULT 'false' NOT NULL, 
  CONSTRAINT post_id_pkey 
    PRIMARY KEY (id));
    
CREATE TABLE "public".post_comment (
  id              uuid NOT NULL, 
  post_id         uuid NOT NULL, 
  user_profile_id uuid NOT NULL, 
  date_created    timestamp with time zone NOT NULL, 
  content         varchar(255) NOT NULL, 
  CONSTRAINT post_comment_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".post_reaction (
  id              uuid NOT NULL, 
  post_id         uuid NOT NULL, 
  user_profile_id uuid NOT NULL, 
  date_created    timestamp with time zone NOT NULL, 
  reaction_id     int4 NOT NULL, 
  CONSTRAINT post_like_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".reaction (
  id   int4 NOT NULL, 
  name varchar(25) NOT NULL, 
  icon varchar(255) NOT NULL, 
  CONSTRAINT reaction_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".user_profile (
  id           uuid NOT NULL, 
  display_name varchar(30) NOT NULL, 
  first_name   varchar(15) NOT NULL, 
  last_name    varchar(30) NOT NULL, 
  email        varchar(255) NOT NULL, 
  phone        varchar(15) NOT NULL, 
  birthdate    date NOT NULL, 
  password     text NOT NULL, 
  profile_pic  text NOT NULL, 
  CONSTRAINT profile_id_pkey 
    PRIMARY KEY (id));

CREATE TABLE "public".user_settings (
  id              SERIAL NOT NULL, 
  "key"           int4 NOT NULL, 
  value           int4 NOT NULL, 
  user_profile_id uuid NOT NULL, 
  PRIMARY KEY (id));

CREATE INDEX b_collection_user_profile_id 
  ON "public".b_collection (user_profile_id);
CREATE INDEX post_comment_post_id 
  ON "public".post_comment (post_id);
CREATE INDEX post_reaction_post_id 
  ON "public".post_reaction (post_id);
  
ALTER TABLE "public".b_collection_comment ADD CONSTRAINT b_collection_comment_b_collection_id_fkey FOREIGN KEY (b_collection_id) REFERENCES "public".b_collection (id);
ALTER TABLE "public".b_collection_comment ADD CONSTRAINT b_collection_comment_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".b_collection_reaction ADD CONSTRAINT b_collection_reaction_b_collection_id_fkey FOREIGN KEY (b_collection_id) REFERENCES "public".b_collection (id);
ALTER TABLE "public".b_collection_reaction ADD CONSTRAINT b_collection_reaction_reaction_id_fkey FOREIGN KEY (reaction_id) REFERENCES "public".reaction (id);
ALTER TABLE "public".b_collection_reaction ADD CONSTRAINT b_collection_reaction_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".b_collection ADD CONSTRAINT b_collection_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".collection_post ADD CONSTRAINT collection_post_b_collection_id_fkey FOREIGN KEY (b_collection_id) REFERENCES "public".b_collection (id);
ALTER TABLE "public".collection_post ADD CONSTRAINT collection_post_post_post_id_fkey FOREIGN KEY (post_id) REFERENCES "public".post (id);
ALTER TABLE "public".post_comment ADD CONSTRAINT comment_post_id_fkey FOREIGN KEY (post_id) REFERENCES "public".post (id);
ALTER TABLE "public".post_comment ADD CONSTRAINT comment_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".friendship ADD CONSTRAINT friendship_request_user_profile_id_fkey FOREIGN KEY (response_user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".friendship ADD CONSTRAINT friendship_response_user_profile_id_fkey FOREIGN KEY (request_user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".post_reaction ADD CONSTRAINT post_reaction_post_id_fkey FOREIGN KEY (post_id) REFERENCES "public".post (id);
ALTER TABLE "public".post_reaction ADD CONSTRAINT post_reaction_reaction_id_fkey FOREIGN KEY (reaction_id) REFERENCES "public".reaction (id);
ALTER TABLE "public".post_reaction ADD CONSTRAINT post_reaction_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".post ADD CONSTRAINT post_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);
ALTER TABLE "public".user_settings ADD CONSTRAINT user_settings_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES "public".user_profile (id);

INSERT INTO
  user_profile (
    id,
    display_name,
    first_name,
    last_name,
    email,
    phone,
    birthdate,
    password,
    profile_pic
  )
VALUES
  (
    '12b02b58-7cad-11ed-a1eb-0242ac120002',
    'TheAlbinoGiannis',
    'Jared',
    'Heidt',
    'jared@babalaas.com',
    '8148675309',
    '2001-02-01',
    'taco',
    'https://library.sportingnews.com/styles/crop_style_16_9_mobile_2x/s3/2023-03/GettyImages-1471642056.jpg?itok=Mcrb_5re'
  ),
  (
    '4f9da08c-62eb-4dc0-9ab5-0d7183437695',
    'HartAttack0328',
    'Sean',
    'Hart',
    'sean@babalaas.com',
    '8148675309',
    '2001-03-28',
    'password',
    'https://library.sportingnews.com/styles/crop_style_16_9_mobile_2x/s3/2023-03/GettyImages-1471642056.jpg?itok=Mcrb_5re'
  );

  INSERT INTO
  post (
    id,
    user_profile_id,
    caption,
    date_created,
    image,
    image_2,
    reaction_count,
    in_collection
  )
VALUES
  (
    '0b318a7a-7cad-11ed-a1eb-0242ac120002',
    '12b02b58-7cad-11ed-a1eb-0242ac120002',
    'Fun in the Sealwoves City!',
    '2022-08-12 17:15:00-5',
    'https://babalaas-bucket.s3.amazonaws.com/sean-seawolves.jpg',
    '',
    0,
    false
  );

INSERT INTO
  post (
    id,
    user_profile_id,
    caption,
    date_created,
    image,
    image_2,
    reaction_count,
    in_collection
  )
VALUES
  (
    '93b4adb8-2b31-47cd-ae92-fd674eadf3b3',
    '12b02b58-7cad-11ed-a1eb-0242ac120002',
    'Walkoff wins are better with beer!',
    '2022-08-12 21:37:00-05',
    'https://babalaas-bucket.s3.amazonaws.com/seawolves.jpg',
    '',
    0,
    false
  );

INSERT INTO
  b_collection (id, user_profile_id, name, day)
VALUES
  (
    '1823ea5c-7cad-11ed-a1eb-0242ac120002',
    '12b02b58-7cad-11ed-a1eb-0242ac120002',
    'Best Buck Night Ever',
    '2022-08-12'
  );

  
INSERT INTO
  collection_post (b_collection_id, post_id, date_added)
VALUES
  (
    '1823ea5c-7cad-11ed-a1eb-0242ac120002',
    '0b318a7a-7cad-11ed-a1eb-0242ac120002',
    '2023-01-22 23:47:00-5'
  ),
  (
    '1823ea5c-7cad-11ed-a1eb-0242ac120002',
    '93b4adb8-2b31-47cd-ae92-fd674eadf3b3',
    '2023-01-22 23:47:00-5'
  );

INSERT INTO
  friendship (request_user_profile_id, response_user_profile_id, status, date_updated)
VALUES
  (
    '12b02b58-7cad-11ed-a1eb-0242ac120002',
    '4f9da08c-62eb-4dc0-9ab5-0d7183437695',
    'accepted',
    '2023-03-26 23:47:00-5'
  );
INSERT INTO "public".reaction
  (id, 
  name, 
  icon) 
VALUES 
  (1, 
  'Like', 
  'https://png.pngtree.com/png-vector/20190224/ourmid/pngtree-vector-like-icon-png-image_699751.jpg');


INSERT INTO "public".post_comment
  (id, 
  post_id, 
  user_profile_id, 
  date_created, 
  content) 
VALUES 
  ('95f709f0-e260-11ed-b5ea-0242ac120002', 
  '0b318a7a-7cad-11ed-a1eb-0242ac120002', 
  '4f9da08c-62eb-4dc0-9ab5-0d7183437695', 
  '2022-08-12 21:37:00-05', 
  'Home is where the heart is');
