--  relayer version: 1.0.0
--
--  (C) Copyright 2022 Satish Babariya (https://satishbabariya.com/) and others.
--
--  Licensed under the Apache License, Version 2.0 (the "License");
--  you may not use this file except in compliance with the License.
--  You may obtain a copy of the License at
--
--       http://www.apache.org/licenses/LICENSE-2.0
--
--  Unless required by applicable law or agreed to in writing, software
--  distributed under the License is distributed on an "AS IS" BASIS,
--  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
--  See the License for the specific language governing permissions and
--  limitations under the License.
--
--  Contributors:
--      satish babariya (satish.babariya@gmail.com)
--
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Drop all existing tables [TODO: Remove this once we have a migration tool]
DROP TABLE IF EXISTS "public"."users";

-- Set Time Zone to UTC
SET TIME ZONE 'UTC';

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT Uuid_generate_v4(), 
  username TEXT, 
  created_at TIMESTAMP NOT NULL DEFAULT Now(), 
  updated_at TIMESTAMP NOT NULL DEFAULT Now()
);
