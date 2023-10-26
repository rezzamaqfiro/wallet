-- Seed data for "orders" table
INSERT INTO "orders" ("invoice","total_payment","user_id","status","created_at", "created_by", "deleted")
VALUES
  (uuid_generate_v4(),260000,'009888aa-43ee-4bf6-8387-ec27a8f2f0ea','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),1500000,'092d38a4-7d05-43d9-9103-7609cb0ed171','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),400000,'d8fbb968-c6fb-4e6b-9c67-e346be789b36','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),3500000,'beb4581e-63cc-4fae-8863-d4e8199be9fc','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),220000,'aba10811-9cac-486e-8673-38a22910b6fa','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),4600000,'0f675e3c-03f8-4775-b76f-dfa5f389c9ab','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),500000,'a8acfae0-3db3-4b5c-a659-9539c40b8b02','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),422000,'c15852f2-45d3-47a1-a0a0-f8cde554467d','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),890000,'a5bb1ba7-14c3-4a87-982a-a2ceba68a3e1','PENDING',NOW(),'Admin',false),
  (uuid_generate_v4(),12000000,'c040ee71-8097-49b3-965f-b3d37e953112','PENDING',NOW(),'Admin',false);

-- Seed data for "users" table
INSERT INTO "users" ("user_id", "name", "balance", "created_at", "created_by", "deleted")
VALUES
  ('009888aa-43ee-4bf6-8387-ec27a8f2f0ea', 'Tina', 200000, NOW(), 'Admin', false),
  ('092d38a4-7d05-43d9-9103-7609cb0ed171', 'Maman', 50000, NOW(), 'Admin', false),
  ('d8fbb968-c6fb-4e6b-9c67-e346be789b36', 'Sari', 600000, NOW(), 'Admin', false),
  ('beb4581e-63cc-4fae-8863-d4e8199be9fc', 'Rina', 30000, NOW(), 'Admin', false),
  ('aba10811-9cac-486e-8673-38a22910b6fa', 'Wawan', 230000, NOW(), 'Admin', false),
  ('0f675e3c-03f8-4775-b76f-dfa5f389c9ab', 'Sandy', 20000, NOW(), 'Admin', false),
  ('a8acfae0-3db3-4b5c-a659-9539c40b8b02', 'Tono', 100000, NOW(), 'Admin', false),
  ('c15852f2-45d3-47a1-a0a0-f8cde554467d', 'Jardem', 5000, NOW(), 'Admin', false),
  ('a5bb1ba7-14c3-4a87-982a-a2ceba68a3e1', 'Kirab', 135000, NOW(), 'Admin', false),
  ('c040ee71-8097-49b3-965f-b3d37e953112', 'Lala', 1000, NOW(), 'Admin', false);